package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Exposes methods for defining authorization config for AppSync APIs.
type IAppSyncAuthConfig interface {
	// Set up Cognito Authorization configuration for AppSync APIs.
	SetupCognitoConfig(config *AppSyncCognitoConfig) interface{}
	// Set up Lambda Authorization configuration AppSync APIs.
	SetupLambdaAuthorizerConfig(config *AppSyncLambdaAuthorizerConfig) interface{}
	// Set up OIDC Authorization configuration for AppSync APIs.
	SetupOpenIdConnectConfig(config *AppSyncOpenIdConnectConfig) interface{}
}

// The jsii proxy for IAppSyncAuthConfig
type jsiiProxy_IAppSyncAuthConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_IAppSyncAuthConfig) SetupCognitoConfig(config *AppSyncCognitoConfig) interface{} {
	if err := i.validateSetupCognitoConfigParameters(config); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"setupCognitoConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAppSyncAuthConfig) SetupLambdaAuthorizerConfig(config *AppSyncLambdaAuthorizerConfig) interface{} {
	if err := i.validateSetupLambdaAuthorizerConfigParameters(config); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"setupLambdaAuthorizerConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAppSyncAuthConfig) SetupOpenIdConnectConfig(config *AppSyncOpenIdConnectConfig) interface{} {
	if err := i.validateSetupOpenIdConnectConfigParameters(config); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"setupOpenIdConnectConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

