//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChannelGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (c *jsiiProxy_ChannelGrants) validateStartParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ChannelGrants) validateStopParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ChannelGrants) validateUpdateScheduleParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateChannelGrants_FromChannelParameters(resource interfacesawsmedialive.IChannelRef) error {
	return nil
}

