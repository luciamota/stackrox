import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form, Checkbox, TextArea } from '@patternfly/react-core';
import * as yup from 'yup';
import merge from 'lodash/merge';

import { ImageIntegrationBase } from 'services/ImageIntegrationsService';

import usePageState from 'Containers/Integrations/hooks/usePageState';
import FormMessage from 'Components/PatternFly/FormMessage';
import FormTestButton from 'Components/PatternFly/FormTestButton';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormLabelGroup from '../FormLabelGroup';

import { getGoogleCredentialsPlaceholder } from '../../utils/integrationUtils';

export type ArtifactRegistryIntegration = {
    categories: 'REGISTRY'[];
    google: {
        endpoint: string;
        project: string;
        wifEnabled: boolean;
        serviceAccount: string;
    };
    type: 'artifactregistry';
} & ImageIntegrationBase;

export type ArtifactRegistryIntegrationFormValues = {
    config: ArtifactRegistryIntegration;
    updatePassword: boolean;
};

export const validationSchema = yup.object().shape({
    config: yup.object().shape({
        name: yup.string().trim().required('An integration name is required'),
        categories: yup
            .array()
            .of(yup.string().trim().oneOf(['REGISTRY']))
            .min(1, 'Must have at least one type selected')
            .required('A category is required'),
        google: yup.object().shape({
            endpoint: yup.string().trim().required('An endpoint is required'),
            project: yup.string().trim().required('A project is required'),
            wifEnabled: yup.boolean(),
            serviceAccount: yup
                .string()
                .trim()
                .test(
                    'serviceAccount-test',
                    'Valid JSON is required for service account key',
                    (value, context: yup.TestContext) => {
                        const requirePasswordField =
                            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                            // @ts-ignore
                            context?.from[2]?.value?.updatePassword || false;
                        const useWorkloadId = context?.parent?.wifEnabled;

                        if (!requirePasswordField || useWorkloadId) {
                            return true;
                        }
                        try {
                            JSON.parse(value as string);
                        } catch (e) {
                            return false;
                        }
                        const trimmedValue = value?.trim();
                        return !!trimmedValue;
                    }
                ),
        }),
        skipTestIntegration: yup.bool(),
        type: yup.string().matches(/artifactregistry/),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: ArtifactRegistryIntegrationFormValues = {
    config: {
        id: '',
        name: '',
        categories: ['REGISTRY'],
        google: {
            endpoint: '',
            project: '',
            wifEnabled: false,
            serviceAccount: '',
        },
        autogenerated: false,
        clusterId: '',
        clusters: [],
        skipTestIntegration: false,
        type: 'artifactregistry',
    },
    updatePassword: true,
};

function ArtifactRegistryIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<ArtifactRegistryIntegration>): ReactElement {
    const formInitialValues = structuredClone(defaultValues);
    if (initialValues) {
        merge(formInitialValues.config, initialValues);

        // We want to clear the password because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.config.google.serviceAccount = '';

        // Don't assume user wants to change password; that has caused confusing UX.
        formInitialValues.updatePassword = false;
    }
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<ArtifactRegistryIntegrationFormValues>({
        initialValues: formInitialValues,
        validationSchema,
    });

    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onUpdateCredentialsChange(value, event) {
        setFieldValue('config.google.serviceAccount', '');
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="config.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.name"
                            placeholder="(ex. Google Artifact Registry)"
                            value={values.config.name}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Registry endpoint"
                        isRequired
                        fieldId="config.google.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.google.endpoint"
                            placeholder="(ex. us-west1-docker.pkg.dev)"
                            value={values.config.google.endpoint}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Project"
                        isRequired
                        fieldId="config.google.project"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.google.project"
                            value={values.config.google.project}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        fieldId="config.google.wifEnabled"
                        touched={touched}
                        errors={errors}
                    >
                        <Checkbox
                            label="Use workload identity"
                            id="config.google.wifEnabled"
                            isChecked={values.config.google.wifEnabled}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    {!isCreating && isEditable && (
                        <FormLabelGroup
                            label=""
                            fieldId="updatePassword"
                            helperText="Enable this option to replace currently stored credentials (if any)"
                            touched={touched}
                            errors={errors}
                        >
                            <Checkbox
                                label="Update stored credentials"
                                id="updatePassword"
                                isChecked={
                                    !values.config.google.wifEnabled && values.updatePassword
                                }
                                onChange={(event, value) => onUpdateCredentialsChange(value, event)}
                                onBlur={handleBlur}
                                isDisabled={!isEditable || values.config.google.wifEnabled}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        label="Service account key (JSON)"
                        isRequired={values.updatePassword && !values.config.google.wifEnabled}
                        fieldId="config.google.serviceAccount"
                        touched={touched}
                        errors={errors}
                    >
                        <TextArea
                            className="json-input"
                            isRequired={values.updatePassword && !values.config.google.wifEnabled}
                            type="text"
                            id="config.google.serviceAccount"
                            name="config.google.serviceAccount"
                            value={values.config.google.serviceAccount}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={
                                !isEditable ||
                                !values.updatePassword ||
                                values.config.google.wifEnabled
                            }
                            placeholder={getGoogleCredentialsPlaceholder(
                                values.config.google.wifEnabled,
                                values.updatePassword
                            )}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        fieldId="config.skipTestIntegration"
                        touched={touched}
                        errors={errors}
                    >
                        <Checkbox
                            label="Create integration without testing"
                            id="config.skipTestIntegration"
                            isChecked={values.config.skipTestIntegration}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default ArtifactRegistryIntegrationForm;
