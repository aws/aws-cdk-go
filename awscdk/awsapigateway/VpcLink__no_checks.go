//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcLink) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcLink) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcLink) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVpcLink_FromVpcLinkIdParameters(scope constructs.Construct, id *string, vpcLinkId *string) error {
	return nil
}

func validateVpcLink_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcLink_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpcLink_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpcLinkParameters(scope constructs.Construct, id *string, props *VpcLinkProps) error {
	return nil
}

