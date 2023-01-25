package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
)

// Auto branch creation configuration.
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
type AutoBranchCreation struct {
	// Whether to enable auto building for the auto created branch.
	// Experimental.
	AutoBuild *bool `field:"optional" json:"autoBuild" yaml:"autoBuild"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection for
	// the auto created branch.
	// Experimental.
	BasicAuth BasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Build spec for the auto created branch.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Environment variables for the auto created branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Automated branch creation glob patterns.
	// Experimental.
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
	// The dedicated backend environment for the pull request previews of the auto created branch.
	// Experimental.
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Whether to enable pull request preview for the auto created branch.
	// Experimental.
	PullRequestPreview *bool `field:"optional" json:"pullRequestPreview" yaml:"pullRequestPreview"`
	// Stage for the auto created branch.
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

