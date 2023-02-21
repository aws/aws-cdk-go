//go:build no_runtime_type_checking

package lambdalayerawscli

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCliLayer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (a *jsiiProxy_AwsCliLayer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AwsCliLayer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AwsCliLayer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAwsCliLayer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateAwsCliLayer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateAwsCliLayer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAwsCliLayer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAwsCliLayer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAwsCliLayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

