import React, { ReactElement } from 'react';
import { Redirect, Switch } from 'react-router-dom';

import {
    mainPath,
    dashboardPath,
    networkPath,
    violationsPath,
    compliancePath,
    clustersPathWithParam,
    clustersListPath,
    integrationsPath,
    policiesPath,
    policiesPathPatternFly,
    riskPath,
    apidocsPath,
    accessControlPathV2,
    userBasePath,
    systemConfigPath,
    systemHealthPath,
    systemHealthPathPF,
    vulnManagementPath,
    vulnManagementReportingPath,
    configManagementPath,
} from 'routePaths';
import { useTheme } from 'Containers/ThemeProvider';

import asyncComponent from 'Components/AsyncComponent';
import ProtectedRoute from 'Components/ProtectedRoute';
import ErrorBoundary from 'Containers/ErrorBoundary';
import { knownBackendFlags } from 'utils/featureFlags';
import useFeatureFlagEnabled from 'hooks/useFeatureFlagEnabled';

const AsyncApiDocsPage = asyncComponent(() => import('Containers/Docs/ApiPage'));
const AsyncDashboardPage = asyncComponent(() => import('Containers/Dashboard/DashboardPage'));
const AsyncNetworkPage = asyncComponent(() => import('Containers/Network/Page'));
const AsyncClustersPage = asyncComponent(() => import('Containers/Clusters/ClustersPage'));
const AsyncPFClustersPage = asyncComponent(() => import('Containers/Clusters/PF/ClustersPage'));
const AsyncIntegrationsPage = asyncComponent(
    () => import('Containers/Integrations/IntegrationsPage')
);
const AsyncViolationsPage = asyncComponent(
    () => import('Containers/Violations/PatternFly/ViolationsPage')
);

const AsyncPoliciesPage = asyncComponent(() => import('Containers/Policies/Page'));
const AsyncPoliciesPagePatternFly = asyncComponent(
    () => import('Containers/Policies/PatternFly/PoliciesPage')
);
const AsyncCompliancePage = asyncComponent(() => import('Containers/Compliance/Page'));
const AsyncRiskPage = asyncComponent(() => import('Containers/Risk/RiskPage'));
const AsyncAccessControlPageV2 = asyncComponent(
    () => import('Containers/AccessControl/AccessControl')
);
const AsyncUserPage = asyncComponent(() => import('Containers/User/UserPage'));
const AsyncSystemConfigPage = asyncComponent(() => import('Containers/SystemConfig/Page'));
const AsyncConfigManagementPage = asyncComponent(() => import('Containers/ConfigManagement/Page'));
const AsyncVulnMgmtReportingPage = asyncComponent(
    () => import('Containers/VulnMgmt/Reporting/VulnMgmtReportPage')
);
const AsyncVulnMgmtPage = asyncComponent(() => import('Containers/Workflow/WorkflowLayout'));
const AsyncSystemHealthPage = asyncComponent(() => import('Containers/SystemHealth/DashboardPage'));
const AsyncSystemHealthPagePF = asyncComponent(
    () => import('Containers/SystemHealth/PatternFly/SystemHealthDashboard')
);

function Body(): ReactElement {
    const { isDarkMode } = useTheme();
    const isSystemHealthPatternFlyEnabled = useFeatureFlagEnabled(
        knownBackendFlags.ROX_SYSTEM_HEALTH_PF
    );
    const isPoliciesPatternFlyEnabled = useFeatureFlagEnabled(
        knownBackendFlags.ROX_POLICIES_PATTERNFLY
    );
    return (
        <div
            className={`flex flex-col h-full w-full relative overflow-auto ${
                isDarkMode ? 'bg-base-0' : 'bg-base-100'
            }`}
        >
            <ErrorBoundary>
                <Switch>
                    <ProtectedRoute path={dashboardPath} component={AsyncDashboardPage} />
                    <ProtectedRoute path={networkPath} component={AsyncNetworkPage} />
                    <ProtectedRoute path={violationsPath} component={AsyncViolationsPage} />
                    <ProtectedRoute path={compliancePath} component={AsyncCompliancePage} />
                    <ProtectedRoute path={integrationsPath} component={AsyncIntegrationsPage} />
                    <ProtectedRoute path={policiesPath} component={AsyncPoliciesPage} />
                    <ProtectedRoute
                        path={policiesPathPatternFly}
                        component={AsyncPoliciesPagePatternFly}
                        featureFlagEnabled={isPoliciesPatternFlyEnabled}
                    />
                    <ProtectedRoute path={riskPath} component={AsyncRiskPage} />
                    <ProtectedRoute
                        path={accessControlPathV2}
                        component={AsyncAccessControlPageV2}
                    />
                    <ProtectedRoute path={apidocsPath} component={AsyncApiDocsPage} />
                    <ProtectedRoute path={userBasePath} component={AsyncUserPage} />
                    <ProtectedRoute path={systemConfigPath} component={AsyncSystemConfigPage} />
                    <ProtectedRoute
                        path={vulnManagementReportingPath}
                        component={AsyncVulnMgmtReportingPage}
                    />
                    <ProtectedRoute path={vulnManagementPath} component={AsyncVulnMgmtPage} />
                    <ProtectedRoute
                        path={configManagementPath}
                        component={AsyncConfigManagementPage}
                    />
                    <ProtectedRoute path={clustersPathWithParam} component={AsyncClustersPage} />
                    {process.env.NODE_ENV === 'development' && (
                        <ProtectedRoute path={clustersListPath} component={AsyncPFClustersPage} />
                    )}
                    <ProtectedRoute path={systemHealthPath} component={AsyncSystemHealthPage} />
                    <ProtectedRoute
                        path={systemHealthPathPF}
                        component={AsyncSystemHealthPagePF}
                        featureFlagEnabled={isSystemHealthPatternFlyEnabled}
                    />
                    <Redirect from={mainPath} to={dashboardPath} />
                </Switch>
            </ErrorBoundary>
        </div>
    );
}

export default Body;
