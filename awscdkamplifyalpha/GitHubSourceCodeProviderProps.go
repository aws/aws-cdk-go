// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a GitHub source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
//   		Owner: jsii.String("<user>"),
//   		Repository: jsii.String("<repo>"),
//   		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
//   	}),
//   	AutoBranchCreation: &AutoBranchCreation{
//   		 // Automatically connect branches that match a pattern set
//   		Patterns: []*string{
//   			jsii.String("feature/*"),
//   			jsii.String("test/*"),
//   		},
//   	},
//   	AutoBranchDeletion: jsii.Boolean(true),
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

