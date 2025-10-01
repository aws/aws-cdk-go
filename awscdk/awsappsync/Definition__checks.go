//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDefinition_FromFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateDefinition_FromSchemaParameters(schema ISchema) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

func validateDefinition_FromSourceApisParameters(sourceApiOptions *SourceApiOptions) error {
	if sourceApiOptions == nil {
		return fmt.Errorf("parameter sourceApiOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sourceApiOptions, func() string { return "parameter sourceApiOptions" }); err != nil {
		return err
	}

	return nil
}

