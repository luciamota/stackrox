import React from 'react';
import { useQuery } from '@apollo/client';
import { Divider } from '@patternfly/react-core';

import useURLSort from 'hooks/useURLSort';
import useURLPagination from 'hooks/useURLPagination';

import { getTableUIState } from 'utils/getTableUIState';
import { getPaginationParams } from 'utils/searchUtils';
import { SearchFilter } from 'types/search';
import DeploymentsTable, { Deployment, deploymentListQuery } from '../Tables/DeploymentsTable';
import TableEntityToolbar, { TableEntityToolbarProps } from '../../components/TableEntityToolbar';
import { VulnerabilitySeverityLabel } from '../../types';

type DeploymentsTableContainerProps = {
    searchFilter: SearchFilter;
    onFilterChange: (searchFilter: SearchFilter) => void;
    filterToolbar: TableEntityToolbarProps['filterToolbar'];
    entityToggleGroup: TableEntityToolbarProps['entityToggleGroup'];
    rowCount: number;
    pagination: ReturnType<typeof useURLPagination>;
    sort: ReturnType<typeof useURLSort>;
    workloadCvesScopedQueryString: string;
    isFiltered: boolean;
    showCveDetailFields: boolean;
};

function DeploymentsTableContainer({
    searchFilter,
    onFilterChange,
    filterToolbar,
    entityToggleGroup,
    rowCount,
    pagination,
    sort,
    workloadCvesScopedQueryString,
    isFiltered,
    showCveDetailFields,
}: DeploymentsTableContainerProps) {
    const { page, perPage } = pagination;
    const { sortOption, getSortParams } = sort;

    const { error, loading, data } = useQuery<{
        deployments: Deployment[];
    }>(deploymentListQuery, {
        variables: {
            query: workloadCvesScopedQueryString,
            pagination: getPaginationParams({ page, perPage, sortOption }),
        },
    });

    const tableState = getTableUIState({
        isLoading: loading,
        error,
        data: data?.deployments,
        searchFilter,
    });

    return (
        <>
            <TableEntityToolbar
                filterToolbar={filterToolbar}
                entityToggleGroup={entityToggleGroup}
                pagination={pagination}
                tableRowCount={rowCount}
                isFiltered={isFiltered}
            />
            <Divider component="div" />
            <div
                className="workload-cves-table-container"
                aria-live="polite"
                aria-busy={loading ? 'true' : 'false'}
            >
                <DeploymentsTable
                    tableState={tableState}
                    getSortParams={getSortParams}
                    isFiltered={isFiltered}
                    filteredSeverities={searchFilter.SEVERITY as VulnerabilitySeverityLabel[]}
                    showCveDetailFields={showCveDetailFields}
                    onClearFilters={() => {
                        onFilterChange({});
                        pagination.setPage(1);
                    }}
                />
            </div>
        </>
    );
}

export default DeploymentsTableContainer;
