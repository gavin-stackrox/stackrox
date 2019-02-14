import React from 'react';
import ComplianceByStandard from 'Containers/Compliance/widgets/ComplianceByStandard';
import PropTypes from 'prop-types';
import entityTypes from 'constants/entityTypes';
import EntityCompliance from 'Containers/Compliance/widgets/EntityCompliance';
import Query from 'Components/ThrowingQuery';
import Labels from 'Containers/Compliance/widgets/Labels';
import IconWidget from 'Components/IconWidget';
import CountWidget from 'Components/CountWidget';
import pluralize from 'pluralize';
import Cluster from 'images/cluster.svg';
import { NAMESPACE_QUERY as QUERY } from 'queries/namespace';
import Widget from 'Components/Widget';
import ResourceRelatedResourceList from 'Containers/Compliance/widgets/ResourceRelatedResourceList';
import Header from './Header';

function processData(data) {
    const defaultValue = {
        labels: []
    };

    if (!data || !data.results || !data.results.metadata) return defaultValue;

    const { metadata, ...rest } = data.results;

    return {
        ...rest,
        ...metadata
    };
}

const NamespacePage = ({ sidePanelMode, params }) => (
    <Query query={QUERY} variables={{ id: params.entityId }} pollInterval={5000}>
        {({ loading, data }) => {
            const namespace = processData(data);
            const header = namespace.name || 'Loading';
            return (
                <section className="flex flex-col h-full w-full">
                    {!sidePanelMode && (
                        <Header header={header} subHeader="Namespace" params={params} />
                    )}
                    <div
                        className={`flex-1 relative bg-base-200 overflow-auto ${
                            !sidePanelMode ? `p-6` : `p-4`
                        } `}
                    >
                        <div
                            className={`grid ${
                                !sidePanelMode
                                    ? `grid grid-gap-6 xxxl:grid-gap-8 md:grid-auto-fit xxl:grid-auto-fit-wide md:grid-dense`
                                    : ``
                            } sm:grid-columns-1 grid-gap-5`}
                        >
                            <div
                                className="grid s-2 md:grid-auto-fit md:grid-dense"
                                style={{ '--min-tile-width': '50%' }}
                            >
                                <div className="s-full pb-3">
                                    <EntityCompliance
                                        entityType={entityTypes.NAMESPACE}
                                        entityId={params.entityId}
                                        entityName={namespace.name}
                                    />
                                </div>
                                <div className="md:pr-3 pt-3">
                                    <IconWidget
                                        title="Parent Cluster"
                                        icon={Cluster}
                                        description={namespace.clusterName}
                                        loading={loading}
                                    />
                                </div>
                                <div className="md:pl-3 pt-3">
                                    <CountWidget
                                        title="Network Policies"
                                        count={namespace.numNetworkPolicies}
                                    />
                                </div>
                            </div>

                            <Widget
                                className="sx-2"
                                header={`${namespace.labels.length} ${pluralize(
                                    'Label',
                                    namespace.labels.length
                                )}`}
                            >
                                <Labels list={namespace.labels.map(label => label.value)} />
                            </Widget>

                            <ComplianceByStandard
                                type={entityTypes.PCI_DSS_3_2}
                                entityName={namespace.name}
                                params={params}
                            />
                            <ComplianceByStandard
                                type={entityTypes.NIST_800_190}
                                entityName={namespace.name}
                                params={params}
                            />
                            <ComplianceByStandard
                                type={entityTypes.HIPAA_164}
                                entityName={namespace.name}
                                params={params}
                            />
                            <ComplianceByStandard
                                type={entityTypes.CIS_KUBERENETES_V1_2_0}
                                entityName={namespace.name}
                                params={params}
                            />
                            <ComplianceByStandard
                                type={entityTypes.CIS_DOCKER_V1_1_0}
                                entityName={namespace.name}
                                params={params}
                            />
                            <ResourceRelatedResourceList
                                listEntityType={entityTypes.DEPLOYMENT}
                                pageEntityType={entityTypes.NAMESPACE}
                                pageEntity={namespace}
                                className="sx-2"
                            />
                            <ResourceRelatedResourceList
                                listEntityType={entityTypes.SECRET}
                                pageEntityType={entityTypes.NAMESPACE}
                                pageEntity={namespace}
                                className="sx-2"
                            />
                        </div>
                    </div>
                </section>
            );
        }}
    </Query>
);

NamespacePage.propTypes = {
    sidePanelMode: PropTypes.bool,
    params: PropTypes.shape({}).isRequired
};

NamespacePage.defaultProps = {
    sidePanelMode: false
};

export default NamespacePage;
