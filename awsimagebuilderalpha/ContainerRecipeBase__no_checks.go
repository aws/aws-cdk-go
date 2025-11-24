//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRecipeBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipeBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipeBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipeBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipeBase) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateContainerRecipeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateContainerRecipeBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateContainerRecipeBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewContainerRecipeBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

