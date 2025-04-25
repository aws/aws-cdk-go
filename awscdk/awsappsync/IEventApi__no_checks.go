//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEventApi) validateAddChannelNamespaceParameters(id *string, options *ChannelNamespaceOptions) error {
	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantParameters(grantee awsiam.IGrantable, resources AppSyncEventResource) error {
	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantPublishAndSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEventApi) validateGrantSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

