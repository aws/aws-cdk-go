//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SchemaFile) validateBindParameters(api IGraphqlApi, _options *SchemaBindOptions) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func validateSchemaFile_FromAssetParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SchemaFile) validateSetDefinitionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSchemaFileParameters(options *SchemaProps) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

