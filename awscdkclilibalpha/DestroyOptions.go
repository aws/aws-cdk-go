package awscdkclilibalpha


// Options to use with cdk destroy.
// Experimental.
type DestroyOptions struct {
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
	// Only destroy the given stack.
	// Default: false.
	//
	// Experimental.
	Exclusively *bool `field:"optional" json:"exclusively" yaml:"exclusively"`
	// Should the script prompt for approval before destroying stacks.
	// Default: false.
	//
	// Experimental.
	RequireApproval *bool `field:"optional" json:"requireApproval" yaml:"requireApproval"`
}

