//go:build no_runtime_type_checking

package awsglobalaccelerator

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Listener) validateAddEndpointGroupParameters(id *string, options *EndpointGroupOptions) error {
	return nil
}

func (l *jsiiProxy_Listener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_Listener) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_Listener) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateListener_FromListenerArnParameters(scope constructs.Construct, id *string, listenerArn *string) error {
	return nil
}

func validateListener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateListener_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateListener_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewListenerParameters(scope constructs.Construct, id *string, props *ListenerProps) error {
	return nil
}

