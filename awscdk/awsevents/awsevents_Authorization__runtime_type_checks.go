//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateAuthorization_ApiKeyParameters(apiKeyName *string, apiKeyValue awscdk.SecretValue) error {
	if apiKeyName == nil {
		return fmt.Errorf("parameter apiKeyName is required, but nil was provided")
	}

	if apiKeyValue == nil {
		return fmt.Errorf("parameter apiKeyValue is required, but nil was provided")
	}

	return nil
}

func validateAuthorization_BasicParameters(username *string, password awscdk.SecretValue) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if password == nil {
		return fmt.Errorf("parameter password is required, but nil was provided")
	}

	return nil
}

func validateAuthorization_OauthParameters(props *OAuthAuthorizationProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

