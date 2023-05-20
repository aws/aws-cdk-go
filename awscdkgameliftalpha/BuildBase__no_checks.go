//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BuildBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BuildBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BuildBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBuildBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBuildBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBuildBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBuildBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

