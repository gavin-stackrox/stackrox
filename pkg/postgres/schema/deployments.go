// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lib/pq"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableDeploymentsStmt holds the create statement for table `deployments`.
	CreateTableDeploymentsStmt = &postgres.CreateStmts{
		GormModel: (*Deployments)(nil),
		Children: []*postgres.CreateStmts{
			&postgres.CreateStmts{
				GormModel: (*DeploymentsContainers)(nil),
				Children: []*postgres.CreateStmts{
					&postgres.CreateStmts{
						GormModel: (*DeploymentsContainersEnvs)(nil),
						Children:  []*postgres.CreateStmts{},
					},
					&postgres.CreateStmts{
						GormModel: (*DeploymentsContainersVolumes)(nil),
						Children:  []*postgres.CreateStmts{},
					},
					&postgres.CreateStmts{
						GormModel: (*DeploymentsContainersSecrets)(nil),
						Children:  []*postgres.CreateStmts{},
					},
				},
			},
			&postgres.CreateStmts{
				GormModel: (*DeploymentsPorts)(nil),
				Children: []*postgres.CreateStmts{
					&postgres.CreateStmts{
						GormModel: (*DeploymentsPortsExposureInfos)(nil),
						Children:  []*postgres.CreateStmts{},
					},
				},
			},
		},
	}

	// DeploymentsSchema is the go schema for table `deployments`.
	DeploymentsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("deployments")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.Deployment)(nil)), "deployments")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Image":             ImagesSchema,
			"storage.NamespaceMetadata": NamespacesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_DEPLOYMENTS, "deployment", (*storage.Deployment)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_IMAGE_VULNERABILITIES,
			v1.SearchCategory_COMPONENT_VULN_EDGE,
			v1.SearchCategory_IMAGE_COMPONENTS,
			v1.SearchCategory_IMAGE_COMPONENT_EDGE,
			v1.SearchCategory_IMAGE_VULN_EDGE,
			v1.SearchCategory_IMAGES,
			v1.SearchCategory_DEPLOYMENTS,
			v1.SearchCategory_NAMESPACES,
			v1.SearchCategory_CLUSTERS,
			v1.SearchCategory_PROCESS_INDICATORS,
		}...)
		RegisterTable(schema, CreateTableDeploymentsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_DEPLOYMENTS, schema)
		return schema
	}()
)

const (
	DeploymentsTableName                   = "deployments"
	DeploymentsContainersTableName         = "deployments_containers"
	DeploymentsContainersEnvsTableName     = "deployments_containers_envs"
	DeploymentsContainersVolumesTableName  = "deployments_containers_volumes"
	DeploymentsContainersSecretsTableName  = "deployments_containers_secrets"
	DeploymentsPortsTableName              = "deployments_ports"
	DeploymentsPortsExposureInfosTableName = "deployments_ports_exposure_infos"
)

// Deployments holds the Gorm model for Postgres table `deployments`.
type Deployments struct {
	Id                            string                  `gorm:"column:id;type:uuid;primaryKey"`
	Name                          string                  `gorm:"column:name;type:varchar"`
	Type                          string                  `gorm:"column:type;type:varchar"`
	Namespace                     string                  `gorm:"column:namespace;type:varchar;index:deployments_sac_filter,type:btree"`
	NamespaceId                   string                  `gorm:"column:namespaceid;type:uuid"`
	OrchestratorComponent         bool                    `gorm:"column:orchestratorcomponent;type:bool"`
	Labels                        map[string]string       `gorm:"column:labels;type:jsonb"`
	PodLabels                     map[string]string       `gorm:"column:podlabels;type:jsonb"`
	Created                       *time.Time              `gorm:"column:created;type:timestamp"`
	ClusterId                     string                  `gorm:"column:clusterid;type:uuid;index:deployments_sac_filter,type:btree"`
	ClusterName                   string                  `gorm:"column:clustername;type:varchar"`
	Annotations                   map[string]string       `gorm:"column:annotations;type:jsonb"`
	Priority                      int64                   `gorm:"column:priority;type:bigint"`
	ImagePullSecrets              *pq.StringArray         `gorm:"column:imagepullsecrets;type:text[]"`
	ServiceAccount                string                  `gorm:"column:serviceaccount;type:varchar"`
	ServiceAccountPermissionLevel storage.PermissionLevel `gorm:"column:serviceaccountpermissionlevel;type:integer"`
	RiskScore                     float32                 `gorm:"column:riskscore;type:numeric"`
	Serialized                    []byte                  `gorm:"column:serialized;type:bytea"`
}

// DeploymentsContainers holds the Gorm model for Postgres table `deployments_containers`.
type DeploymentsContainers struct {
	DeploymentsId                         string          `gorm:"column:deployments_id;type:uuid;primaryKey"`
	Idx                                   int             `gorm:"column:idx;type:integer;primaryKey;index:deploymentscontainers_idx,type:btree"`
	ImageId                               string          `gorm:"column:image_id;type:varchar;index:deploymentscontainers_image_id,type:hash"`
	ImageNameRegistry                     string          `gorm:"column:image_name_registry;type:varchar"`
	ImageNameRemote                       string          `gorm:"column:image_name_remote;type:varchar"`
	ImageNameTag                          string          `gorm:"column:image_name_tag;type:varchar"`
	ImageNameFullName                     string          `gorm:"column:image_name_fullname;type:varchar"`
	SecurityContextPrivileged             bool            `gorm:"column:securitycontext_privileged;type:bool"`
	SecurityContextDropCapabilities       *pq.StringArray `gorm:"column:securitycontext_dropcapabilities;type:text[]"`
	SecurityContextAddCapabilities        *pq.StringArray `gorm:"column:securitycontext_addcapabilities;type:text[]"`
	SecurityContextReadOnlyRootFilesystem bool            `gorm:"column:securitycontext_readonlyrootfilesystem;type:bool"`
	ResourcesCpuCoresRequest              float32         `gorm:"column:resources_cpucoresrequest;type:numeric"`
	ResourcesCpuCoresLimit                float32         `gorm:"column:resources_cpucoreslimit;type:numeric"`
	ResourcesMemoryMbRequest              float32         `gorm:"column:resources_memorymbrequest;type:numeric"`
	ResourcesMemoryMbLimit                float32         `gorm:"column:resources_memorymblimit;type:numeric"`
	DeploymentsRef                        Deployments     `gorm:"foreignKey:deployments_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}

// DeploymentsContainersEnvs holds the Gorm model for Postgres table `deployments_containers_envs`.
type DeploymentsContainersEnvs struct {
	DeploymentsId            string                                                 `gorm:"column:deployments_id;type:uuid;primaryKey"`
	DeploymentsContainersIdx int                                                    `gorm:"column:deployments_containers_idx;type:integer;primaryKey"`
	Idx                      int                                                    `gorm:"column:idx;type:integer;primaryKey;index:deploymentscontainersenvs_idx,type:btree"`
	Key                      string                                                 `gorm:"column:key;type:varchar"`
	Value                    string                                                 `gorm:"column:value;type:varchar"`
	EnvVarSource             storage.ContainerConfig_EnvironmentConfig_EnvVarSource `gorm:"column:envvarsource;type:integer"`
	DeploymentsContainersRef DeploymentsContainers                                  `gorm:"foreignKey:deployments_id,deployments_containers_idx;references:deployments_id,idx;belongsTo;constraint:OnDelete:CASCADE"`
}

// DeploymentsContainersVolumes holds the Gorm model for Postgres table `deployments_containers_volumes`.
type DeploymentsContainersVolumes struct {
	DeploymentsId            string                `gorm:"column:deployments_id;type:uuid;primaryKey"`
	DeploymentsContainersIdx int                   `gorm:"column:deployments_containers_idx;type:integer;primaryKey"`
	Idx                      int                   `gorm:"column:idx;type:integer;primaryKey;index:deploymentscontainersvolumes_idx,type:btree"`
	Name                     string                `gorm:"column:name;type:varchar"`
	Source                   string                `gorm:"column:source;type:varchar"`
	Destination              string                `gorm:"column:destination;type:varchar"`
	ReadOnly                 bool                  `gorm:"column:readonly;type:bool"`
	Type                     string                `gorm:"column:type;type:varchar"`
	DeploymentsContainersRef DeploymentsContainers `gorm:"foreignKey:deployments_id,deployments_containers_idx;references:deployments_id,idx;belongsTo;constraint:OnDelete:CASCADE"`
}

// DeploymentsContainersSecrets holds the Gorm model for Postgres table `deployments_containers_secrets`.
type DeploymentsContainersSecrets struct {
	DeploymentsId            string                `gorm:"column:deployments_id;type:uuid;primaryKey"`
	DeploymentsContainersIdx int                   `gorm:"column:deployments_containers_idx;type:integer;primaryKey"`
	Idx                      int                   `gorm:"column:idx;type:integer;primaryKey;index:deploymentscontainerssecrets_idx,type:btree"`
	Name                     string                `gorm:"column:name;type:varchar"`
	Path                     string                `gorm:"column:path;type:varchar"`
	DeploymentsContainersRef DeploymentsContainers `gorm:"foreignKey:deployments_id,deployments_containers_idx;references:deployments_id,idx;belongsTo;constraint:OnDelete:CASCADE"`
}

// DeploymentsPorts holds the Gorm model for Postgres table `deployments_ports`.
type DeploymentsPorts struct {
	DeploymentsId  string                           `gorm:"column:deployments_id;type:uuid;primaryKey"`
	Idx            int                              `gorm:"column:idx;type:integer;primaryKey;index:deploymentsports_idx,type:btree"`
	ContainerPort  int32                            `gorm:"column:containerport;type:integer"`
	Protocol       string                           `gorm:"column:protocol;type:varchar"`
	Exposure       storage.PortConfig_ExposureLevel `gorm:"column:exposure;type:integer"`
	DeploymentsRef Deployments                      `gorm:"foreignKey:deployments_id;references:id;belongsTo;constraint:OnDelete:CASCADE"`
}

// DeploymentsPortsExposureInfos holds the Gorm model for Postgres table `deployments_ports_exposure_infos`.
type DeploymentsPortsExposureInfos struct {
	DeploymentsId       string                           `gorm:"column:deployments_id;type:uuid;primaryKey"`
	DeploymentsPortsIdx int                              `gorm:"column:deployments_ports_idx;type:integer;primaryKey"`
	Idx                 int                              `gorm:"column:idx;type:integer;primaryKey;index:deploymentsportsexposureinfos_idx,type:btree"`
	Level               storage.PortConfig_ExposureLevel `gorm:"column:level;type:integer"`
	ServiceName         string                           `gorm:"column:servicename;type:varchar"`
	ServicePort         int32                            `gorm:"column:serviceport;type:integer"`
	NodePort            int32                            `gorm:"column:nodeport;type:integer"`
	ExternalIps         *pq.StringArray                  `gorm:"column:externalips;type:text[]"`
	ExternalHostnames   *pq.StringArray                  `gorm:"column:externalhostnames;type:text[]"`
	DeploymentsPortsRef DeploymentsPorts                 `gorm:"foreignKey:deployments_id,deployments_ports_idx;references:deployments_id,idx;belongsTo;constraint:OnDelete:CASCADE"`
}
