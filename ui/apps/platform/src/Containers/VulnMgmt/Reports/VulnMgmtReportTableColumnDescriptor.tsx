import React, { ReactElement } from 'react';
import { Link } from 'react-router-dom';
import { Button, ButtonVariant, Label } from '@patternfly/react-core';

import DateTimeFormat from 'Components/PatternFly/DateTimeFormat';
import { fixabilityLabels, vulnerabilitySeverityLabels } from 'constants/reportConstants';
import { vulnManagementReportsPath } from 'routePaths';

const VulnMgmtReportTableColumnDescriptor = [
    {
        Header: 'Report',
        accessor: 'report.name',
        sortField: 'Report',
        Cell: ({ original }) => {
            const url = `${vulnManagementReportsPath}/${original.id as string}`;
            return (
                <Button
                    variant={ButtonVariant.link}
                    isInline
                    component={(props) => <Link {...props} to={url} />}
                >
                    {original?.name}
                </Button>
            );
        },
    },
    {
        Header: 'Description',
        accessor: 'description',
        Cell: ({ value }): ReactElement => {
            return <span>{value}</span>;
        },
    },
    {
        Header: 'CVE fixability type',
        accessor: 'vulnReportFiltersMappedValues.fixabilityMappedValues',
        Cell: ({ value }): ReactElement => {
            const fixabilityStrings = value.map((fixValue) => fixabilityLabels[fixValue] as string);
            return <span>{fixabilityStrings.join(', ') || 'Issue: Fixabiltiy unset'}</span>;
        },
    },
    {
        Header: 'CVE severities',
        accessor: 'vulnReportFiltersMappedValues.severities',
        sortField: 'CVE severities',
        Cell: ({ value }): ReactElement => {
            const severityLabels = value.map((fixValue) => (
                <Label className="pf-u-mr-sm" color="red" isCompact>
                    {vulnerabilitySeverityLabels[fixValue]}
                </Label>
            ));
            return <>{severityLabels}</>;
        },
    },
    {
        Header: 'Last run',
        accessor: 'runStatus',
        sortField: 'Last run',
        Cell: ({ value }): ReactElement => {
            const lastRunAttempt = value?.lastTimeRun;
            return lastRunAttempt ? (
                <DateTimeFormat time={lastRunAttempt} />
            ) : (
                <span>Not run yet</span>
            );
        },
    },
];

export default VulnMgmtReportTableColumnDescriptor;
