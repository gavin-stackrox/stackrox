import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import { Redirect, Switch } from 'react-router-dom';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import {
    mainPath,
    dashboardPath,
    violationsPath,
    compliancePath,
    integrationsPath,
    policiesPath,
    riskPath,
    imagesPath
} from 'routePaths';
import { selectors } from 'reducers';
import { actions as globalSearchActions } from 'reducers/globalSearch';

import ProtectedRoute from 'Components/ProtectedRoute';
import Notifications from 'Containers/Notifications';
import DashboardPage from 'Containers/Dashboard/DashboardPage';
import IntegrationsPage from 'Containers/Integrations/IntegrationsPage';
import ViolationsPage from 'Containers/Violations/ViolationsPage';
import PoliciesPage from 'Containers/Policies/PoliciesPage';
import ImagesPage from 'Containers/Images/ImagesPage';
import CompliancePage from 'Containers/Compliance/CompliancePage';
import RiskPage from 'Containers/Risk/RiskPage';
import TopNavigation from 'Containers/Navigation/TopNavigation';
import LeftNavigation from 'Containers/Navigation/LeftNavigation';
import SearchModal from 'Containers/Search/SearchModal';

class MainPage extends Component {
    static propTypes = {
        history: ReactRouterPropTypes.history.isRequired,
        toggleGlobalSearchView: PropTypes.func.isRequired,
        globalSearchView: PropTypes.bool.isRequired
    };

    onCloseHandler = toURL => {
        this.props.toggleGlobalSearchView();
        if (toURL && typeof toURL === 'string') this.props.history.push(toURL);
    };

    renderSearchModal = () => {
        if (!this.props.globalSearchView) return '';
        return <SearchModal className="h-full w-full" onClose={this.onCloseHandler} />;
    };

    renderRouter = () => (
        <section className="flex-auto overflow-auto">
            <Switch>
                <ProtectedRoute path={dashboardPath} component={DashboardPage} />
                <ProtectedRoute path={violationsPath} component={ViolationsPage} />
                <ProtectedRoute path={compliancePath} component={CompliancePage} />
                <ProtectedRoute path={integrationsPath} component={IntegrationsPage} />
                <ProtectedRoute path={policiesPath} component={PoliciesPage} />
                <ProtectedRoute path={riskPath} component={RiskPage} />
                <ProtectedRoute path={imagesPath} component={ImagesPage} />
                <Redirect from={mainPath} to={dashboardPath} />
            </Switch>
        </section>
    );

    render() {
        return (
            <section className="flex flex-1 flex-col h-full relative">
                <Notifications />
                <div className="navigation-gradient" />
                <header className="flex">
                    <TopNavigation />
                </header>
                <section className="flex flex-1 flex-row">
                    <LeftNavigation />
                    {this.renderRouter()}
                </section>
                {this.renderSearchModal()}
            </section>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    globalSearchView: selectors.getGlobalSearchView
});

const mapDispatchToProps = dispatch => ({
    toggleGlobalSearchView: () => dispatch(globalSearchActions.toggleGlobalSearchView())
});

export default connect(mapStateToProps, mapDispatchToProps)(MainPage);
