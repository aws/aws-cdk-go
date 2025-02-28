package awscdkcognitoidentitypoolalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Keys for Login Providers - each correspond to the client IDs of their respective federation Identity Providers.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	RoleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			ProviderUrl: awscdkcognitoidentitypoolalpha.IdentityPoolProviderUrl_Custom(jsii.String("my-custom-provider.com")),
//   			UseToken: jsii.Boolean(true),
//   		},
//   	},
//   })
//
// Experimental.
type IdentityPoolProviderUrl interface {
	// The type of Identity Pool Provider.
	// Experimental.
	Type() IdentityPoolProviderType
	// The value of the Identity Pool Provider.
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

// Custom Provider url.
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

// OpenId Provider url.
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

// Saml Provider url.
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
func IdentityPoolProviderUrl_UserPool(userPool awscognito.IUserPool, userPoolClient awscognito.IUserPoolClient) IdentityPoolProviderUrl {
	_init_.Initialize()

	if err := validateIdentityPoolProviderUrl_UserPoolParameters(userPool, userPoolClient); err != nil {
		panic(err)
	}
	var returns IdentityPoolProviderUrl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		"userPool",
		[]interface{}{userPool, userPoolClient},
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

