//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagOptions) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TagOptions) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TagOptions) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTagOptions_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTagOptions_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTagOptions_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTagOptionsParameters(scope constructs.Construct, id *string, props *TagOptionsProps) error {
	return nil
}

