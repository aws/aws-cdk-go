//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSyncEventResource) validateResourceArnsParameters(api EventApiBase) error {
	return nil
}

func validateAppSyncEventResource_OfChannelNamespaceParameters(channelNamespace *string) error {
	return nil
}

