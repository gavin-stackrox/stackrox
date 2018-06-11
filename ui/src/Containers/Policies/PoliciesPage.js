import React, { Component } from 'react';
import PropTypes from 'prop-types';
import ReactRouterPropTypes from 'react-router-prop-types';
import { connect } from 'react-redux';
import { selectors } from 'reducers';
import { actions as policyActions } from 'reducers/policies';
import { createSelector, createStructuredSelector } from 'reselect';

import { formValueSelector } from 'redux-form';
import * as Icon from 'react-feather';
import Dialog from 'Components/Dialog';
import Table from 'Components/Table';
import Panel from 'Components/Panel';
import { formatPolicyFields, getPolicyFormDataKeys } from 'Containers/Policies/policyFormUtils';
import { deletePolicy } from 'services/PoliciesService';
import PolicyDetails from 'Containers/Policies/PolicyDetails';
import PageHeader from 'Components/PageHeader';
import SearchInput from 'Components/SearchInput';

import { severityLabels } from 'messages/common';
import { sortSeverity } from 'sorters/sorters';
import PolicyCreationWizard from 'Containers/Policies/PolicyCreationWizard';

class PoliciesPage extends Component {
    static propTypes = {
        policies: PropTypes.arrayOf(PropTypes.object).isRequired,
        selectedPolicy: PropTypes.shape({
            id: PropTypes.string.isRequired
        }),
        fetchPolicies: PropTypes.func.isRequired,
        reassessPolicies: PropTypes.func.isRequired,
        updatePolicy: PropTypes.func.isRequired,
        formData: PropTypes.shape({}),
        wizardState: PropTypes.shape({
            current: PropTypes.string,
            policy: PropTypes.shape({}),
            isNew: PropTypes.bool,
            disabled: PropTypes.bool
        }).isRequired,
        setWizardState: PropTypes.func.isRequired,
        history: ReactRouterPropTypes.history.isRequired,
        match: ReactRouterPropTypes.match.isRequired,
        searchOptions: PropTypes.arrayOf(PropTypes.object).isRequired,
        searchModifiers: PropTypes.arrayOf(PropTypes.object).isRequired,
        searchSuggestions: PropTypes.arrayOf(PropTypes.object).isRequired,
        setSearchOptions: PropTypes.func.isRequired,
        setSearchModifiers: PropTypes.func.isRequired,
        setSearchSuggestions: PropTypes.func.isRequired,
        isViewFiltered: PropTypes.bool.isRequired
    };

    static defaultProps = {
        selectedPolicy: null,
        formData: {}
    };

    constructor(props) {
        super(props);

        this.state = {
            showConfirmationDialog: false
        };
    }

    componentWillUnmount() {
        this.props.setWizardState({ current: '' });
    }

    onSubmit = () => {
        const { selectedPolicy } = this.props;
        const { isNew, policy, disabled } = this.props.wizardState;
        const newPolicy = Object.assign({}, selectedPolicy, policy);
        const newState = {};
        newState.current = isNew ? 'CREATE' : 'SAVE';
        newState.policy = newPolicy;
        if (disabled) newState.policy.disabled = disabled;
        this.props.setWizardState(newState);
    };

    getPanelButtons = () => {
        switch (this.props.wizardState.current) {
            case 'EDIT':
            case 'PRE_PREVIEW':
                return [
                    {
                        renderIcon: () => <Icon.ArrowRight className="h-4 w-4" />,
                        text: 'Next',
                        className: 'btn-primary',
                        onClick: () => {
                            this.getPolicyDryRun();
                        }
                    }
                ];
            case 'PREVIEW':
                return [
                    {
                        renderIcon: () => <Icon.ArrowLeft className="h-4 w-4" />,
                        text: 'Previous',
                        className: 'btn-primary',
                        onClick: () => {
                            this.props.setWizardState({ current: 'EDIT' });
                        }
                    },
                    {
                        renderIcon: () => <Icon.Save className="h-4 w-4" />,
                        text: 'Save',
                        className: 'btn-success',
                        onClick: () => {
                            this.onSubmit();
                        }
                    }
                ];
            default:
                return [
                    {
                        renderIcon: () => <Icon.Edit className="h-4 w-4" />,
                        text: 'Edit',
                        className: 'btn-success',
                        onClick: () => {
                            this.props.setWizardState({ current: 'EDIT', policy: null });
                        }
                    }
                ];
        }
    };

    getPolicyDryRun = () => {
        const serverFormattedPolicy = formatPolicyFields(this.props.formData);
        const enabledPolicy = Object.assign({}, serverFormattedPolicy);
        // set disabled to false for dryrun so that we can see what deployments the policy will affect
        enabledPolicy.disabled = false;

        const wizardState = {
            current: 'PRE_PREVIEW',
            policy: enabledPolicy,
            disabled: serverFormattedPolicy.disabled
        };
        this.props.setWizardState(wizardState);
    };

    setSelectedPolicy = policy => {
        const urlSuffix = policy && policy.id ? `/${policy.id}` : '';
        this.props.history.push({
            pathname: `/main/policies${urlSuffix}`
        });
        this.props.setWizardState({ current: '', isNew: false });
    };

    preSubmit = policy => {
        const newPolicy = formatPolicyFields(policy);
        return newPolicy;
    };

    deletePolicies = () => {
        const promises = [];
        this.policyTable.getSelectedRows().forEach(row => {
            // close the view panel if that policy is being deleted
            if (row.id === this.props.match.params.id) {
                this.setSelectedPolicy();
            }
            const promise = deletePolicy(row.id);
            promises.push(promise);
        });
        Promise.all(promises).then(() => {
            this.policyTable.clearSelectedRows();
            this.hideConfirmationDialog();
            this.props.fetchPolicies();
        });
    };

    addPolicy = () => {
        this.setSelectedPolicy();
        this.props.setWizardState({ current: 'EDIT', policy: null, isNew: true });
    };

    toggleEnabledDisabledPolicy = policy => {
        this.props.updatePolicy({ ...policy, disabled: !policy.disabled });
    };

    showConfirmationDialog = () => {
        this.setState({ showConfirmationDialog: true });
    };

    hideConfirmationDialog = () => {
        this.setState({ showConfirmationDialog: false });
    };

    renderTablePanel = () => {
        const header = `${this.props.policies.length} Policies`;
        const buttons = [
            {
                renderIcon: () => <Icon.Trash2 className="h-4 w-4" />,
                text: 'Delete Policies',
                className: 'btn-danger',
                onClick: this.showConfirmationDialog,
                disabled: this.props.wizardState.current !== ''
            },
            {
                renderIcon: () => <Icon.FileText className="h-4 w-4" />,
                text: 'Reassess Policies',
                className: 'btn-success',
                onClick: this.props.reassessPolicies,
                disabled: this.props.wizardState.current !== '',
                tooltip: 'Manually enrich external data'
            },
            {
                renderIcon: () => <Icon.Plus className="h-4 w-4" />,
                text: 'Add',
                className: 'btn-success',
                onClick: this.addPolicy,
                disabled: this.props.wizardState.current !== ''
            }
        ];
        const columns = [
            {
                key: 'name',
                keys: ['name', 'disabled'],
                keyValueFunc: (name, disabled) => (
                    <div className="flex items-center relative">
                        <div
                            className={`h-2 w-2 rounded-lg absolute -ml-4 ${
                                !disabled ? 'bg-success-500' : 'bg-base-300'
                            }`}
                        />
                        <div>{name}</div>
                    </div>
                ),
                label: 'Name'
            },
            { key: 'description', label: 'Description' },
            {
                key: 'severity',
                label: 'Severity',
                keyValueFunc: severity => severityLabels[severity],
                classFunc: severity => {
                    switch (severity) {
                        case 'Low':
                            return 'text-low-500';
                        case 'Medium':
                            return 'text-medium-500';
                        case 'High':
                            return 'text-high-severity';
                        case 'Critical':
                            return 'text-critical-severity';
                        default:
                            return '';
                    }
                },
                sortMethod: sortSeverity
            }
        ];
        const actions = [
            {
                renderIcon: row =>
                    row.disabled ? (
                        <Icon.Power className="h-5 w-4 text-base-600" />
                    ) : (
                        <Icon.Power className="h-5 w-4 text-success-500" />
                    ),
                className: 'flex rounded-sm uppercase text-center text-sm items-center',
                onClick: this.toggleEnabledDisabledPolicy
            }
        ];
        const rows = this.props.policies;
        return (
            <Panel header={header} buttons={buttons}>
                <Table
                    columns={columns}
                    rows={rows}
                    onRowClick={this.setSelectedPolicy}
                    actions={actions}
                    checkboxes
                    ref={table => {
                        this.policyTable = table;
                    }}
                />
            </Panel>
        );
    };

    renderSidePanel = () => {
        const { selectedPolicy } = this.props;
        if (!this.props.wizardState.current && !selectedPolicy) return null;

        const editingPolicy = Object.assign({}, selectedPolicy, this.props.wizardState.policy);
        const header = editingPolicy ? editingPolicy.name : '';
        const buttons = this.getPanelButtons();
        return (
            <Panel header={header} buttons={buttons} onClose={this.setSelectedPolicy} width="w-2/3">
                {this.props.wizardState.current === '' ? (
                    <PolicyDetails policyId={selectedPolicy.id} />
                ) : (
                    <PolicyCreationWizard />
                )}
            </Panel>
        );
    };

    renderConfirmationDialog = () => {
        const numSelectedRows = this.policyTable ? this.policyTable.getSelectedRows().length : 0;
        return (
            <Dialog
                isOpen={this.state.showConfirmationDialog}
                text={`Are you sure you want to delete ${numSelectedRows} ${
                    numSelectedRows === 1 ? 'policy' : 'policies'
                }?`}
                onConfirm={this.deletePolicies}
                onCancel={this.hideConfirmationDialog}
            />
        );
    };

    render() {
        const subHeader = this.props.isViewFiltered ? 'Filtered view' : 'Default view';
        return (
            <section className="flex flex-1 flex-col h-full">
                <div>
                    <PageHeader header="Policies" subHeader={subHeader}>
                        <SearchInput
                            id="risk"
                            searchOptions={this.props.searchOptions}
                            searchModifiers={this.props.searchModifiers}
                            searchSuggestions={this.props.searchSuggestions}
                            setSearchOptions={this.props.setSearchOptions}
                            setSearchModifiers={this.props.setSearchModifiers}
                            setSearchSuggestions={this.props.setSearchSuggestions}
                        />
                    </PageHeader>
                </div>
                <div className="flex flex-1 bg-base-100">
                    <div className="flex flex-row w-full h-full bg-white rounded-sm shadow">
                        {this.renderTablePanel()}
                        {this.renderSidePanel()}
                    </div>
                </div>
                {this.renderConfirmationDialog()}
            </section>
        );
    }
}

const isViewFiltered = createSelector(
    [selectors.getPoliciesSearchOptions],
    searchOptions => searchOptions.length !== 0
);

const getFormData = state =>
    formValueSelector('policyCreationForm')(state, ...getPolicyFormDataKeys());

const getSelectedPolicy = (state, props) => {
    const { policyId } = props.match.params;
    return policyId ? selectors.getPolicy(state, policyId) : null;
};

const mapStateToProps = createStructuredSelector({
    policies: state => Object.values(selectors.getPoliciesById(state)),
    selectedPolicy: getSelectedPolicy,
    formData: getFormData,
    wizardState: selectors.getPolicyWizardState,
    searchOptions: selectors.getPoliciesSearchOptions,
    searchModifiers: selectors.getPoliciesSearchModifiers,
    searchSuggestions: selectors.getPoliciesSearchSuggestions,
    isViewFiltered
});

const mapDispatchToProps = {
    setSearchOptions: policyActions.setPoliciesSearchOptions,
    setSearchModifiers: policyActions.setPoliciesSearchModifiers,
    setSearchSuggestions: policyActions.setPoliciesSearchSuggestions,
    fetchPolicies: policyActions.fetchPolicies.request,
    reassessPolicies: policyActions.reassessPolicies,
    updatePolicy: policyActions.updatePolicy,
    setWizardState: policyActions.setPolicyWizardState
};

export default connect(mapStateToProps, mapDispatchToProps)(PoliciesPage);
