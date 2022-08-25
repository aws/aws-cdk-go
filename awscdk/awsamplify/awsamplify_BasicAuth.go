package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
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
	Bind(scope awscdk.Construct, id *string) *BasicAuthConfig
}

// The jsii proxy struct for BasicAuth
type jsiiProxy_BasicAuth struct {
	_ byte // padding
}

// Experimental.
func NewBasicAuth(props *BasicAuthProps) BasicAuth {
	_init_.Initialize()

	j := jsiiProxy_BasicAuth{}

	_jsii_.Create(
		"monocdk.aws_amplify.BasicAuth",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBasicAuth_Override(b BasicAuth, props *BasicAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.BasicAuth",
		[]interface{}{props},
		b,
	)
}

// Creates a Basic Auth configuration from a username and a password.
// Experimental.
func BasicAuth_FromCredentials(username *string, password awscdk.SecretValue) BasicAuth {
	_init_.Initialize()

	var returns BasicAuth

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.BasicAuth",
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

	var returns BasicAuth

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.BasicAuth",
		"fromGeneratedPassword",
		[]interface{}{username, encryptionKey},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BasicAuth) Bind(scope awscdk.Construct, id *string) *BasicAuthConfig {
	var returns *BasicAuthConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

