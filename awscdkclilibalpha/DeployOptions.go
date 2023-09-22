package awscdkclilibalpha


// Options to use with cdk deploy.
//
// Example:
//   // await this asynchronous method call using a language feature
//   cli.Deploy(&DeployOptions{
//   	Stacks: []*string{
//   		jsii.String("MyTestStack"),
//   	},
//   })
//
// Experimental.
type DeployOptions struct {
	// Include "aws:asset:*" CloudFormation metadata for resources that use assets.
	// Default: true.
	//
	// Experimental.
	AssetMetadata *bool `field:"optional" json:"assetMetadata" yaml:"assetMetadata"`
	// Path to CA certificate to use when validating HTTPS requests.
	// Default: - read from AWS_CA_BUNDLE environment variable.
	//
	// Experimental.
	CaBundlePath *string `field:"optional" json:"caBundlePath" yaml:"caBundlePath"`
	// Show colors and other style from console output.
	// Default: - `true` unless the environment variable `NO_COLOR` is set.
	//
	// Experimental.
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Additional context.
	// Default: - no additional context.
	//
	// Experimental.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// enable emission of additional debugging information, such as creation stack traces of tokens.
	// Default: false.
	//
	// Experimental.
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
	// Force trying to fetch EC2 instance credentials.
	// Default: - guess EC2 instance status.
	//
	// Experimental.
	Ec2Creds *bool `field:"optional" json:"ec2Creds" yaml:"ec2Creds"`
	// Ignores synthesis errors, which will likely produce an invalid output.
	// Default: false.
	//
	// Experimental.
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Use JSON output instead of YAML when templates are printed to STDOUT.
	// Default: false.
	//
	// Experimental.
	Json *bool `field:"optional" json:"json" yaml:"json"`
	// Perform context lookups.
	//
	// Synthesis fails if this is disabled and context lookups need
	// to be performed.
	// Default: true.
	//
	// Experimental.
	Lookups *bool `field:"optional" json:"lookups" yaml:"lookups"`
	// Show relevant notices.
	// Default: true.
	//
	// Experimental.
	Notices *bool `field:"optional" json:"notices" yaml:"notices"`
	// Include "aws:cdk:path" CloudFormation metadata for each resource.
	// Default: true.
	//
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// Use the indicated AWS profile as the default environment.
	// Default: - no profile is used.
	//
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Use the indicated proxy.
	//
	// Will read from
	// HTTPS_PROXY environment if specified.
	// Default: - no proxy.
	//
	// Experimental.
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Role to pass to CloudFormation for deployment.
	// Default: - use the bootstrap cfn-exec role.
	//
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of stacks to deploy.
	// Default: - all stacks.
	//
	// Experimental.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
	// Copy assets to the output directory.
	//
	// Needed for local debugging the source files with SAM CLI.
	// Default: false.
	//
	// Experimental.
	Staging *bool `field:"optional" json:"staging" yaml:"staging"`
	// Do not construct stacks with warnings.
	// Default: false.
	//
	// Experimental.
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// Print trace for stack warnings.
	// Default: false.
	//
	// Experimental.
	Trace *bool `field:"optional" json:"trace" yaml:"trace"`
	// show debug logs.
	// Default: false.
	//
	// Experimental.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Include "AWS::CDK::Metadata" resource in synthesized templates.
	// Default: true.
	//
	// Experimental.
	VersionReporting *bool `field:"optional" json:"versionReporting" yaml:"versionReporting"`
	// Whether to build/publish assets in parallel.
	// Default: false.
	//
	// Experimental.
	AssetParallelism *bool `field:"optional" json:"assetParallelism" yaml:"assetParallelism"`
	// Whether to build all assets before deploying the first stack (useful for failing Docker builds).
	// Default: true.
	//
	// Experimental.
	AssetPrebuild *bool `field:"optional" json:"assetPrebuild" yaml:"assetPrebuild"`
	// Optional name to use for the CloudFormation change set.
	//
	// If not provided, a name will be generated automatically.
	// Default: - auto generate a name.
	//
	// Experimental.
	ChangeSetName *string `field:"optional" json:"changeSetName" yaml:"changeSetName"`
	// Whether we are on a CI system.
	// Default: - `false` unless the environment variable `CI` is set.
	//
	// Experimental.
	Ci *bool `field:"optional" json:"ci" yaml:"ci"`
	// Maximum number of simultaneous deployments (dependency permitting) to execute.
	// Default: 1.
	//
	// Experimental.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Only perform action on the given stack.
	// Default: false.
	//
	// Experimental.
	Exclusively *bool `field:"optional" json:"exclusively" yaml:"exclusively"`
	// Whether to execute the ChangeSet Not providing `execute` parameter will result in execution of ChangeSet.
	// Default: true.
	//
	// Experimental.
	Execute *bool `field:"optional" json:"execute" yaml:"execute"`
	// Always deploy, even if templates are identical.
	// Default: false.
	//
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Experimental.
	Hotswap HotswapMode `field:"optional" json:"hotswap" yaml:"hotswap"`
	// ARNs of SNS topics that CloudFormation will notify with stack related events.
	// Default: - no notifications.
	//
	// Experimental.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Path to file where stack outputs will be written after a successful deploy as JSON.
	// Default: - Outputs are not written to any file.
	//
	// Experimental.
	OutputsFile *string `field:"optional" json:"outputsFile" yaml:"outputsFile"`
	// Additional parameters for CloudFormation at deploy time.
	// Default: {}.
	//
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Display mode for stack activity events.
	//
	// The default in the CLI is StackActivityProgress.BAR. But since this is an API
	// it makes more sense to set the default to StackActivityProgress.EVENTS
	// Default: StackActivityProgress.EVENTS
	//
	// Experimental.
	Progress StackActivityProgress `field:"optional" json:"progress" yaml:"progress"`
	// What kind of security changes require approval.
	// Default: RequireApproval.Never
	//
	// Experimental.
	RequireApproval RequireApproval `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Reuse the assets with the given asset IDs.
	// Default: - do not reuse assets.
	//
	// Experimental.
	ReuseAssets *[]*string `field:"optional" json:"reuseAssets" yaml:"reuseAssets"`
	// Rollback failed deployments.
	// Default: true.
	//
	// Experimental.
	Rollback *bool `field:"optional" json:"rollback" yaml:"rollback"`
	// Name of the toolkit stack to use/deploy.
	// Default: CDKToolkit.
	//
	// Experimental.
	ToolkitStackName *string `field:"optional" json:"toolkitStackName" yaml:"toolkitStackName"`
	// Use previous values for unspecified parameters.
	//
	// If not set, all parameters must be specified for every deployment.
	// Default: true.
	//
	// Experimental.
	UsePreviousParameters *bool `field:"optional" json:"usePreviousParameters" yaml:"usePreviousParameters"`
}

