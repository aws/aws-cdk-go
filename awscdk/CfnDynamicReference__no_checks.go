//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDynamicReference) validateNewErrorParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_CfnDynamicReference) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateNewCfnDynamicReferenceParameters(service CfnDynamicReferenceService, key *string) error {
	return nil
}

