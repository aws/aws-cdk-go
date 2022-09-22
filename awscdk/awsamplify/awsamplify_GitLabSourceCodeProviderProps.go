package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a GitLab source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
//   	sourceCodeProvider: amplify.NewGitLabSourceCodeProvider(&gitLabSourceCodeProviderProps{
//   		owner: jsii.String("<user>"),
//   		repository: jsii.String("<repo>"),
//   		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-gitlab-token")),
//   	}),
//   })
//
// Experimental.
type GitLabSourceCodeProviderProps struct {
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

