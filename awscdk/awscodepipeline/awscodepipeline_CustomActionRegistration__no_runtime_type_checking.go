//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomActionRegistration) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CustomActionRegistration) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCustomActionRegistration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomActionRegistrationParameters(scope constructs.Construct, id *string, props *CustomActionRegistrationProps) error {
	return nil
}

