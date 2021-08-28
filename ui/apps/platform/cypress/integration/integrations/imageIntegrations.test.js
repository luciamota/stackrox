import { selectors } from '../../constants/IntegrationsPage';
import * as api from '../../constants/apiEndpoints';
import withAuth from '../../helpers/basicAuth';
import { editIntegration } from './integrationUtils';

describe('Image Integrations Test', () => {
    withAuth();

    beforeEach(() => {
        cy.intercept('GET', api.integrations.imageIntegrations, {
            fixture: 'integrations/imageIntegrations.json',
        }).as('getImageIntegrations');

        cy.visit('/');
        cy.get(selectors.configure).click();
        cy.get(selectors.navLink).click({ force: true });
        cy.wait('@getImageIntegrations');
    });

    it('should show a hint about stored credentials for Docker Trusted Registry', () => {
        cy.get(selectors.dockerTrustedRegistryTile).click();
        editIntegration('DTR Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Quay', () => {
        cy.get(selectors.quayTile).click();
        editIntegration('Quay Test');
        cy.get('div:contains("OAuth Token"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Amazon ECR', () => {
        cy.get(selectors.amazonECRTile).click();
        editIntegration('Amazon ECR Test');
        cy.get('div:contains("Access Key ID"):last [data-testid="help-icon"]').trigger(
            'mouseenter'
        );
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
        cy.get('div:contains("Secret Access Key"):last [data-testid="help-icon"]').trigger(
            'mouseenter'
        );
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Tenable', () => {
        cy.get(selectors.tenableTile).click();
        editIntegration('Tenable Test');
        cy.get('div:contains("Access Key"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
        cy.get('div:contains("Secret Key"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Google Container Registry', () => {
        cy.get(selectors.googleContainerRegistryTile).click();
        editIntegration('Google Container Registry Test');
        cy.get('div:contains("Service Account Key"):last [data-testid="help-icon"]').trigger(
            'mouseenter'
        );
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Anchore Scanner', () => {
        cy.get(selectors.anchoreScannerTile).click();
        editIntegration('Anchore Scanner Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for IBM Cloud', () => {
        cy.get(selectors.ibmCloudTile).click();
        editIntegration('IBM Cloud Test');
        cy.get('div:contains("API Key"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Microsoft ACR', () => {
        cy.get(selectors.microsoftACRTile).click();
        editIntegration('Microsoft ACR Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for JFrog Artifactory', () => {
        cy.get(selectors.jFrogArtifactoryTile).click();
        editIntegration('JFrog Artifactory Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Sonatype Nexus', () => {
        cy.get(selectors.sonatypeNexusTile).click();
        editIntegration('Sonatype Nexus Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });

    it('should show a hint about stored credentials for Red Hat', () => {
        cy.get(selectors.redHatTile).click();
        editIntegration('Red Hat Test');
        cy.get('div:contains("Password"):last [data-testid="help-icon"]').trigger('mouseenter');
        cy.get(selectors.tooltip.overlay).contains(
            'Leave this empty to use the currently stored credentials'
        );
    });
});
