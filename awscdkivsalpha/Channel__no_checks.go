//go:build no_runtime_type_checking

package awscdkivsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Channel) validateAddStreamKeyParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Channel) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Channel) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Channel) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateChannel_FromChannelArnParameters(scope constructs.Construct, id *string, channelArn *string) error {
	return nil
}

func validateChannel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateChannel_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateChannel_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewChannelParameters(scope constructs.Construct, id *string, props *ChannelProps) error {
	return nil
}

