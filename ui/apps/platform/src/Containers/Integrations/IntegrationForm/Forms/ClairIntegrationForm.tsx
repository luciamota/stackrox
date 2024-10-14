import React, { ReactElement } from 'react';
import {
    Alert,
    AlertVariant,
    Checkbox,
    Form,
    PageSection,
    Text,
    TextInput,
} from '@patternfly/react-core';
import * as yup from 'yup';
import merge from 'lodash/merge';

import { ImageIntegrationBase } from 'services/ImageIntegrationsService';

import FormMessage from 'Components/PatternFly/FormMessage';
import FormTestButton from 'Components/PatternFly/FormTestButton';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormLabelGroup from '../FormLabelGroup';

export type ClairIntegration = {
    categories: 'SCANNER'[];
    clair: {
        endpoint: string;
        insecure: boolean;
    };
    type: 'clair';
} & ImageIntegrationBase;

export const validationSchema = yup.object().shape({
    name: yup.string().trim().required('An integration name is required'),
    categories: yup
        .array()
        .of(yup.string().trim().oneOf(['SCANNER']))
        .min(1, 'Must have at least one type selected')
        .required('A category is required'),
    clair: yup.object().shape({
        endpoint: yup.string().trim().required('An endpoint is required').min(1),
        insecure: yup.bool(),
    }),
    type: yup.string().matches(/clair/),
});

export const defaultValues: ClairIntegration = {
    id: '',
    name: '',
    categories: ['SCANNER'],
    clair: {
        endpoint: '',
        insecure: false,
    },
    autogenerated: false,
    clusterId: '',
    clusters: [],
    skipTestIntegration: false,
    type: 'clair',
};

function ClairIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<ClairIntegration>): ReactElement {
    const formInitialValues: ClairIntegration = merge({}, defaultValues, initialValues);
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
    } = useIntegrationForm<ClairIntegration>({
        initialValues: formInitialValues,
        validationSchema,
    });

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <Alert
                    title="Deprecation notice"
                    component="p"
                    variant={AlertVariant.warning}
                    isInline
                    className="pf-v5-u-mb-lg"
                >
                    <Text>
                        CoreOS Clair v2 integration will be removed in Red Hat Advanced Cluster
                        Security 4.5 release.
                    </Text>
                    <Text>Use Clair v4 integration instead.</Text>
                </Alert>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="name"
                            value={values.name}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        isRequired
                        fieldId="clair.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="clair.endpoint"
                            value={values.clair.endpoint}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup fieldId="clair.insecure" touched={touched} errors={errors}>
                        <Checkbox
                            label="Disable TLS certificate validation (insecure)"
                            id="clair.insecure"
                            isChecked={values.clair.insecure}
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

export default ClairIntegrationForm;
