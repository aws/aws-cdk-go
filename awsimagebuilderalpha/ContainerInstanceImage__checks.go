//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateContainerInstanceImage_FromAmiIdParameters(amiId *string) error {
	if amiId == nil {
		return fmt.Errorf("parameter amiId is required, but nil was provided")
	}

	return nil
}

func validateContainerInstanceImage_FromSsmParameterParameters(parameter awsssm.IStringParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateContainerInstanceImage_FromSsmParameterNameParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

func validateContainerInstanceImage_FromStringParameters(containerInstanceImageString *string) error {
	if containerInstanceImageString == nil {
		return fmt.Errorf("parameter containerInstanceImageString is required, but nil was provided")
	}

	return nil
}

func validateNewContainerInstanceImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

