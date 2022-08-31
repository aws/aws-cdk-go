//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awselasticloadbalancing

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancer) validateAddListenerParameters(listener *LoadBalancerListener) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateAddTargetParameters(target ILoadBalancerTarget) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LoadBalancer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateLoadBalancer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLoadBalancer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewLoadBalancerParameters(scope constructs.Construct, id *string, props *LoadBalancerProps) error {
	return nil
}

