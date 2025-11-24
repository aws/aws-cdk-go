//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRecipe) validateAddInstanceBlockDeviceParameters(instanceBlockDevices *[]*awsec2.BlockDevice) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipe) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipe) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipe) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipe) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ContainerRecipe) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateContainerRecipe_FromContainerRecipeArnParameters(scope constructs.Construct, id *string, containerRecipeArn *string) error {
	return nil
}

func validateContainerRecipe_FromContainerRecipeAttributesParameters(scope constructs.Construct, id *string, attrs *ContainerRecipeAttributes) error {
	return nil
}

func validateContainerRecipe_FromContainerRecipeNameParameters(scope constructs.Construct, id *string, containerRecipeName *string) error {
	return nil
}

func validateContainerRecipe_IsConstructParameters(x interface{}) error {
	return nil
}

func validateContainerRecipe_IsContainerRecipeParameters(x interface{}) error {
	return nil
}

func validateContainerRecipe_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateContainerRecipe_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewContainerRecipeParameters(scope constructs.Construct, id *string, props *ContainerRecipeProps) error {
	return nil
}

