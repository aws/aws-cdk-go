package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a GitHub source code provider.
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
type GitHubSourceCodeProviderProps struct {
	// A personal access token with the `repo` scope.
	// Experimental.
	OauthToken awscdk.SecretValue `field:"required" json:"oauthToken" yaml:"oauthToken"`
	// The user or organization owning the repository.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository.
	// Experimental.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}

