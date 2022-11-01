//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplicationLoadBalancer) validateAddListenerParameters(id *string, props *BaseApplicationListenerProps) error {
	return nil
}

func (i *jsiiProxy_IApplicationLoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

