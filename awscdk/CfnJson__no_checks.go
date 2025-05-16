//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnJson) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateCfnJson_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnJsonParameters(scope constructs.Construct, id *string, props *CfnJsonProps) error {
	return nil
}

