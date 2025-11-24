//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateBaseImage_FromAmiIdParameters(amiId *string) error {
	return nil
}

func validateBaseImage_FromMarketplaceProductIdParameters(productId *string) error {
	return nil
}

func validateBaseImage_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateBaseImage_FromSsmParameterNameParameters(parameterName *string) error {
	return nil
}

func validateBaseImage_FromStringParameters(baseImageString *string) error {
	return nil
}

func validateNewBaseImageParameters(image *string) error {
	return nil
}

