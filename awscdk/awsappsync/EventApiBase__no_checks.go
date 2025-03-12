//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventApiBase) validateAddChannelNamespaceParameters(id *string, options *ChannelNamespaceOptions) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGrantParameters(grantee awsiam.IGrantable, resources AppSyncEventResource) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGrantPublishAndSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EventApiBase) validateGrantSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEventApiBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventApiBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEventApiBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEventApiBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

