//go:build no_runtime_type_checking

package awsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcConnector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcConnector) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcConnector) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpcConnector) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpcConnector) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpcConnector_FromVpcConnectorAttributesParameters(scope constructs.Construct, id *string, attrs *VpcConnectorAttributes) error {
	return nil
}

func validateVpcConnector_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcConnector_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewVpcConnectorParameters(scope constructs.Construct, id *string, props *VpcConnectorProps) error {
	return nil
}

