//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseLoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateLogAccessLogsParameters(bucket awss3.IBucket) error {
	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateRemoveAttributeParameters(key *string) error {
	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateSetAttributeParameters(key *string) error {
	return nil
}

func validateBaseLoadBalancer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBaseLoadBalancer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBaseLoadBalancer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBaseLoadBalancerParameters(scope constructs.Construct, id *string, baseProps *BaseLoadBalancerProps, additionalProps interface{}) error {
	return nil
}

