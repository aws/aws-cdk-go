package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// GitHub source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("App"), &AppProps{
//   	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
//   		Owner: jsii.String("<user>"),
//   		Repository: jsii.String("<repo>"),
//   		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
//   	}),
//   	CustomResponseHeaders: []CustomResponseHeader{
//   		&CustomResponseHeader{
//   			Pattern: jsii.String("*.json"),
//   			Headers: map[string]*string{
//   				"custom-header-name-1": jsii.String("custom-header-value-1"),
//   				"custom-header-name-2": jsii.String("custom-header-value-2"),
//   			},
//   		},
//   		&CustomResponseHeader{
//   			Pattern: jsii.String("/path/*"),
//   			Headers: map[string]*string{
//   				"custom-header-name-1": jsii.String("custom-header-value-2"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type GitHubSourceCodeProvider interface {
	ISourceCodeProvider
	// Binds the source code provider to an app.
	// Experimental.
	Bind(_app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for GitHubSourceCodeProvider
type jsiiProxy_GitHubSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewGitHubSourceCodeProvider(props *GitHubSourceCodeProviderProps) GitHubSourceCodeProvider {
	_init_.Initialize()

	if err := validateNewGitHubSourceCodeProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubSourceCodeProvider_Override(g GitHubSourceCodeProvider, props *GitHubSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProvider",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitHubSourceCodeProvider) Bind(_app App) *SourceCodeProviderConfig {
	if err := g.validateBindParameters(_app); err != nil {
		panic(err)
	}
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_app},
		&returns,
	)

	return returns
}

