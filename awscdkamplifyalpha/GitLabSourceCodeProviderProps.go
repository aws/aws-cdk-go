package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a GitLab source code provider.
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

