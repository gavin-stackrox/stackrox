package service

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/complianceoperator/v2/checkresults/datastore"
	resultMocks "github.com/stackrox/rox/central/complianceoperator/v2/checkresults/datastore/mocks"
	integrationMocks "github.com/stackrox/rox/central/complianceoperator/v2/integration/datastore/mocks"
	profileDatastore "github.com/stackrox/rox/central/complianceoperator/v2/profiles/datastore/mocks"
	ruleMocks "github.com/stackrox/rox/central/complianceoperator/v2/rules/datastore/mocks"
	scanConfigMocks "github.com/stackrox/rox/central/complianceoperator/v2/scanconfigurations/datastore/mocks"
	convertUtils "github.com/stackrox/rox/central/convert/testutils"
	v1 "github.com/stackrox/rox/generated/api/v1"
	apiV2 "github.com/stackrox/rox/generated/api/v2"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/fixtures/fixtureconsts"
	"github.com/stackrox/rox/pkg/grpc/testutils"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/testconsts"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var (
	integration1 = &storage.ComplianceIntegration{
		Id:                  "",
		ClusterId:           testconsts.Cluster1,
		ComplianceNamespace: fixtureconsts.Namespace1,
		Version:             "2",
		StatusErrors:        []string{"test error"},
	}

	integration2 = &storage.ComplianceIntegration{
		Id:                  "",
		ClusterId:           testconsts.Cluster2,
		ComplianceNamespace: fixtureconsts.Namespace1,
		Version:             "2",
		StatusErrors:        []string{"test error"},
	}

	integration3 = &storage.ComplianceIntegration{
		Id:                  "",
		ClusterId:           testconsts.Cluster3,
		ComplianceNamespace: fixtureconsts.Namespace1,
		Version:             "2",
		StatusErrors:        []string{"test error"},
	}
)

func TestAuthz(t *testing.T) {
	testutils.AssertAuthzWorks(t, &serviceImpl{})
}

func TestComplianceResultsService(t *testing.T) {
	suite.Run(t, new(ComplianceResultsServiceTestSuite))
}

type ComplianceResultsServiceTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller

	ctx             context.Context
	resultDatastore *resultMocks.MockDataStore
	scanConfigDS    *scanConfigMocks.MockDataStore
	integrationDS   *integrationMocks.MockDataStore
	ruleDS          *ruleMocks.MockDataStore
	service         Service
	profilsDS       *profileDatastore.MockDataStore
}

func (s *ComplianceResultsServiceTestSuite) SetupSuite() {
	s.T().Setenv(features.ComplianceEnhancements.EnvVar(), "true")
	if !features.ComplianceEnhancements.Enabled() {
		s.T().Skip("Skip test when compliance enhancements are disabled")
		s.T().SkipNow()
	}

	s.ctx = sac.WithAllAccess(context.Background())
}

func (s *ComplianceResultsServiceTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.resultDatastore = resultMocks.NewMockDataStore(s.mockCtrl)
	s.scanConfigDS = scanConfigMocks.NewMockDataStore(s.mockCtrl)
	s.integrationDS = integrationMocks.NewMockDataStore(s.mockCtrl)
	s.profilsDS = profileDatastore.NewMockDataStore(s.mockCtrl)
	s.ruleDS = ruleMocks.NewMockDataStore(s.mockCtrl)

	s.service = New(s.resultDatastore, s.scanConfigDS, s.integrationDS, s.profilsDS, s.ruleDS)
}

func (s *ComplianceResultsServiceTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceScanResults() {
	testCases := []struct {
		desc         string
		query        *apiV2.RawQuery
		expectedResp []*apiV2.ComplianceScanResult
		expectedErr  error
		found        bool
		setMocks     func()
	}{
		{
			desc:         "Empty query",
			query:        &apiV2.RawQuery{Query: ""},
			expectedErr:  nil,
			expectedResp: convertUtils.GetConvertedComplianceResults(s.T()),
			found:        true,
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(convertUtils.GetComplianceStorageResults(s.T()), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig2").Return(getTestRec("scanConfig2"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig3").Return(getTestRec("scanConfig3"), nil).Times(1)
			},
		},
		{
			desc:        "Query with search field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			expectedErr: nil,
			found:       true,
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(convertUtils.GetOneClusterComplianceStorageResults(s.T(), fixtureconsts.Cluster1), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig2").Return(getTestRec("scanConfig2"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig3").Return(getTestRec("scanConfig3"), nil).Times(1)
			},
		},
		{
			desc: "Query with custom pagination",
			query: &apiV2.RawQuery{
				Query:      "",
				Pagination: &apiV2.Pagination{Limit: 1},
			},
			expectedErr: nil,
			found:       true,
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(1)).ProtoQuery()
				returnResults := []*storage.ComplianceOperatorCheckResultV2{
					convertUtils.GetComplianceStorageResults(s.T())[0],
				}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(returnResults, nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
			},
		},
		{
			desc:        "Query with non-existent field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:id"},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance scan results for query %v", &apiV2.RawQuery{Query: "Cluster ID:id"}),
			found:       false,
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(nil, nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceScanResults(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(convertUtils.GetConvertedComplianceResults(s.T()), results.GetScanResults())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceClusterScanStats() {
	testCases := []struct {
		desc         string
		query        *apiV2.RawQuery
		expectedResp []*apiV2.ComplianceClusterScanStats
		expectedErr  error
		setMocks     func()
	}{
		{
			desc:        "Empty query",
			query:       &apiV2.RawQuery{Query: ""},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterScanStats{
				convertUtils.GetComplianceClusterScanV2Count(s.T(), fixtureconsts.Cluster1),
				convertUtils.GetComplianceClusterScanV2Count(s.T(), fixtureconsts.Cluster2),
				convertUtils.GetComplianceClusterScanV2Count(s.T(), fixtureconsts.Cluster3),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				results := []*datastore.ResourceResultCountByClusterScan{
					convertUtils.GetComplianceStorageClusterScanCount(s.T(), fixtureconsts.Cluster1),
					convertUtils.GetComplianceStorageClusterScanCount(s.T(), fixtureconsts.Cluster2),
					convertUtils.GetComplianceStorageClusterScanCount(s.T(), fixtureconsts.Cluster3),
				}
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.resultDatastore.EXPECT().ComplianceCheckResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc:        "Query with search field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterScanStats{
				convertUtils.GetComplianceClusterScanV2Count(s.T(), fixtureconsts.Cluster1),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				results := []*datastore.ResourceResultCountByClusterScan{
					convertUtils.GetComplianceStorageClusterScanCount(s.T(), fixtureconsts.Cluster1),
				}
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.resultDatastore.EXPECT().ComplianceCheckResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc:        "Query with non-existent field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:id"},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance cluster scan stats for query %v", &apiV2.RawQuery{Query: "Cluster ID:id"}),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				s.resultDatastore.EXPECT().ComplianceCheckResultStats(gomock.Any(), expectedQ).Return(nil, nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceClusterScanStats(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(tc.expectedResp, results.GetScanStats())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceOverallClusterStats() {
	testCases := []struct {
		desc         string
		query        *apiV2.RawQuery
		expectedResp []*apiV2.ComplianceClusterOverallStats
		expectedErr  error
		setMocks     func()
	}{
		{
			desc:        "Empty query",
			query:       &apiV2.RawQuery{Query: ""},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterOverallStats{
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster1),
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster2),
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster3),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				results := []*datastore.ResultStatusCountByCluster{
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster1),
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster2),
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster3),
				}

				s.resultDatastore.EXPECT().CountByField(gomock.Any(), search.EmptyQuery(), search.ClusterID)
				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster1).Return([]*storage.ComplianceIntegration{integration1}, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster2).Return([]*storage.ComplianceIntegration{integration2}, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster3).Return([]*storage.ComplianceIntegration{integration3}, nil).Times(1)
			},
		},
		{
			desc:        "Query with search field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterOverallStats{
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster1),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				countQuery := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()

				results := []*datastore.ResultStatusCountByCluster{
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster1),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ClusterID)
				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster1).Return([]*storage.ComplianceIntegration{integration1}, nil).Times(1)
			},
		},
		{
			desc:        "Query with non-existent field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:id"},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance cluster scan stats for query %v", &apiV2.RawQuery{Query: "Cluster ID:id"}),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				countQuery := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").ProtoQuery()

				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ClusterID)
				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(nil, nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceOverallClusterStats(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(tc.expectedResp, results.GetScanStats())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceClusterStats() {
	testCases := []struct {
		desc         string
		request      *apiV2.ComplianceProfileResultsRequest
		expectedResp []*apiV2.ComplianceClusterOverallStats
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query just profile",
			request: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "test-profile",
			},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterOverallStats{
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster1),
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster2),
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster3),
			},
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "test-profile").ProtoQuery(),
					search.EmptyQuery(),
				)

				countQuery := expectedQ.Clone()

				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResultStatusCountByCluster{
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster1),
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster2),
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster3),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ClusterID)
				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster1).Return([]*storage.ComplianceIntegration{integration1}, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster2).Return([]*storage.ComplianceIntegration{integration2}, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster3).Return([]*storage.ComplianceIntegration{integration3}, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			request: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "test-profile",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceClusterOverallStats{
				convertUtils.GetComplianceClusterV2Count(s.T(), fixtureconsts.Cluster1),
			},
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "test-profile").ProtoQuery(),
					search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery(),
				)

				countQuery := expectedQ.Clone()

				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResultStatusCountByCluster{
					convertUtils.GetComplianceStorageClusterCount(s.T(), fixtureconsts.Cluster1),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ClusterID)
				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				s.integrationDS.EXPECT().GetComplianceIntegrationByCluster(gomock.Any(), fixtureconsts.Cluster1).Return([]*storage.ComplianceIntegration{integration1}, nil).Times(1)
			},
		},
		{
			desc: "Query with non-existent field",
			request: &apiV2.ComplianceProfileResultsRequest{
				Query: &apiV2.RawQuery{Query: "Cluster ID:id"},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {

			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceClusterStats(s.ctx, tc.request)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(tc.expectedResp, results.GetScanStats())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) GetComplianceOverallClusterCount() {
	testCases := []struct {
		desc         string
		query        *apiV2.RawQuery
		expectedResp *apiV2.CountComplianceScanResults
		expectedErr  error
		setMocks     func()
	}{
		{
			desc:        "Empty query",
			query:       &apiV2.RawQuery{Query: ""},
			expectedErr: nil,
			expectedResp: &apiV2.CountComplianceScanResults{
				Count: 3,
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), expectedQ, search.ClusterID).Return(3, nil).Times(1)
			},
		},
		{
			desc:        "Query with search field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			expectedErr: nil,
			expectedResp: &apiV2.CountComplianceScanResults{
				Count: 1,
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), expectedQ, search.ClusterID).Return(1, nil).Times(1)
			},
		},
		{
			desc:        "Query with non-existent field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:id"},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance cluster scan stats for query %v", &apiV2.RawQuery{Query: "Cluster ID:id"}),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()

				s.resultDatastore.EXPECT().ComplianceClusterStats(gomock.Any(), expectedQ).Return(nil, nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceOverallClusterCount(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(tc.expectedResp.Count, results.Count)
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceScanResult() {
	testCases := []struct {
		desc         string
		query        *apiV2.ResourceByID
		expectedResp *apiV2.ComplianceClusterCheckStatus
		expectedErr  error
		found        bool
		setMocks     func()
	}{
		{
			desc:         "ID exists",
			query:        &apiV2.ResourceByID{Id: uuid.NewDummy().String()},
			expectedErr:  nil,
			expectedResp: convertUtils.GetConvertedComplianceResult(s.T()),
			found:        true,
			setMocks: func() {
				s.resultDatastore.EXPECT().GetComplianceCheckResult(gomock.Any(), uuid.NewDummy().String()).Return(convertUtils.GetComplianceStorageResult(s.T()), true, nil).Times(1)
			},
		},
		{
			desc:        "Query with non-existent record",
			query:       &apiV2.ResourceByID{Id: uuid.NewDummy().String()},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "compliance check result with id %q does not exist", uuid.NewDummy().String()),
			found:       false,
			setMocks: func() {
				s.resultDatastore.EXPECT().GetComplianceCheckResult(gomock.Any(), uuid.NewDummy().String()).Return(nil, false, nil).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			result, err := s.service.GetComplianceScanCheckResult(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(convertUtils.GetConvertedComplianceResult(s.T()), result)
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceScanConfigurationResults() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceScanResultsRequest
		expectedResp []*apiV2.ComplianceScanResult
		expectedErr  error
		found        bool
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceScanResultsRequest{
				ScanConfigName: "scanConfig1",
				Query:          &apiV2.RawQuery{Query: ""},
			},
			expectedErr:  nil,
			expectedResp: convertUtils.GetConvertedComplianceResults(s.T()),
			found:        true,
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorScanConfigName, "scanConfig1").ProtoQuery(),
					search.EmptyQuery(),
				)
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(convertUtils.GetComplianceStorageResults(s.T()), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig2").Return(getTestRec("scanConfig2"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig3").Return(getTestRec("scanConfig3"), nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceScanResultsRequest{
				ScanConfigName: "scanConfig1",
				Query:          &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: nil,
			found:       true,
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorScanConfigName, "scanConfig1").ProtoQuery(),
					expectedQ,
				)
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(convertUtils.GetOneClusterComplianceStorageResults(s.T(), fixtureconsts.Cluster1), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig2").Return(getTestRec("scanConfig2"), nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig3").Return(getTestRec("scanConfig3"), nil).Times(1)
			},
		},
		{
			desc: "Query with custom pagination",
			query: &apiV2.ComplianceScanResultsRequest{
				ScanConfigName: "scanConfig1",
				Query: &apiV2.RawQuery{Query: "",
					Pagination: &apiV2.Pagination{Limit: 1}},
			},
			expectedErr: nil,
			found:       true,
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorScanConfigName, "scanConfig1").ProtoQuery(),
					search.EmptyQuery(),
				)
				expectedQ.Pagination = &v1.QueryPagination{Limit: 1}
				returnResults := []*storage.ComplianceOperatorCheckResultV2{
					convertUtils.GetComplianceStorageResults(s.T())[0],
				}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).Return(returnResults, nil).Times(1)
				s.scanConfigDS.EXPECT().GetScanConfigurationByName(gomock.Any(), "scanConfig1").Return(getTestRec("scanConfig1"), nil).Times(1)
			},
		},
		{
			desc: "Query with no scan configuration name field",
			query: &apiV2.ComplianceScanResultsRequest{
				Query: &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Scan configuration name is required"),
			found:       false,
			setMocks:    func() {},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceScanConfigurationResults(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(convertUtils.GetConvertedComplianceResults(s.T()), results.GetScanResults())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileScanStats() {
	testCases := []struct {
		desc         string
		query        *apiV2.RawQuery
		expectedResp []*apiV2.ComplianceProfileScanStats
		expectedErr  error
		setMocks     func()
	}{
		{
			desc:        "Empty query",
			query:       &apiV2.RawQuery{Query: ""},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceProfileScanStats{
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "ocp4"),
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "rhcos4-moderate"),
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "ocp4-node"),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				results := []*datastore.ResourceResultCountByProfile{
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "ocp4"),
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "rhcos4-moderate"),
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "ocp4-node"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), search.EmptyQuery(), search.ComplianceOperatorProfileName)
				s.resultDatastore.EXPECT().ComplianceProfileResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				profiles_ocp := []*storage.ComplianceOperatorProfileV2{{
					Name:           "ocp4",
					ProfileVersion: "test_version_ocp4",
					Title:          "test_title_ocp4",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), search.NewQueryBuilder().
					AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").ProtoQuery()).Return(profiles_ocp, nil).Times(1)
				profiles := []*storage.ComplianceOperatorProfileV2{{
					Name:           "rhcos4-moderate",
					ProfileVersion: "test_version_rhcos4-moderate",
					Title:          "test_title_rhcos4-moderate",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), search.NewQueryBuilder().
					AddExactMatches(search.ComplianceOperatorProfileName, "rhcos4-moderate").ProtoQuery()).Return(profiles, nil).Times(1)
				profiles = []*storage.ComplianceOperatorProfileV2{{
					Name:           "ocp4-node",
					ProfileVersion: "test_version_ocp4-node",
					Title:          "test_title_ocp4-node",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), search.NewQueryBuilder().
					AddExactMatches(search.ComplianceOperatorProfileName, "ocp4-node").ProtoQuery()).Return(profiles, nil).Times(1)
			},
		},
		{
			desc:        "Query with search field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceProfileScanStats{
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "ocp4-node"),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				countQuery := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()

				results := []*datastore.ResourceResultCountByProfile{
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "ocp4-node"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorProfileName)
				s.resultDatastore.EXPECT().ComplianceProfileResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				profiles := []*storage.ComplianceOperatorProfileV2{{
					Name:           "ocp4-node",
					ProfileVersion: "test_version_ocp4-node",
					Title:          "test_title_ocp4-node",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), gomock.Any()).Return(profiles, nil).AnyTimes()
			},
		},
		{
			desc:        "Query with non-existent field",
			query:       &apiV2.RawQuery{Query: "Cluster ID:id"},
			expectedErr: errors.Wrapf(errox.InvalidArgs, "Unable to retrieve compliance cluster scan stats for query %v", &apiV2.RawQuery{Query: "Cluster ID:id"}),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").
					WithPagination(search.NewPagination().Limit(maxPaginationLimit)).ProtoQuery()
				countQuery := search.NewQueryBuilder().AddStrings(search.ClusterID, "id").ProtoQuery()

				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorProfileName)
				s.resultDatastore.EXPECT().ComplianceProfileResultStats(gomock.Any(), expectedQ).Return(nil, nil).Times(1)
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfilesStats(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().ElementsMatch(tc.expectedResp, results.GetScanStats())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileStats() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceProfileResultsRequest
		expectedResp []*apiV2.ComplianceProfileScanStats
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: ""},
			},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceProfileScanStats{
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "ocp4"),
			},
			setMocks: func() {

				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").ProtoQuery(),
					search.EmptyQuery(),
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultCountByProfile{
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "ocp4"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorProfileName)
				s.resultDatastore.EXPECT().ComplianceProfileResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				profiles_ocp := []*storage.ComplianceOperatorProfileV2{{
					Name:           "ocp4",
					ProfileVersion: "test_version_ocp4",
					Title:          "test_title_ocp4",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), search.NewQueryBuilder().
					AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").ProtoQuery()).Return(profiles_ocp, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "ocp4-node",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: nil,
			expectedResp: []*apiV2.ComplianceProfileScanStats{
				convertUtils.GetComplianceProfileScanV2Count(s.T(), "ocp4-node"),
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4-node").ProtoQuery(),
					expectedQ,
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultCountByProfile{
					convertUtils.GetComplianceStorageProfileScanCount(s.T(), "ocp4-node"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorProfileName)
				s.resultDatastore.EXPECT().ComplianceProfileResultStats(gomock.Any(), expectedQ).Return(results, nil).Times(1)
				profiles := []*storage.ComplianceOperatorProfileV2{{
					Name:           "ocp4-node",
					ProfileVersion: "test_version_ocp4-node",
					Title:          "test_title_ocp4-node",
				}}
				s.profilsDS.EXPECT().SearchProfiles(gomock.Any(), search.NewQueryBuilder().
					AddExactMatches(search.ComplianceOperatorProfileName, "ocp4-node").ProtoQuery()).Return(profiles, nil).Times(1)
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfileStats(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
			} else {
				s.Require().Error(tc.expectedErr, err)
			}

			if tc.expectedResp != nil {
				s.Require().Equal(tc.expectedResp, results.GetScanStats())
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileResults() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceProfileResultsRequest
		expectedResp *apiV2.ListComplianceProfileResults
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: ""},
			},
			expectedErr:  nil,
			expectedResp: convertUtils.GetComplianceProfileResultsV2(s.T(), "ocp4"),
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").ProtoQuery(),
					search.EmptyQuery(),
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultsByProfile{
					convertUtils.GetComplianceStorageProfileResults(s.T(), "ocp4"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorCheckName).Return(1, nil).Times(1)
				s.resultDatastore.EXPECT().ComplianceProfileResults(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr:  nil,
			expectedResp: convertUtils.GetComplianceProfileResultsV2(s.T(), "ocp4"),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").ProtoQuery(),
					expectedQ,
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultsByProfile{
					convertUtils.GetComplianceStorageProfileResults(s.T(), "ocp4"),
				}
				s.resultDatastore.EXPECT().CountByField(gomock.Any(), countQuery, search.ComplianceOperatorCheckName).Return(1, nil).Times(1)
				s.resultDatastore.EXPECT().ComplianceProfileResults(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileResultsRequest{
				ProfileName: "",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfileResults(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
				s.Require().Equal(tc.expectedResp, results)
			} else {
				s.Require().Error(tc.expectedErr, err)
				s.Require().Nil(results)
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileCheckStats() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceProfileCheckRequest
		expectedResp *apiV2.ListComplianceProfileResults
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				CheckName:   "check-name",
				Query:       &apiV2.RawQuery{Query: ""},
			},
			expectedErr:  nil,
			expectedResp: convertUtils.GetComplianceProfileResultsV2(s.T(), "ocp4"),
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ComplianceOperatorCheckName, "check-name").ProtoQuery(),
					search.EmptyQuery(),
				)
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultsByProfile{
					convertUtils.GetComplianceStorageProfileResults(s.T(), "ocp4"),
				}
				s.resultDatastore.EXPECT().ComplianceProfileResults(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				CheckName:   "check-name",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr:  nil,
			expectedResp: convertUtils.GetComplianceProfileResultsV2(s.T(), "ocp4"),
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ComplianceOperatorCheckName, "check-name").ProtoQuery(),
					expectedQ,
				)
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				results := []*datastore.ResourceResultsByProfile{
					convertUtils.GetComplianceStorageProfileResults(s.T(), "ocp4"),
				}
				s.resultDatastore.EXPECT().ComplianceProfileResults(gomock.Any(), expectedQ).Return(results, nil).Times(1)
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Compliance check name is required"),
			setMocks: func() {
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfileCheckStats(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
				s.Require().Equal(tc.expectedResp, results)
			} else {
				s.Require().Error(tc.expectedErr, err)
				s.Require().Nil(results)
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileCheckResult() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceProfileCheckRequest
		expectedResp *apiV2.ListComplianceCheckClusterResponse
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				CheckName:   "check-name",
				Query:       &apiV2.RawQuery{Query: ""},
			},
			expectedErr: nil,
			expectedResp: &apiV2.ListComplianceCheckClusterResponse{
				CheckResults: convertUtils.GetConvertedComplianceResult(s.T()).Clusters,
				ProfileName:  "ocp4",
				CheckName:    "check-name",
				TotalCount:   7,
			},
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ComplianceOperatorCheckName, "check-name").ProtoQuery(),
					search.EmptyQuery(),
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).
					Return([]*storage.ComplianceOperatorCheckResultV2{convertUtils.GetComplianceStorageResult(s.T())}, nil).
					Times(1)
				s.resultDatastore.EXPECT().CountCheckResults(gomock.Any(), countQuery).Return(7, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				CheckName:   "check-name",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: nil,
			expectedResp: &apiV2.ListComplianceCheckClusterResponse{
				CheckResults: convertUtils.GetConvertedComplianceResult(s.T()).Clusters,
				ProfileName:  "ocp4",
				CheckName:    "check-name",
				TotalCount:   3,
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ClusterID, fixtureconsts.Cluster1).ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ComplianceOperatorCheckName, "check-name").ProtoQuery(),
					expectedQ,
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).
					Return([]*storage.ComplianceOperatorCheckResultV2{convertUtils.GetComplianceStorageResult(s.T())}, nil).
					Times(1)
				s.resultDatastore.EXPECT().CountCheckResults(gomock.Any(), countQuery).Return(3, nil).Times(1)
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {
			},
		},
		{
			desc: "Query with non-existent field",
			query: &apiV2.ComplianceProfileCheckRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + fixtureconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Compliance check name is required"),
			setMocks: func() {
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfileCheckResult(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
				s.Require().Equal(tc.expectedResp, results)
			} else {
				s.Require().Error(tc.expectedErr, err)
				s.Require().Nil(results)
			}
		})
	}
}

func (s *ComplianceResultsServiceTestSuite) TestGetComplianceProfileClusterResults() {
	testCases := []struct {
		desc         string
		query        *apiV2.ComplianceProfileClusterRequest
		expectedResp *apiV2.ListComplianceCheckResultResponse
		expectedErr  error
		setMocks     func()
	}{
		{
			desc: "Empty query",
			query: &apiV2.ComplianceProfileClusterRequest{
				ProfileName: "ocp4",
				ClusterId:   testconsts.Cluster1,
				Query:       &apiV2.RawQuery{Query: ""},
			},
			expectedErr: nil,
			expectedResp: &apiV2.ListComplianceCheckResultResponse{
				CheckResults: convertUtils.GetConvertedCheckResult(s.T()),
				ProfileName:  "ocp4",
				ClusterId:    testconsts.Cluster1,
				TotalCount:   7,
			},
			setMocks: func() {
				expectedQ := search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ClusterID, testconsts.Cluster1).ProtoQuery(),
					search.EmptyQuery(),
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				ruleQuery := search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorRuleRef, "test-ref-id").ProtoQuery()
				s.ruleDS.EXPECT().SearchRules(gomock.Any(), ruleQuery).Return([]*storage.ComplianceOperatorRuleV2{{Name: "test-rule-name"}}, nil).Times(1)

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).
					Return([]*storage.ComplianceOperatorCheckResultV2{convertUtils.GetComplianceStorageResult(s.T())}, nil).
					Times(1)
				s.resultDatastore.EXPECT().CountCheckResults(gomock.Any(), countQuery).Return(7, nil).Times(1)
			},
		},
		{
			desc: "Query with search field",
			query: &apiV2.ComplianceProfileClusterRequest{
				ProfileName: "ocp4",
				ClusterId:   testconsts.Cluster1,
				Query:       &apiV2.RawQuery{Query: "Compliance Check Name:" + "check-name"},
			},
			expectedErr: nil,
			expectedResp: &apiV2.ListComplianceCheckResultResponse{
				CheckResults: convertUtils.GetConvertedCheckResult(s.T()),
				ProfileName:  "ocp4",
				ClusterId:    testconsts.Cluster1,
				TotalCount:   3,
			},
			setMocks: func() {
				expectedQ := search.NewQueryBuilder().AddStrings(search.ComplianceOperatorCheckName, "check-name").ProtoQuery()
				expectedQ = search.ConjunctionQuery(
					search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorProfileName, "ocp4").
						AddExactMatches(search.ClusterID, testconsts.Cluster1).ProtoQuery(),
					expectedQ,
				)
				countQuery := expectedQ.Clone()
				expectedQ.Pagination = &v1.QueryPagination{Limit: maxPaginationLimit}

				ruleQuery := search.NewQueryBuilder().AddExactMatches(search.ComplianceOperatorRuleRef, "test-ref-id").ProtoQuery()
				s.ruleDS.EXPECT().SearchRules(gomock.Any(), ruleQuery).Return([]*storage.ComplianceOperatorRuleV2{{Name: "test-rule-name"}}, nil).Times(1)

				s.resultDatastore.EXPECT().SearchComplianceCheckResults(gomock.Any(), expectedQ).
					Return([]*storage.ComplianceOperatorCheckResultV2{convertUtils.GetComplianceStorageResult(s.T())}, nil).
					Times(1)
				s.resultDatastore.EXPECT().CountCheckResults(gomock.Any(), countQuery).Return(3, nil).Times(1)
			},
		},
		{
			desc: "Request with missing field",
			query: &apiV2.ComplianceProfileClusterRequest{
				ProfileName: "",
				Query:       &apiV2.RawQuery{Query: "Cluster ID:" + testconsts.Cluster1},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Profile name is required"),
			setMocks: func() {
			},
		},
		{
			desc: "Query with missing cluster",
			query: &apiV2.ComplianceProfileClusterRequest{
				ProfileName: "ocp4",
				Query:       &apiV2.RawQuery{Query: "Compliance Operator Check Name:" + "check-name"},
			},
			expectedErr: errors.Wrap(errox.InvalidArgs, "Cluster ID is required"),
			setMocks: func() {
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.desc, func(t *testing.T) {
			tc.setMocks()

			results, err := s.service.GetComplianceProfileClusterResults(s.ctx, tc.query)
			if tc.expectedErr == nil {
				s.Require().NoError(err)
				s.Require().Equal(tc.expectedResp, results)
			} else {
				s.Require().Error(tc.expectedErr, err)
				s.Require().Nil(results)
			}
		})
	}
}

func getTestRec(scanName string) *storage.ComplianceOperatorScanConfigurationV2 {
	return &storage.ComplianceOperatorScanConfigurationV2{
		Id:                     scanName,
		ScanConfigName:         scanName,
		AutoApplyRemediations:  false,
		AutoUpdateRemediations: false,
		OneTimeScan:            false,
		Profiles: []*storage.ComplianceOperatorScanConfigurationV2_ProfileName{
			{
				ProfileName: "ocp4-cis",
			},
		},
		StrictNodeScan: false,
		Description:    "test-description",
		Clusters: []*storage.ComplianceOperatorScanConfigurationV2_Cluster{
			{
				ClusterId: fixtureconsts.Cluster1,
			},
			{
				ClusterId: fixtureconsts.Cluster2,
			},
		},
	}
}
