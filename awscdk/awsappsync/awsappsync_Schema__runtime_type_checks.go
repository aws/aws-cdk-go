//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_Schema) validateAddMutationParameters(fieldName *string, field ResolvableField) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Schema) validateAddQueryParameters(fieldName *string, field ResolvableField) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Schema) validateAddSubscriptionParameters(fieldName *string, field Field) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Schema) validateAddToSchemaParameters(addition *string) error {
	if addition == nil {
		return fmt.Errorf("parameter addition is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Schema) validateAddTypeParameters(type_ IIntermediateType) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_Schema) validateBindParameters(api GraphqlApi) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func validateSchema_FromAssetParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Schema) validateSetDefinitionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSchemaParameters(options *SchemaOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

