//go:build no_runtime_type_checking

package awslambdapython

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PythonLayerVersion) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (p *jsiiProxy_PythonLayerVersion) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PythonLayerVersion) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PythonLayerVersion) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PythonLayerVersion) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_PythonLayerVersion) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validatePythonLayerVersion_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validatePythonLayerVersion_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validatePythonLayerVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePythonLayerVersion_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewPythonLayerVersionParameters(scope awscdk.Construct, id *string, props *PythonLayerVersionProps) error {
	return nil
}

