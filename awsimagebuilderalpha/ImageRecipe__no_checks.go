//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImageRecipe) validateAddBlockDeviceParameters(blockDevices *[]*awsec2.BlockDevice) error {
	return nil
}

func (i *jsiiProxy_ImageRecipe) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_ImageRecipe) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_ImageRecipe) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_ImageRecipe) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ImageRecipe) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateImageRecipe_FromImageRecipeArnParameters(scope constructs.Construct, id *string, imageRecipeArn *string) error {
	return nil
}

func validateImageRecipe_FromImageRecipeAttributesParameters(scope constructs.Construct, id *string, attrs *ImageRecipeAttributes) error {
	return nil
}

func validateImageRecipe_FromImageRecipeNameParameters(scope constructs.Construct, id *string, imageRecipeName *string) error {
	return nil
}

func validateImageRecipe_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImageRecipe_IsImageRecipeParameters(x interface{}) error {
	return nil
}

func validateImageRecipe_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateImageRecipe_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewImageRecipeParameters(scope constructs.Construct, id *string, props *ImageRecipeProps) error {
	return nil
}

