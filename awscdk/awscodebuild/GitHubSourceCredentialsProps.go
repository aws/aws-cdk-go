package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Creation properties for `GitHubSourceCredentials`.
//
// Example:
//   codebuild.NewGitHubSourceCredentials(this, jsii.String("CodeBuildGitHubCreds"), &GitHubSourceCredentialsProps{
//   	AccessToken: awscdk.SecretValue_SecretsManager(jsii.String("my-token")),
//   })
//
type GitHubSourceCredentialsProps struct {
	// The personal access token to use when contacting the GitHub API.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
}

