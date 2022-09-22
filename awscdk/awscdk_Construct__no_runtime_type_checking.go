//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Construct) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_Construct) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateConstruct_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewConstructParameters(scope constructs.Construct, id *string) error {
	return nil
}

