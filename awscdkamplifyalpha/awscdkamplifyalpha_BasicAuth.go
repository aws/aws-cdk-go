// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Basic Auth configuration.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
//   	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
//   		owner: jsii.String("<user>"),
//   		repository: jsii.String("<repo>"),
//   		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	}),
//   	basicAuth: amplify.basicAuth.fromGeneratedPassword(jsii.String("username")),
//   })
//
// Experimental.
type BasicAuth interface {
	// Binds this Basic Auth configuration to an App.
	// Experimental.
	Bind(scope constructs.Construct, id *string) *BasicAuthConfig
}

// The jsii proxy struct for BasicAuth
type jsiiProxy_BasicAuth struct {
	_ byte // padding
}

// Experimental.
func NewBasicAuth(props *BasicAuthProps) BasicAuth {
	_init_.Initialize()

	if err := validateNewBasicAuthParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BasicAuth{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBasicAuth_Override(b BasicAuth, props *BasicAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		[]interface{}{props},
		b,
	)
}

// Creates a Basic Auth configuration from a username and a password.
// Experimental.
func BasicAuth_FromCredentials(username *string, password awscdk.SecretValue) BasicAuth {
	_init_.Initialize()

	if err := validateBasicAuth_FromCredentialsParameters(username, password); err != nil {
		panic(err)
	}
	var returns BasicAuth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		"fromCredentials",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Creates a Basic Auth configuration with a password generated in Secrets Manager.
// Experimental.
func BasicAuth_FromGeneratedPassword(username *string, encryptionKey awskms.IKey) BasicAuth {
	_init_.Initialize()

	if err := validateBasicAuth_FromGeneratedPasswordParameters(username); err != nil {
		panic(err)
	}
	var returns BasicAuth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		"fromGeneratedPassword",
		[]interface{}{username, encryptionKey},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuth) Bind(scope constructs.Construct, id *string) *BasicAuthConfig {
	if err := b.validateBindParameters(scope, id); err != nil {
		panic(err)
	}
	var returns *BasicAuthConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

