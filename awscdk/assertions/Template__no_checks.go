//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Template) validateAllResourcesParameters(type_ *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateAllResourcesPropertiesParameters(type_ *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateFindConditionsParameters(logicalId *string) error {
	return nil
}

func (t *jsiiProxy_Template) validateFindMappingsParameters(logicalId *string) error {
	return nil
}

func (t *jsiiProxy_Template) validateFindOutputsParameters(logicalId *string) error {
	return nil
}

func (t *jsiiProxy_Template) validateFindParametersParameters(logicalId *string) error {
	return nil
}

func (t *jsiiProxy_Template) validateFindResourcesParameters(type_ *string) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasConditionParameters(logicalId *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasMappingParameters(logicalId *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasOutputParameters(logicalId *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasParameterParameters(logicalId *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasResourceParameters(type_ *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateHasResourcePropertiesParameters(type_ *string, props interface{}) error {
	return nil
}

func (t *jsiiProxy_Template) validateResourceCountIsParameters(type_ *string, count *float64) error {
	return nil
}

func (t *jsiiProxy_Template) validateResourcePropertiesCountIsParameters(type_ *string, props interface{}, count *float64) error {
	return nil
}

func (t *jsiiProxy_Template) validateTemplateMatchesParameters(expected interface{}) error {
	return nil
}

func validateTemplate_FromJSONParameters(template *map[string]interface{}, templateParsingOptions *TemplateParsingOptions) error {
	return nil
}

func validateTemplate_FromStackParameters(stack awscdk.Stack, templateParsingOptions *TemplateParsingOptions) error {
	return nil
}

func validateTemplate_FromStringParameters(template *string, templateParsingOptions *TemplateParsingOptions) error {
	return nil
}

