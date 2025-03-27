//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateParamsAndSecretsLayerVersion_FromVersionParameters(version ParamsAndSecretsVersions, options *ParamsAndSecretsOptions) error {
	if version == "" {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateParamsAndSecretsLayerVersion_FromVersionArnParameters(arn *string, options *ParamsAndSecretsOptions) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

