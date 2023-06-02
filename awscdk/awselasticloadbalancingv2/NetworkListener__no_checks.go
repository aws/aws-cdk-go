//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkListener) validateAddActionParameters(_id *string, props *AddNetworkActionProps) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateAddCertificatesParameters(id *string, certificates *[]IListenerCertificate) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateAddTargetGroupsParameters(_id *string) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateAddTargetsParameters(id *string, props *AddNetworkTargetsProps) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkListener) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNetworkListener_FromLookupParameters(scope constructs.Construct, id *string, options *NetworkListenerLookupOptions) error {
	return nil
}

func validateNetworkListener_FromNetworkListenerArnParameters(scope constructs.Construct, id *string, networkListenerArn *string) error {
	return nil
}

func validateNetworkListener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkListener_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetworkListener_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkListenerParameters(scope constructs.Construct, id *string, props *NetworkListenerProps) error {
	return nil
}

