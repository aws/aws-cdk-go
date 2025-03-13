package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

// Auto branch creation configuration.
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
type AutoBranchCreation struct {
	// Whether to enable auto building for the auto created branch.
	// Default: true.
	//
	// Experimental.
	AutoBuild *bool `field:"optional" json:"autoBuild" yaml:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the auto created branch.
	// Default: - no password protection.
	//
	// Experimental.
	BasicAuth BasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Build spec for the auto created branch.
	// Default: - application build spec.
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Environment variables for the auto created branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Default: - application environment variables.
	//
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Automated branch creation glob patterns.
	// Default: - all repository branches.
	//
	// Experimental.
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
	// The dedicated backend environment for the pull request previews of the auto created branch.
	// Default: - automatically provision a temporary backend.
	//
	// Experimental.
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the auto created branch.
	// Default: true.
	//
	// Experimental.
	PullRequestPreview *bool `field:"optional" json:"pullRequestPreview" yaml:"pullRequestPreview"`
	// Stage for the auto created branch.
	// Default: - no stage.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

