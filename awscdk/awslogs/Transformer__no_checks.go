//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Transformer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Transformer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Transformer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransformer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransformer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransformer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransformerParameters(scope constructs.Construct, id *string, props *TransformerProps) error {
	return nil
}

