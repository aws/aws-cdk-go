//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VirtualService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VirtualService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVirtualService_FromVirtualServiceArnParameters(scope constructs.Construct, id *string, virtualServiceArn *string) error {
	return nil
}

func validateVirtualService_FromVirtualServiceAttributesParameters(scope constructs.Construct, id *string, attrs *VirtualServiceAttributes) error {
	return nil
}

func validateVirtualService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVirtualService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVirtualService_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVirtualServiceParameters(scope constructs.Construct, id *string, props *VirtualServiceProps) error {
	return nil
}

