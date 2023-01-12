// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Keys for Login Providers - correspond to client id's of respective federation identity providers.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   var userPool userPool
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &identityPoolProps{
//   	identityPoolName: jsii.String("myidentitypool"),
//   	roleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			mappingKey: jsii.String("cognito"),
//   			providerUrl: awscdkcognitoidentitypoolalpha.IdentityPoolProviderUrl.userPool(userPool.userPoolProviderUrl),
//   			useToken: jsii.Boolean(true),
//   		},
//   	},
//   })
//
// Experimental.
type IdentityPoolProviderUrl interface {
	// type of Provider Url.
	// Experimental.
	Type() IdentityPoolProviderType
	// value of Provider Url.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for IdentityPoolProviderUrl
type jsiiProxy_IdentityPoolProviderUrl struct {
	_ byte // padding
}

func (j *jsiiProxy_IdentityPoolProviderUrl) Type() IdentityPoolProviderType {
	var returns IdentityPoolProviderType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolProviderUrl) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewIdentityPoolProviderUrl(type_ IdentityPoolProviderType, value *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateNewIdentityPoolProviderUrlParameters(type_, value); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPoolProviderUrl{}

	_jsii_.Create(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		[]interface{}{type_, value},
		&j,
	)

	return &j
}

// Experimental.
func NewIdentityPoolProviderUrl_Override(i IdentityPoolProviderUrl, type_ IdentityPoolProviderType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		[]interface{}{type_, value},
		i,
	)
}

// Custom Provider Url.
// Experimental.
func IdentityPoolProviderUrl_Custom(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateIdentityPoolProviderUrl_CustomParameters(url); err != nil {
		panic(err)
	}
	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"custom",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// OpenId Provider Url.
// Experimental.
func IdentityPoolProviderUrl_OpenId(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateIdentityPoolProviderUrl_OpenIdParameters(url); err != nil {
		panic(err)
	}
	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"openId",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// Saml Provider Url.
// Experimental.
func IdentityPoolProviderUrl_Saml(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateIdentityPoolProviderUrl_SamlParameters(url); err != nil {
		panic(err)
	}
	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"saml",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// User Pool Provider Url.
// Experimental.
func IdentityPoolProviderUrl_UserPool(url *string) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateIdentityPoolProviderUrl_UserPoolParameters(url); err != nil {
		panic(err)
	}
	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"userPool",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func IdentityPoolProviderUrl_AMAZON() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"AMAZON",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_APPLE() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"APPLE",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_DIGITS() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"DIGITS",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_FACEBOOK() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"FACEBOOK",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_GOOGLE() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"GOOGLE",
		&returns,
	)
	return returns
}

func IdentityPoolProviderUrl_TWITTER() IdentityPoolProviderUrl {
	_init_.Initialize()
	var returns IdentityPoolProviderUrl
	_jsii_.StaticGet(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"TWITTER",
		&returns,
	)
	return returns
}

