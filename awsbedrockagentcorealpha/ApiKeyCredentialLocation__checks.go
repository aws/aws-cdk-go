//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateApiKeyCredentialLocation_HeaderParameters(config *ApiKeyAdditionalConfiguration) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func validateApiKeyCredentialLocation_QueryParameterParameters(config *ApiKeyAdditionalConfiguration) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

