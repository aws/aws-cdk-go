//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConsumableResourceGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (c *jsiiProxy_ConsumableResourceGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ConsumableResourceGrants) validateUpdateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateConsumableResourceGrants_FromConsumableResourceParameters(resource interfacesawsbatch.IConsumableResourceRef) error {
	return nil
}

