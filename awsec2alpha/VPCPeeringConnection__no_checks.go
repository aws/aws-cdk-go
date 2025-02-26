//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VPCPeeringConnection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VPCPeeringConnection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VPCPeeringConnection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVPCPeeringConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVPCPeeringConnection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVPCPeeringConnection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVPCPeeringConnectionParameters(scope constructs.Construct, id *string, props *VPCPeeringConnectionProps) error {
	return nil
}

