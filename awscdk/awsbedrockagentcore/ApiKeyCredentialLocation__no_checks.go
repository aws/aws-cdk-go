//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateApiKeyCredentialLocation_HeaderParameters(config *ApiKeyAdditionalConfiguration) error {
	return nil
}

func validateApiKeyCredentialLocation_QueryParameterParameters(config *ApiKeyAdditionalConfiguration) error {
	return nil
}

