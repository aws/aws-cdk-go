// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a GitHub source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
//   	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
//   		owner: jsii.String("<user>"),
//   		repository: jsii.String("<repo>"),
//   		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	}),
//   	autoBranchCreation: &autoBranchCreation{
//   		 // Automatically connect branches that match a pattern set
//   		patterns: []*string{
//   			jsii.String("feature/*"),
//   			jsii.String("test/*"),
//   		},
//   	},
//   	autoBranchDeletion: jsii.Boolean(true),
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

