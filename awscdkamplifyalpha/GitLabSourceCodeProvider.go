package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// GitLab source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	SourceCodeProvider: amplify.NewGitLabSourceCodeProvider(&GitLabSourceCodeProviderProps{
//   		Owner: jsii.String("<user>"),
//   		Repository: jsii.String("<repo>"),
//   		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-gitlab-token")),
//   	}),
//   })
//
// Experimental.
type GitLabSourceCodeProvider interface {
	ISourceCodeProvider
	// Binds the source code provider to an app.
	// Experimental.
	Bind(_app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for GitLabSourceCodeProvider
type jsiiProxy_GitLabSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewGitLabSourceCodeProvider(props *GitLabSourceCodeProviderProps) GitLabSourceCodeProvider {
	_init_.Initialize()

	if err := validateNewGitLabSourceCodeProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitLabSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitLabSourceCodeProvider_Override(g GitLabSourceCodeProvider, props *GitLabSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProvider",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitLabSourceCodeProvider) Bind(_app App) *SourceCodeProviderConfig {
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

