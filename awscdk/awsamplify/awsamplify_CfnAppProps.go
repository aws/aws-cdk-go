package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &cfnAppProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	autoBranchCreationConfig: &autoBranchCreationConfigProperty{
//   		autoBranchCreationPatterns: []*string{
//   			jsii.String("autoBranchCreationPatterns"),
//   		},
//   		basicAuthConfig: &basicAuthConfigProperty{
//   			enableBasicAuth: jsii.Boolean(false),
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		buildSpec: jsii.String("buildSpec"),
//   		enableAutoBranchCreation: jsii.Boolean(false),
//   		enableAutoBuild: jsii.Boolean(false),
//   		enablePerformanceMode: jsii.Boolean(false),
//   		enablePullRequestPreview: jsii.Boolean(false),
//   		environmentVariables: []interface{}{
//   			&environmentVariableProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		pullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   		stage: jsii.String("stage"),
//   	},
//   	basicAuthConfig: &basicAuthConfigProperty{
//   		enableBasicAuth: jsii.Boolean(false),
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	buildSpec: jsii.String("buildSpec"),
//   	customHeaders: jsii.String("customHeaders"),
//   	customRules: []interface{}{
//   		&customRuleProperty{
//   			source: jsii.String("source"),
//   			target: jsii.String("target"),
//
//   			// the properties below are optional
//   			condition: jsii.String("condition"),
//   			status: jsii.String("status"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	enableBranchAutoDeletion: jsii.Boolean(false),
//   	environmentVariables: []interface{}{
//   		&environmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	iamServiceRole: jsii.String("iamServiceRole"),
//   	oauthToken: jsii.String("oauthToken"),
//   	repository: jsii.String("repository"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppProps struct {
	// The name for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	Name *string `field:"required" json:"name" yaml:"name"`
	// The personal access token for a GitHub repository for an Amplify app.
	//
	// The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig interface{} `field:"optional" json:"autoBranchCreationConfig" yaml:"autoBranchCreationConfig"`
	// The credentials for basic authorization for an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig interface{} `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	//
	// *Pattern:* (?s).+
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// The custom HTTP headers for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 25000.
	//
	// *Pattern:* (?s).*
	CustomHeaders *string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules interface{} `field:"optional" json:"customRules" yaml:"customRules"`
	// The description for an Amplify app.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion interface{} `field:"optional" json:"enableBranchAutoDeletion" yaml:"enableBranchAutoDeletion"`
	// The environment variables map for an Amplify app.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) of the Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	IamServiceRole *string `field:"optional" json:"iamServiceRole" yaml:"iamServiceRole"`
	// The OAuth token for a third-party source control system for an Amplify app.
	//
	// The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// The repository for an Amplify app.
	//
	// *Pattern:* (?s).*
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The tag for an Amplify app.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

