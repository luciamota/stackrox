/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useEffect, useState, ReactElement } from 'react';
import { Link } from 'react-router-dom';
import {
    Button,
    ButtonVariant,
    Flex,
    FlexItem,
    PageSection,
    PageSectionVariants,
    Text,
    TextContent,
    Title,
} from '@patternfly/react-core';

import ACSEmptyState from 'Components/ACSEmptyState';
import { vulnManagementReportsPath } from 'routePaths';
import { fetchReports } from 'services/ReportsService';
import { ReportConfigurationMappedValues } from 'types/report.proto';
import VulnMgmtReportTablePanel from './VulnMgmtReportTablePanel';
import VulnMgmtReportTableColumnDescriptor from './VulnMgmtReportTableColumnDescriptor';

function ReportTablePage(): ReactElement {
    const [reports, setReports] = useState<ReportConfigurationMappedValues[]>([]);
    const columns = VulnMgmtReportTableColumnDescriptor;

    useEffect(() => {
        fetchReports()
            .then((reportsResponse) => {
                setReports(reportsResponse);
            })
            .catch(() => {
                // TODO: show error message on failure
            });
    }, []);

    return (
        <>
            <PageSection variant={PageSectionVariants.light}>
                <Flex
                    alignItems={{
                        default: 'alignItemsFlexStart',
                        md: 'alignItemsCenter',
                    }}
                    direction={{ default: 'column', md: 'row' }}
                    flexWrap={{ default: 'nowrap' }}
                    spaceItems={{ default: 'spaceItemsXl' }}
                >
                    <FlexItem grow={{ default: 'grow' }}>
                        <TextContent>
                            <Title headingLevel="h1">Vulnerability reporting</Title>
                            <Text component="p">
                                Configure reports, define resource scopes, and assign distribution
                                lists to report on vulnerabilities across the organization.
                            </Text>
                        </TextContent>
                    </FlexItem>
                    <FlexItem
                        align={{
                            default: 'alignLeft',
                            md: 'alignRight',
                            lg: 'alignRight',
                            xl: 'alignRight',
                            '2xl': 'alignRight',
                        }}
                    >
                        <Button
                            variant={ButtonVariant.primary}
                            isInline
                            component={(props) => (
                                <Link
                                    {...props}
                                    to={`${vulnManagementReportsPath}?action=create`}
                                />
                            )}
                        >
                            Create report
                        </Button>
                    </FlexItem>
                </Flex>
            </PageSection>
            <PageSection variant={PageSectionVariants.light}>
                {reports.length > 0 ? (
                    <VulnMgmtReportTablePanel
                        reports={reports}
                        reportCount={0}
                        currentPage={0}
                        setCurrentPage={function (page: number): void {
                            throw new Error('Function not implemented.');
                        }}
                        perPage={0}
                        setPerPage={function (perPage: number): void {
                            throw new Error('Function not implemented.');
                        }}
                        activeSortIndex={0}
                        setActiveSortIndex={function (idx: number): void {
                            throw new Error('Function not implemented.');
                        }}
                        activeSortDirection="desc"
                        setActiveSortDirection={function (dir: string): void {
                            throw new Error('Function not implemented.');
                        }}
                        columns={columns}
                    />
                ) : (
                    <ACSEmptyState title="No reports are currently configured." />
                )}
            </PageSection>
        </>
    );
}

export default ReportTablePage;
