package cloudassemblyschema


// Options to use with cdk deploy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployOptions := &DeployOptions{
//   	All: jsii.Boolean(false),
//   	App: jsii.String("app"),
//   	AssetMetadata: jsii.Boolean(false),
//   	CaBundlePath: jsii.String("caBundlePath"),
//   	ChangeSetName: jsii.String("changeSetName"),
//   	Ci: jsii.Boolean(false),
//   	Color: jsii.Boolean(false),
//   	Concurrency: jsii.Number(123),
//   	Context: map[string]*string{
//   		"contextKey": jsii.String("context"),
//   	},
//   	Debug: jsii.Boolean(false),
//   	Ec2Creds: jsii.Boolean(false),
//   	Exclusively: jsii.Boolean(false),
//   	Execute: jsii.Boolean(false),
//   	Force: jsii.Boolean(false),
//   	IgnoreErrors: jsii.Boolean(false),
//   	Json: jsii.Boolean(false),
//   	Lookups: jsii.Boolean(false),
//   	Notices: jsii.Boolean(false),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	Output: jsii.String("output"),
//   	OutputsFile: jsii.String("outputsFile"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	PathMetadata: jsii.Boolean(false),
//   	Profile: jsii.String("profile"),
//   	Proxy: jsii.String("proxy"),
//   	RequireApproval: awscdk.Cloud_assembly_schema.RequireApproval_NEVER,
//   	ReuseAssets: []*string{
//   		jsii.String("reuseAssets"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Rollback: jsii.Boolean(false),
//   	Stacks: []*string{
//   		jsii.String("stacks"),
//   	},
//   	Staging: jsii.Boolean(false),
//   	Strict: jsii.Boolean(false),
//   	ToolkitStackName: jsii.String("toolkitStackName"),
//   	Trace: jsii.Boolean(false),
//   	UsePreviousParameters: jsii.Boolean(false),
//   	Verbose: jsii.Boolean(false),
//   	VersionReporting: jsii.Boolean(false),
//   }
//
type DeployOptions struct {
	// Deploy all stacks.
	//
	// Requried if `stacks` is not set.
	// Default: - false.
	//
	All *bool `field:"optional" json:"all" yaml:"all"`
	// command-line for executing your app or a cloud assembly directory e.g. "node bin/my-app.js" or "cdk.out".
	// Default: - read from cdk.json
	//
	App *string `field:"optional" json:"app" yaml:"app"`
	// Include "aws:asset:*" CloudFormation metadata for resources that use assets.
	// Default: true.
	//
	AssetMetadata *bool `field:"optional" json:"assetMetadata" yaml:"assetMetadata"`
	// Path to CA certificate to use when validating HTTPS requests.
	// Default: - read from AWS_CA_BUNDLE environment variable.
	//
	CaBundlePath *string `field:"optional" json:"caBundlePath" yaml:"caBundlePath"`
	// Show colors and other style from console output.
	// Default: true.
	//
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Additional context.
	// Default: - no additional context.
	//
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// enable emission of additional debugging information, such as creation stack traces of tokens.
	// Default: false.
	//
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
	// Force trying to fetch EC2 instance credentials.
	// Default: - guess EC2 instance status.
	//
	Ec2Creds *bool `field:"optional" json:"ec2Creds" yaml:"ec2Creds"`
	// Ignores synthesis errors, which will likely produce an invalid output.
	// Default: false.
	//
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Use JSON output instead of YAML when templates are printed to STDOUT.
	// Default: false.
	//
	Json *bool `field:"optional" json:"json" yaml:"json"`
	// Perform context lookups.
	//
	// Synthesis fails if this is disabled and context lookups need
	// to be performed.
	// Default: true.
	//
	Lookups *bool `field:"optional" json:"lookups" yaml:"lookups"`
	// Show relevant notices.
	// Default: true.
	//
	Notices *bool `field:"optional" json:"notices" yaml:"notices"`
	// Emits the synthesized cloud assembly into a directory.
	// Default: cdk.out
	//
	Output *string `field:"optional" json:"output" yaml:"output"`
	// Include "aws:cdk:path" CloudFormation metadata for each resource.
	// Default: true.
	//
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// Use the indicated AWS profile as the default environment.
	// Default: - no profile is used.
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Use the indicated proxy.
	//
	// Will read from
	// HTTPS_PROXY environment if specified.
	// Default: - no proxy.
	//
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Role to pass to CloudFormation for deployment.
	// Default: - use the bootstrap cfn-exec role.
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of stacks to deploy.
	//
	// Requried if `all` is not set.
	// Default: - [].
	//
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
	// Copy assets to the output directory.
	//
	// Needed for local debugging the source files with SAM CLI.
	// Default: false.
	//
	Staging *bool `field:"optional" json:"staging" yaml:"staging"`
	// Do not construct stacks with warnings.
	// Default: false.
	//
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// Print trace for stack warnings.
	// Default: false.
	//
	Trace *bool `field:"optional" json:"trace" yaml:"trace"`
	// show debug logs.
	// Default: false.
	//
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Include "AWS::CDK::Metadata" resource in synthesized templates.
	// Default: true.
	//
	VersionReporting *bool `field:"optional" json:"versionReporting" yaml:"versionReporting"`
	// Optional name to use for the CloudFormation change set.
	//
	// If not provided, a name will be generated automatically.
	// Default: - auto generate a name.
	//
	ChangeSetName *string `field:"optional" json:"changeSetName" yaml:"changeSetName"`
	// Whether we are on a CI system.
	// Default: false.
	//
	Ci *bool `field:"optional" json:"ci" yaml:"ci"`
	// Deploy multiple stacks in parallel.
	// Default: 1.
	//
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Only perform action on the given stack.
	// Default: false.
	//
	Exclusively *bool `field:"optional" json:"exclusively" yaml:"exclusively"`
	// Whether to execute the ChangeSet Not providing `execute` parameter will result in execution of ChangeSet.
	// Default: true.
	//
	Execute *bool `field:"optional" json:"execute" yaml:"execute"`
	// Always deploy, even if templates are identical.
	// Default: false.
	//
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// ARNs of SNS topics that CloudFormation will notify with stack related events.
	// Default: - no notifications.
	//
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Path to file where stack outputs will be written after a successful deploy as JSON.
	// Default: - Outputs are not written to any file.
	//
	OutputsFile *string `field:"optional" json:"outputsFile" yaml:"outputsFile"`
	// Additional parameters for CloudFormation at deploy time.
	// Default: {}.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// What kind of security changes require approval.
	// Default: RequireApproval.Never
	//
	RequireApproval RequireApproval `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Reuse the assets with the given asset IDs.
	// Default: - do not reuse assets.
	//
	ReuseAssets *[]*string `field:"optional" json:"reuseAssets" yaml:"reuseAssets"`
	// Rollback failed deployments.
	// Default: true.
	//
	Rollback *bool `field:"optional" json:"rollback" yaml:"rollback"`
	// Name of the toolkit stack to use/deploy.
	// Default: CDKToolkit.
	//
	ToolkitStackName *string `field:"optional" json:"toolkitStackName" yaml:"toolkitStackName"`
	// Use previous values for unspecified parameters.
	//
	// If not set, all parameters must be specified for every deployment.
	// Default: true.
	//
	UsePreviousParameters *bool `field:"optional" json:"usePreviousParameters" yaml:"usePreviousParameters"`
}

