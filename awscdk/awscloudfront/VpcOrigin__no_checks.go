//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcOrigin) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcOrigin) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcOrigin) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVpcOrigin_FromVpcOriginArnParameters(scope constructs.Construct, id *string, vpcOriginArn *string) error {
	return nil
}

func validateVpcOrigin_FromVpcOriginAttributesParameters(scope constructs.Construct, id *string, attrs *VpcOriginAttributes) error {
	return nil
}

func validateVpcOrigin_FromVpcOriginIdParameters(scope constructs.Construct, id *string, vpcOriginId *string) error {
	return nil
}

func validateVpcOrigin_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcOrigin_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpcOrigin_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpcOriginParameters(scope constructs.Construct, id *string, props *VpcOriginProps) error {
	return nil
}

