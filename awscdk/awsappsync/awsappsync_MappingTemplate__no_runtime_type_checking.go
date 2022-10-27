//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateMappingTemplate_DynamoDbDeleteItemParameters(keyName *string, idArg *string) error {
	return nil
}

func validateMappingTemplate_DynamoDbGetItemParameters(keyName *string, idArg *string) error {
	return nil
}

func validateMappingTemplate_DynamoDbPutItemParameters(key PrimaryKey, values AttributeValues) error {
	return nil
}

func validateMappingTemplate_DynamoDbQueryParameters(cond KeyCondition) error {
	return nil
}

func validateMappingTemplate_FromFileParameters(fileName *string) error {
	return nil
}

func validateMappingTemplate_FromStringParameters(template *string) error {
	return nil
}

