//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkListenerAction) validateBindParameters(scope constructs.Construct, listener INetworkListener) error {
	return nil
}

func (n *jsiiProxy_NetworkListenerAction) validateRenumberParameters(actions *[]*CfnListener_ActionProperty) error {
	return nil
}

func validateNetworkListenerAction_ForwardParameters(targetGroups *[]INetworkTargetGroup, options *NetworkForwardOptions) error {
	return nil
}

func validateNetworkListenerAction_WeightedForwardParameters(targetGroups *[]*NetworkWeightedTargetGroup, options *NetworkForwardOptions) error {
	return nil
}

func validateNewNetworkListenerActionParameters(defaultActionJson *CfnListener_ActionProperty) error {
	return nil
}

