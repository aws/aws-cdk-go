//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package lambdalayerkubectl

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlLayer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlLayer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlLayer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlLayer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (k *jsiiProxy_KubectlLayer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (k *jsiiProxy_KubectlLayer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateKubectlLayer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlLayer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlLayer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlLayer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewKubectlLayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

