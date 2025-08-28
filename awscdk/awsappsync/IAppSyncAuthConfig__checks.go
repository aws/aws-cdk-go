//go:build !no_runtime_type_checking

package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupCognitoConfigParameters(config *AppSyncCognitoConfig) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupLambdaAuthorizerConfigParameters(config *AppSyncLambdaAuthorizerConfig) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupOpenIdConnectConfigParameters(config *AppSyncOpenIdConnectConfig) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

