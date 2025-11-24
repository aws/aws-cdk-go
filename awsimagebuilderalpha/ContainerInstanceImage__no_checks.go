//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateContainerInstanceImage_FromAmiIdParameters(amiId *string) error {
	return nil
}

func validateContainerInstanceImage_FromSsmParameterParameters(parameter awsssm.IStringParameter) error {
	return nil
}

func validateContainerInstanceImage_FromSsmParameterNameParameters(parameterName *string) error {
	return nil
}

func validateContainerInstanceImage_FromStringParameters(containerInstanceImageString *string) error {
	return nil
}

func validateNewContainerInstanceImageParameters(image *string) error {
	return nil
}

