//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupCognitoConfigParameters(config *AppSyncCognitoConfig) error {
	return nil
}

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupLambdaAuthorizerConfigParameters(config *AppSyncLambdaAuthorizerConfig) error {
	return nil
}

func (i *jsiiProxy_IAppSyncAuthConfig) validateSetupOpenIdConnectConfigParameters(config *AppSyncOpenIdConnectConfig) error {
	return nil
}

