//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Image) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_Image) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_Image) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_Image) validateGrantDefaultExecutionRolePermissionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_Image) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateImage_FromImageArnParameters(scope constructs.Construct, id *string, imageArn *string) error {
	return nil
}

func validateImage_FromImageAttributesParameters(scope constructs.Construct, id *string, attrs *ImageAttributes) error {
	return nil
}

func validateImage_FromImageNameParameters(scope constructs.Construct, id *string, imageName *string) error {
	return nil
}

func validateImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImage_IsImageParameters(x interface{}) error {
	return nil
}

func validateImage_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateImage_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewImageParameters(scope constructs.Construct, id *string, props *ImageProps) error {
	return nil
}

