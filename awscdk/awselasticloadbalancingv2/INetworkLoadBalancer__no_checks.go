//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INetworkLoadBalancer) validateAddListenerParameters(id *string, props *BaseNetworkListenerProps) error {
	return nil
}

func (i *jsiiProxy_INetworkLoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

