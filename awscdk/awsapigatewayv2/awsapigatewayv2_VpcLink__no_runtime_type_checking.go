//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

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

func (v *jsiiProxy_VpcLink) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpcLink) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpcLink_FromVpcLinkAttributesParameters(scope constructs.Construct, id *string, attrs *VpcLinkAttributes) error {
	return nil
}

func validateVpcLink_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcLink_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewVpcLinkParameters(scope constructs.Construct, id *string, props *VpcLinkProps) error {
	return nil
}

