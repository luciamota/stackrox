import { VulnerabilitySeverity } from 'types/cve.proto';
import { Fixability } from 'types/report.proto';

type VulnerabilitySeverityLabels = Record<VulnerabilitySeverity, string>;

export const vulnerabilitySeverityLabels: VulnerabilitySeverityLabels = {
    CRITICAL_VULNERABILITY_SEVERITY: 'Critical',
    IMPORTANT_VULNERABILITY_SEVERITY: 'Important',
    MODERATE_VULNERABILITY_SEVERITY: 'Medium',
    LOW_VULNERABILITY_SEVERITY: 'Low',
};

type FixabilityLabelKeys = Exclude<Fixability, 'BOTH'>;
type FixabilityLabels = Record<FixabilityLabelKeys, string>;

export const fixabilityLabels: FixabilityLabels = {
    FIXABLE: 'Fixable',
    NOT_FIXABLE: 'Unfixable',
};
