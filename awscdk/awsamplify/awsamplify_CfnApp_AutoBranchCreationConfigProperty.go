package awsamplify


// Use the AutoBranchCreationConfig property type to automatically create branches that match a certain pattern.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoBranchCreationConfigProperty := &autoBranchCreationConfigProperty{
//   	autoBranchCreationPatterns: []*string{
//   		jsii.String("autoBranchCreationPatterns"),
//   	},
//   	basicAuthConfig: &basicAuthConfigProperty{
//   		enableBasicAuth: jsii.Boolean(false),
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	buildSpec: jsii.String("buildSpec"),
//   	enableAutoBranchCreation: jsii.Boolean(false),
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
//   }
//
type CfnApp_AutoBranchCreationConfigProperty struct {
	// Automated branch creation glob patterns for the Amplify app.
	AutoBranchCreationPatterns *[]*string `field:"optional" json:"autoBranchCreationPatterns" yaml:"autoBranchCreationPatterns"`
	// Sets password protection for your auto created branch.
	BasicAuthConfig interface{} `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// The build specification (build spec) for the autocreated branch.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Enables automated branch creation for the Amplify app.
	EnableAutoBranchCreation interface{} `field:"optional" json:"enableAutoBranchCreation" yaml:"enableAutoBranchCreation"`
	// Enables auto building for the auto created branch.
	EnableAutoBuild interface{} `field:"optional" json:"enableAutoBuild" yaml:"enableAutoBuild"`
	// Enables performance mode for the branch.
	//
	// Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out.
	EnablePerformanceMode interface{} `field:"optional" json:"enablePerformanceMode" yaml:"enablePerformanceMode"`
	// Sets whether pull request previews are enabled for each branch that Amplify Console automatically creates for your app.
	//
	// Amplify Console creates previews by deploying your app to a unique URL whenever a pull request is opened for the branch. Development and QA teams can use this preview to test the pull request before it's merged into a production or integration branch.
	//
	// To provide backend support for your preview, the Amplify Console automatically provisions a temporary backend environment that it deletes when the pull request is closed. If you want to specify a dedicated backend environment for your previews, use the `PullRequestEnvironmentName` property.
	//
	// For more information, see [Web Previews](https://docs.aws.amazon.com/amplify/latest/userguide/pr-previews.html) in the *AWS Amplify Hosting User Guide* .
	EnablePullRequestPreview interface{} `field:"optional" json:"enablePullRequestPreview" yaml:"enablePullRequestPreview"`
	// Environment variables for the auto created branch.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// `CfnApp.AutoBranchCreationConfigProperty.Framework`.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// If pull request previews are enabled, you can use this property to specify a dedicated backend environment for your previews.
	//
	// For example, you could specify an environment named `prod` , `test` , or `dev` that you initialized with the Amplify CLI.
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
	// Stage for the auto created branch.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

