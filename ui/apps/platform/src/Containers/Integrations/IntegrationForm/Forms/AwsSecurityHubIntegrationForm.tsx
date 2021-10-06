import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form, FormSelect, Checkbox } from '@patternfly/react-core';
import * as yup from 'yup';

import { NotifierIntegrationBase } from 'services/NotifierIntegrationsService';

import usePageState from 'Containers/Integrations/hooks/usePageState';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessage from '../FormMessage';
import FormLabelGroup from '../FormLabelGroup';
import AwsRegionOptions from '../AwsRegionOptions';

export type AwsSecurityHubIntegration = {
    awsSecurityHub: {
        accountId: string;
        region: string;
        credentials: {
            accessKeyId: string;
            secretAccessKey: string;
        };
    };
    type: 'awsSecurityHub';
} & NotifierIntegrationBase;

export type AwsSecurityHubIntegrationFormValues = {
    notifier: AwsSecurityHubIntegration;
    updatePassword: boolean;
};

export const validationSchema = yup.object().shape({
    notifier: yup.object().shape({
        name: yup.string().required('An integration name is required'),
        awsSecurityHub: yup.object().shape({
            accountId: yup
                .string()
                .required('An AWS account number is required')
                .length(12, 'AWS account numbers must be 12 characters long'),
            region: yup.string().required('An AWS region is required'),
            credentials: yup.object().shape({
                accessKeyId: yup
                    .string()
                    .test(
                        'accessKeyId-test',
                        'An access key ID is required',
                        (value, context: yup.TestContext) => {
                            const requirePasswordField =
                                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                                // @ts-ignore
                                context?.from[3]?.value?.updatePassword || false;

                            if (!requirePasswordField) {
                                return true;
                            }

                            const trimmedValue = value?.trim();
                            return !!trimmedValue;
                        }
                    ),
                secretAccessKey: yup
                    .string()
                    .test(
                        'secretAccessKey-test',
                        'A secret access key is required',
                        (value, context: yup.TestContext) => {
                            const requirePasswordField =
                                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                                // @ts-ignore
                                context?.from[3]?.value?.updatePassword || false;

                            if (!requirePasswordField) {
                                return true;
                            }

                            const trimmedValue = value?.trim();
                            return !!trimmedValue;
                        }
                    ),
            }),
        }),
        uiEndpoint: yup.string(),
        type: yup.string().matches(/awsSecurityHub/),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: AwsSecurityHubIntegrationFormValues = {
    notifier: {
        id: '',
        name: '',
        awsSecurityHub: {
            accountId: '',
            region: '',
            credentials: {
                accessKeyId: '',
                secretAccessKey: '',
            },
        },
        labelDefault: '',
        labelKey: '',
        uiEndpoint: window.location.origin,
        type: 'awsSecurityHub',
    },
    updatePassword: true,
};

function AwsSecurityHubIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<AwsSecurityHubIntegration>): ReactElement {
    const formInitialValues = { ...defaultValues, ...initialValues };
    if (initialValues) {
        formInitialValues.notifier = {
            ...formInitialValues.notifier,
            ...initialValues,
        };
        // We want to clear these values because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.notifier.awsSecurityHub.credentials.accessKeyId = '';
        formInitialValues.notifier.awsSecurityHub.credentials.secretAccessKey = '';
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
    } = useIntegrationForm<AwsSecurityHubIntegrationFormValues>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onUpdateCredentialsChange(value, event) {
        setFieldValue('notifier.awsSecurityHub.credentials.accessKeyId', '');
        setFieldValue('notifier.awsSecurityHub.credentials.secretAccessKey', '');
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        isRequired
                        label="Integration name"
                        fieldId="notifier.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="notifier.name"
                            value={values.notifier.name}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        isRequired
                        label="AWS account number"
                        fieldId="notifier.awsSecurityHub.accountId"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="notifier.awsSecurityHub.accountId"
                            value={values.notifier.awsSecurityHub.accountId}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        isRequired
                        label="AWS region"
                        fieldId="notifier.awsSecurityHub.region"
                        touched={touched}
                        errors={errors}
                    >
                        <FormSelect
                            id="notifier.awsSecurityHub.region"
                            value={values.notifier.awsSecurityHub.region}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        >
                            <AwsRegionOptions />
                        </FormSelect>
                    </FormLabelGroup>
                    {!isCreating && isEditable && (
                        <FormLabelGroup
                            label=""
                            fieldId="updatePassword"
                            helperText="Leave this off to use the currently stored credentials."
                            errors={errors}
                        >
                            <Checkbox
                                label="Update stored credentials"
                                id="updatePassword"
                                isChecked={values.updatePassword}
                                onChange={onUpdateCredentialsChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        isRequired={values.updatePassword}
                        label="Access key ID"
                        fieldId="notifier.awsSecurityHub.credentials.accessKeyId"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired={values.updatePassword}
                            type="password"
                            id="notifier.awsSecurityHub.credentials.accessKeyId"
                            value={values.notifier.awsSecurityHub.credentials.accessKeyId}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || !values.updatePassword}
                            placeholder={
                                values.updatePassword
                                    ? ''
                                    : 'Currently-stored access key ID will be used.'
                            }
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        isRequired={values.updatePassword}
                        label="Secret access key"
                        fieldId="notifier.awsSecurityHub.credentials.secretAccessKey"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired={values.updatePassword}
                            type="password"
                            id="notifier.awsSecurityHub.credentials.secretAccessKey"
                            value={values.notifier.awsSecurityHub.credentials.secretAccessKey}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || !values.updatePassword}
                            placeholder={
                                values.updatePassword
                                    ? ''
                                    : 'Currently-stored secret access key will be used.'
                            }
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

export default AwsSecurityHubIntegrationForm;
