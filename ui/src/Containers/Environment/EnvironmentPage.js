import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createSelector, createStructuredSelector } from 'reselect';
import { selectors } from 'reducers';
import { actions as environmentActions } from 'reducers/environment';

import PageHeader from 'Components/PageHeader';
import SearchInput from 'Components/SearchInput';
import EnvironmentGraph from 'Components/EnvironmentGraph';
import * as Icon from 'react-feather';

class EnvironmentPage extends Component {
    static propTypes = {
        searchOptions: PropTypes.arrayOf(PropTypes.object).isRequired,
        searchModifiers: PropTypes.arrayOf(PropTypes.object).isRequired,
        searchSuggestions: PropTypes.arrayOf(PropTypes.object).isRequired,
        setSearchOptions: PropTypes.func.isRequired,
        setSearchModifiers: PropTypes.func.isRequired,
        setSearchSuggestions: PropTypes.func.isRequired,
        isViewFiltered: PropTypes.bool.isRequired,
        fetchEnvironmentGraph: PropTypes.func.isRequired,
        environmentGraph: PropTypes.shape({
            nodes: PropTypes.arrayOf(
                PropTypes.shape({
                    id: PropTypes.string.isRequired
                })
            ),
            edges: PropTypes.arrayOf(
                PropTypes.shape({
                    source: PropTypes.string.isRequired,
                    target: PropTypes.string.isRequired
                })
            )
        }).isRequired
    };

    renderGraph = () => (
        <EnvironmentGraph
            nodes={this.props.environmentGraph.nodes}
            edges={this.props.environmentGraph.edges}
        />
    );

    render() {
        const subHeader = this.props.isViewFiltered ? 'Filtered view' : 'Default view';
        return (
            <section className="flex flex-1 h-full w-full">
                <div className="flex flex-1 flex-col w-full">
                    <div>
                        <PageHeader header="Environment" subHeader={subHeader}>
                            <SearchInput
                                id="images"
                                searchOptions={this.props.searchOptions}
                                searchModifiers={this.props.searchModifiers}
                                searchSuggestions={this.props.searchSuggestions}
                                setSearchOptions={this.props.setSearchOptions}
                                setSearchModifiers={this.props.setSearchModifiers}
                                setSearchSuggestions={this.props.setSearchSuggestions}
                            />
                        </PageHeader>
                    </div>
                    <section className="environment-grid-bg flex flex-1 relative">
                        {this.renderGraph()}
                        <button
                            className="btn-graph-refresh absolute pin-t pin-r mt-2 mr-2 p-2 bg-primary-300 rounded-sm text-sm"
                            onClick={this.props.fetchEnvironmentGraph}
                        >
                            <Icon.Circle className="h-2 w-2" />
                            <span className="pl-1">Node updates available</span>
                        </button>
                    </section>
                </div>
            </section>
        );
    }
}

const isViewFiltered = createSelector(
    [selectors.getEnvironmentSearchOptions],
    searchOptions => searchOptions.length !== 0
);

const mapStateToProps = createStructuredSelector({
    environmentGraph: selectors.getEnvironmentGraph,
    searchOptions: selectors.getEnvironmentSearchOptions,
    searchModifiers: selectors.getEnvironmentSearchModifiers,
    searchSuggestions: selectors.getEnvironmentSearchSuggestions,
    isViewFiltered
});

const mapDispatchToProps = {
    fetchEnvironmentGraph: environmentActions.fetchEnvironmentGraph.request,
    setSearchOptions: environmentActions.setEnvironmentSearchOptions,
    setSearchModifiers: environmentActions.setEnvironmentSearchModifiers,
    setSearchSuggestions: environmentActions.setEnvironmentSearchSuggestions
};

export default connect(mapStateToProps, mapDispatchToProps)(EnvironmentPage);
