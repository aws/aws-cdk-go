package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CodeCommit source code provider.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repository := codecommit.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	RepositoryName: jsii.String("my-repo"),
//   })
//
//   amplifyApp := amplify.NewApp(this, jsii.String("App"), &AppProps{
//   	SourceCodeProvider: amplify.NewCodeCommitSourceCodeProvider(&CodeCommitSourceCodeProviderProps{
//   		Repository: *Repository,
//   	}),
//   })
//
// Experimental.
type CodeCommitSourceCodeProvider interface {
	ISourceCodeProvider
	// Binds the source code provider to an app.
	// Experimental.
	Bind(app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for CodeCommitSourceCodeProvider
type jsiiProxy_CodeCommitSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewCodeCommitSourceCodeProvider(props *CodeCommitSourceCodeProviderProps) CodeCommitSourceCodeProvider {
	_init_.Initialize()

	if err := validateNewCodeCommitSourceCodeProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeCommitSourceCodeProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitSourceCodeProvider_Override(c CodeCommitSourceCodeProvider, props *CodeCommitSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProvider",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeCommitSourceCodeProvider) Bind(app App) *SourceCodeProviderConfig {
	if err := c.validateBindParameters(app); err != nil {
		panic(err)
	}
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{app},
		&returns,
	)

	return returns
}

