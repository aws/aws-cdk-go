//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Validations) validateAcknowledgeParameters(rules *[]*Acknowledgment) error {
	return nil
}

func (v *jsiiProxy_Validations) validateAddErrorParameters(id *string, message *string) error {
	return nil
}

func (v *jsiiProxy_Validations) validateAddWarningParameters(id *string, message *string) error {
	return nil
}

func validateValidations_OfParameters(scope constructs.IConstruct) error {
	return nil
}

