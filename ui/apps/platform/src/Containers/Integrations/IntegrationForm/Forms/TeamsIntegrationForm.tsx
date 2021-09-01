import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form } from '@patternfly/react-core';
import * as yup from 'yup';

import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessage from '../FormMessage';
import FormLabelGroup from '../FormLabelGroup';
import AnnotationKeyLabelIcon from '../AnnotationKeyLabelIcon';

export type TeamsIntegration = {
    id?: string;
    name: string;
    labelDefault: string;
    labelKey: string;
    uiEndpoint: string;
    type: 'teams';
    enabled: boolean;
};

const validTeamsWebhookRegex = /^((https?):\/\/)?(outlook.office365.com\/webhook\/)([a-zA-Z0-9-]+)$/;

export const validationSchema = yup.object().shape({
    name: yup.string().trim().required('Integration name is required'),
    labelDefault: yup
        .string()
        .trim()
        .required('Webhook is required')
        .matches(
            validTeamsWebhookRegex,
            'Must be a valid Teams webhook URL, like https://outlook.office365.com/webhook/EXAMPLE'
        ),
    labelKey: yup.string().trim(),
    uiEndpoint: yup.string(),
    type: yup.string().matches(/teams/),
    enabled: yup.bool(),
});

export const defaultValues: TeamsIntegration = {
    name: '',
    labelDefault: '',
    labelKey: '',
    uiEndpoint: window.location.origin,
    type: 'teams',
    enabled: true,
};

function TeamsIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<TeamsIntegration>): ReactElement {
    const formInitialValues = initialValues
        ? { ...defaultValues, ...initialValues }
        : defaultValues;
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
    } = useIntegrationForm<TeamsIntegration, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });

    function onChange(value, event) {
        return setFieldValue(event.target.id, value, false);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                {message && <FormMessage message={message} />}
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
                            placeholder="(example, Teams Integration)"
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Default Teams webhook"
                        isRequired
                        fieldId="labelDefault"
                        touched={touched}
                        errors={errors}
                        helperText="For example, https://outlook.office365.com/webhook/EXAMPLE"
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="labelDefault"
                            value={values.labelDefault}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Annotation key for Teams webhook"
                        labelIcon={<AnnotationKeyLabelIcon />}
                        fieldId="labelKey"
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="labelKey"
                            value={values.labelKey}
                            onChange={onChange}
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
                        isValid={isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default TeamsIntegrationForm;
