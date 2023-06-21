package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Identity providers supported by the UserPoolClient.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	// ...
//   	SupportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   		cognito.*userPoolClientIdentityProvider_COGNITO(),
//   	},
//   })
//
type UserPoolClientIdentityProvider interface {
	// The name of the identity provider as recognized by CloudFormation property `SupportedIdentityProviders`.
	Name() *string
}

// The jsii proxy struct for UserPoolClientIdentityProvider
type jsiiProxy_UserPoolClientIdentityProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_UserPoolClientIdentityProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Specify a provider not yet supported by the CDK.
func UserPoolClientIdentityProvider_Custom(name *string) UserPoolClientIdentityProvider {
	_init_.Initialize()

	if err := validateUserPoolClientIdentityProvider_CustomParameters(name); err != nil {
		panic(err)
	}
	var returns UserPoolClientIdentityProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"custom",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func UserPoolClientIdentityProvider_AMAZON() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"AMAZON",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_APPLE() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"APPLE",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_COGNITO() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"COGNITO",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_FACEBOOK() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"FACEBOOK",
		&returns,
	)
	return returns
}

func UserPoolClientIdentityProvider_GOOGLE() UserPoolClientIdentityProvider {
	_init_.Initialize()
	var returns UserPoolClientIdentityProvider
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPoolClientIdentityProvider",
		"GOOGLE",
		&returns,
	)
	return returns
}

