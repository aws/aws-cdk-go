//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LayerVersion) validateAddPermissionParameters(id *string, permission *LayerVersionPermission) error {
	return nil
}

func (l *jsiiProxy_LayerVersion) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LayerVersion) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LayerVersion) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLayerVersion_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateLayerVersion_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *LayerVersionAttributes) error {
	return nil
}

func validateLayerVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLayerVersion_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLayerVersion_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLayerVersionParameters(scope constructs.Construct, id *string, props *LayerVersionProps) error {
	return nil
}

