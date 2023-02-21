//go:build no_runtime_type_checking

package lambdalayernodeproxyagent

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodeProxyAgentLayer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (n *jsiiProxy_NodeProxyAgentLayer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NodeProxyAgentLayer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NodeProxyAgentLayer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNodeProxyAgentLayer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateNodeProxyAgentLayer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateNodeProxyAgentLayer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodeProxyAgentLayer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNodeProxyAgentLayer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNodeProxyAgentLayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

