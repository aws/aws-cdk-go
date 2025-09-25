//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcIngressConnection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcIngressConnection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcIngressConnection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVpcIngressConnection_FromArnParameters(scope constructs.Construct, id *string, vpcIngressConnectionArn *string) error {
	return nil
}

func validateVpcIngressConnection_FromVpcIngressConnectionAttributesParameters(scope constructs.Construct, id *string, attrs *VpcIngressConnectionAttributes) error {
	return nil
}

func validateVpcIngressConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcIngressConnection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpcIngressConnection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpcIngressConnectionParameters(scope constructs.Construct, id *string, props *VpcIngressConnectionProps) error {
	return nil
}

