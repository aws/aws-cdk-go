//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateRegisterCnameInstanceParameters(id *string, props *CnameInstanceBaseProps) error {
	return nil
}

func (s *jsiiProxy_Service) validateRegisterIpInstanceParameters(id *string, props *IpInstanceBaseProps) error {
	return nil
}

func (s *jsiiProxy_Service) validateRegisterLoadBalancerParameters(id *string, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) error {
	return nil
}

func (s *jsiiProxy_Service) validateRegisterNonIpInstanceParameters(id *string, props *NonIpInstanceBaseProps) error {
	return nil
}

func validateService_FromServiceAttributesParameters(scope constructs.Construct, id *string, attrs *ServiceAttributes) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, props *ServiceProps) error {
	return nil
}

