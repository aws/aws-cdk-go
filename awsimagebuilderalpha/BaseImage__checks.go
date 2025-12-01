//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateBaseImage_FromAmiIdParameters(amiId *string) error {
	if amiId == nil {
		return fmt.Errorf("parameter amiId is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromImageParameters(image IImage) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromMarketplaceProductIdParameters(productId *string) error {
	if productId == nil {
		return fmt.Errorf("parameter productId is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromSsmParameterNameParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

func validateBaseImage_FromStringParameters(baseImageString *string) error {
	if baseImageString == nil {
		return fmt.Errorf("parameter baseImageString is required, but nil was provided")
	}

	return nil
}

func validateNewBaseImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

