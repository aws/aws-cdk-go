//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConstructNode) validateAddErrorParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddInfoParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddMetadataParameters(type_ *string, data interface{}) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddValidationParameters(validation constructs.IValidation) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddWarningParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateApplyAspectParameters(aspect IAspect) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateFindChildParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateSetContextParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryFindChildParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryGetContextParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryRemoveChildParameters(childName *string) error {
	return nil
}

func validateConstructNode_PrepareParameters(node ConstructNode) error {
	return nil
}

func validateConstructNode_SynthParameters(node ConstructNode, options *SynthesisOptions) error {
	return nil
}

func validateConstructNode_ValidateParameters(node ConstructNode) error {
	return nil
}

func validateNewConstructNodeParameters(host Construct, scope IConstruct, id *string) error {
	return nil
}

