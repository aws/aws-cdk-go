//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

func (s *jsiiProxy_SchemaFile) validateBindParameters(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
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

