//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChannelGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (c *jsiiProxy_ChannelGrants) validateIngestParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateChannelGrants_FromChannelParameters(resource interfacesawsmediapackagev2.IChannelRef) error {
	return nil
}

