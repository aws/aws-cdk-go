//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func validateMappingTemplate_DynamoDbDeleteItemParameters(keyName *string, idArg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if idArg == nil {
		return fmt.Errorf("parameter idArg is required, but nil was provided")
	}

	return nil
}

func validateMappingTemplate_DynamoDbGetItemParameters(keyName *string, idArg *string) error {
	if keyName == nil {
		return fmt.Errorf("parameter keyName is required, but nil was provided")
	}

	if idArg == nil {
		return fmt.Errorf("parameter idArg is required, but nil was provided")
	}

	return nil
}

func validateMappingTemplate_DynamoDbPutItemParameters(key PrimaryKey, values AttributeValues) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateMappingTemplate_DynamoDbQueryParameters(cond KeyCondition) error {
	if cond == nil {
		return fmt.Errorf("parameter cond is required, but nil was provided")
	}

	return nil
}

func validateMappingTemplate_FromFileParameters(fileName *string) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	return nil
}

func validateMappingTemplate_FromStringParameters(template *string) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	return nil
}

