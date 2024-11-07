import withAuth from '../../../helpers/basicAuth';
import { hasFeatureFlag } from '../../../helpers/features';

import {
    cancelAllCveExceptions,
    fillAndSubmitExceptionForm,
    getDateString,
    getFutureDateByDays,
    selectMultipleCvesForException,
    selectSingleCveForException,
    verifyExceptionConfirmationDetails,
    verifySelectedCvesInModal,
    visitAnyImageSinglePageWithMockedCves,
} from './WorkloadCves.helpers';

describe('Workload CVE Image page deferral and false positive flows', () => {
    withAuth();

    before(function () {
        if (!hasFeatureFlag('ROX_VULN_MGMT_WORKLOAD_CVES')) {
            this.skip();
        }
    });

    beforeEach(() => {
        cancelAllCveExceptions();
    });

    after(() => {
        if (hasFeatureFlag('ROX_VULN_MGMT_WORKLOAD_CVES')) {
            cancelAllCveExceptions();
        }
    });

    it('should defer a single CVE', () => {
        visitAnyImageSinglePageWithMockedCves().then((image) => {
            selectSingleCveForException('DEFERRAL').then((cveName) => {
                verifySelectedCvesInModal([cveName]);
                fillAndSubmitExceptionForm({
                    comment: 'Test comment',
                    expiryLabel: 'When all CVEs are fixable',
                });
                verifyExceptionConfirmationDetails({
                    expectedAction: 'Deferral',
                    cves: [cveName],
                    scope: `${image}:*`,
                    expiry: 'When all CVEs are fixable',
                });
            });
        });
    });

    it('should defer multiple selected CVEs', () => {
        visitAnyImageSinglePageWithMockedCves().then((image) => {
            selectMultipleCvesForException('DEFERRAL').then((cveNames) => {
                verifySelectedCvesInModal(cveNames);
                fillAndSubmitExceptionForm({
                    comment: 'Test comment',
                    expiryLabel: '30 days',
                });

                verifyExceptionConfirmationDetails({
                    expectedAction: 'Deferral',
                    cves: cveNames,
                    scope: `${image}:*`,
                    expiry: `${getDateString(getFutureDateByDays(30))} (30 days)`,
                });
            });
        });
    });

    it('should mark a single CVE as false positive', () => {
        visitAnyImageSinglePageWithMockedCves().then((image) => {
            selectSingleCveForException('FALSE_POSITIVE').then((cveName) => {
                verifySelectedCvesInModal([cveName]);
                fillAndSubmitExceptionForm({ comment: 'Test comment' });
                verifyExceptionConfirmationDetails({
                    expectedAction: 'False positive',
                    cves: [cveName],
                    scope: `${image}:*`,
                });
            });
        });
    });

    it('should mark multiple selected CVEs as false positive', () => {
        visitAnyImageSinglePageWithMockedCves().then((image) => {
            selectMultipleCvesForException('FALSE_POSITIVE').then((cveNames) => {
                verifySelectedCvesInModal(cveNames);
                fillAndSubmitExceptionForm({
                    comment: 'Test comment',
                });
                verifyExceptionConfirmationDetails({
                    expectedAction: 'False positive',
                    cves: cveNames,
                    scope: `${image}:*`,
                });
            });
        });
    });
});
