//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateGatewayCredentialProvider_FromApiKeyIdentityParameters(provider IApiKeyCredentialProvider, options *FromApiKeyIdentityOptions) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGatewayCredentialProvider_FromApiKeyIdentityArnParameters(props *ApiKeyCredentialProviderProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateGatewayCredentialProvider_FromOauthIdentityParameters(provider IOAuth2CredentialProvider, options *FromOauthIdentityOptions) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGatewayCredentialProvider_FromOauthIdentityArnParameters(props *OAuthConfiguration) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

