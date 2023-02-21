//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseListener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BaseListener) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BaseListener) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBaseListener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBaseListener_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBaseListener_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBaseListenerParameters(scope constructs.Construct, id *string, additionalProps interface{}) error {
	return nil
}

