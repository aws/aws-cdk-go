package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBranch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBranchProps := &cfnBranchProps{
//   	appId: jsii.String("appId"),
//   	branchName: jsii.String("branchName"),
//
//   	// the properties below are optional
//   	basicAuthConfig: &basicAuthConfigProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//
//   		// the properties below are optional
//   		enableBasicAuth: jsii.Boolean(false),
//   	},
//   	buildSpec: jsii.String("buildSpec"),
//   	description: jsii.String("description"),
//   	enableAutoBuild: jsii.Boolean(false),
//   	enablePerformanceMode: jsii.Boolean(false),
//   	enablePullRequestPreview: jsii.Boolean(false),
//   	environmentVariables: []interface{}{
//   		&environmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	framework: jsii.String("framework"),
//   	pullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   	stage: jsii.String("stage"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBranchProps struct {
	// The unique ID for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 20.
	//
	// *Pattern:* d[a-z0-9]+.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The name for the branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The basic authorization credentials for a branch of an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig interface{} `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for the branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	//
	// *Pattern:* (?s).+
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// The description for the branch that is part of an Amplify app.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enables auto building for the branch.
	EnableAutoBuild interface{} `field:"optional" json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out.
	EnablePerformanceMode interface{} `field:"optional" json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Sets whether the Amplify Console creates a preview for each pull request that is made for this branch.
	//
	// If this property is enabled, the Amplify Console deploys your app to a unique preview URL after each pull request is opened. Development and QA teams can use this preview to test the pull request before it's merged into a production or integration branch.
	//
	// To provide backend support for your preview, the Amplify Console automatically provisions a temporary backend environment that it deletes when the pull request is closed. If you want to specify a dedicated backend environment for your previews, use the `PullRequestEnvironmentName` property.
	//
	// For more information, see [Web Previews](https://docs.aws.amazon.com/amplify/latest/userguide/pr-previews.html) in the *AWS Amplify Hosting User Guide* .
	EnablePullRequestPreview interface{} `field:"optional" json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// The environment variables for the branch.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// `AWS::Amplify::Branch.Framework`.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// If pull request previews are enabled for this branch, you can use this property to specify a dedicated backend environment for your previews.
	//
	// For example, you could specify an environment named `prod` , `test` , or `dev` that you initialized with the Amplify CLI and mapped to this branch.
	//
	// To enable pull request previews, set the `EnablePullRequestPreview` property to `true` .
	//
	// If you don't specify an environment, the Amplify Console provides backend support for each preview by automatically provisioning a temporary backend environment. Amplify Console deletes this environment when the pull request is closed.
	//
	// For more information about creating backend environments, see [Feature Branch Deployments and Team Workflows](https://docs.aws.amazon.com/amplify/latest/userguide/multi-environments.html) in the *AWS Amplify Hosting User Guide* .
	//
	// *Length Constraints:* Maximum length of 20.
	//
	// *Pattern:* (?s).*
	PullRequestEnvironmentName *string `field:"optional" json:"pullRequestEnvironmentName" yaml:"pullRequestEnvironmentName"`
	// Describes the current stage for the branch.
	//
	// *Valid Values:* PRODUCTION | BETA | DEVELOPMENT | EXPERIMENTAL | PULL_REQUEST.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The tag for the branch.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

