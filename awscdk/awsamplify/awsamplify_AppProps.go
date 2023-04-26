package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for an App.
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
type AppProps struct {
	// The name for the application.
	// Experimental.
	AppName *string `field:"optional" json:"appName" yaml:"appName"`
	// The auto branch creation configuration.
	//
	// Use this to automatically create
	// branches that match a certain pattern.
	// Experimental.
	AutoBranchCreation *AutoBranchCreation `field:"optional" json:"autoBranchCreation" yaml:"autoBranchCreation"`
	// Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.
	// Experimental.
	AutoBranchDeletion *bool `field:"optional" json:"autoBranchDeletion" yaml:"autoBranchDeletion"`
	// The Basic Auth configuration.
	//
	// Use this to set password protection at an
	// app level to all your branches.
	// Experimental.
	BasicAuth BasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// BuildSpec for the application.
	//
	// Alternatively, add a `amplify.yml`
	// file to the repository.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html
	//
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// The custom HTTP response headers for an Amplify app.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/custom-headers.html
	//
	// Experimental.
	CustomResponseHeaders *[]*CustomResponseHeader `field:"optional" json:"customResponseHeaders" yaml:"customResponseHeaders"`
	// Custom rewrite/redirect rules for the application.
	// Experimental.
	CustomRules *[]CustomRule `field:"optional" json:"customRules" yaml:"customRules"`
	// A description for the application.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables for the application.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The IAM service role to associate with the application.
	//
	// The App
	// implements IGrantable.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The source code provider for this application.
	// Experimental.
	SourceCodeProvider ISourceCodeProvider `field:"optional" json:"sourceCodeProvider" yaml:"sourceCodeProvider"`
}

