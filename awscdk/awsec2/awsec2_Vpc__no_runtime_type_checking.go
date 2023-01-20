//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Vpc) validateAddClientVpnEndpointParameters(id *string, options *ClientVpnEndpointOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateAddFlowLogParameters(id *string, options *FlowLogOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateAddGatewayEndpointParameters(id *string, options *GatewayVpcEndpointOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateAddInterfaceEndpointParameters(id *string, options *InterfaceVpcEndpointOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateAddVpnConnectionParameters(id *string, options *VpnConnectionOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateEnableVpnGatewayParameters(options *EnableVpnGatewayOptions) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateSelectSubnetObjectsParameters(selection *SubnetSelection) error {
	return nil
}

func (v *jsiiProxy_Vpc) validateSelectSubnetsParameters(selection *SubnetSelection) error {
	return nil
}

func validateVpc_FromLookupParameters(scope constructs.Construct, id *string, options *VpcLookupOptions) error {
	return nil
}

func validateVpc_FromVpcAttributesParameters(scope constructs.Construct, id *string, attrs *VpcAttributes) error {
	return nil
}

func validateVpc_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpc_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpc_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Vpc) validateSetIncompleteSubnetDefinitionParameters(val *bool) error {
	return nil
}

func validateNewVpcParameters(scope constructs.Construct, id *string, props *VpcProps) error {
	return nil
}

