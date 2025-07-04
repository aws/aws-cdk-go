//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChannelNamespace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ChannelNamespace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ChannelNamespace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ChannelNamespace) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ChannelNamespace) validateGrantPublishAndSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ChannelNamespace) validateGrantSubscribeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateChannelNamespace_FromChannelNamespaceArnParameters(scope constructs.Construct, id *string, channelNamespaceArn *string) error {
	return nil
}

func validateChannelNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateChannelNamespace_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateChannelNamespace_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewChannelNamespaceParameters(scope constructs.Construct, id *string, props *ChannelNamespaceProps) error {
	return nil
}

