//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Ipam) validateAddScopeParameters(scope constructs.Construct, id *string, options *IpamScopeOptions) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_Ipam) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateIpam_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpam_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateIpam_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIpamParameters(scope constructs.Construct, id *string, props *IpamProps) error {
	return nil
}

