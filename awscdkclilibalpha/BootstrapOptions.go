package awscdkclilibalpha


// Options to use with cdk bootstrap.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cli_lib_alpha "github.com/aws/aws-cdk-go/awscdkclilibalpha"
//
//   bootstrapOptions := &BootstrapOptions{
//   	AssetMetadata: jsii.Boolean(false),
//   	BootstrapBucketName: jsii.String("bootstrapBucketName"),
//   	BootstrapCustomerKey: jsii.String("bootstrapCustomerKey"),
//   	BootstrapKmsKeyId: jsii.String("bootstrapKmsKeyId"),
//   	CaBundlePath: jsii.String("caBundlePath"),
//   	CfnExecutionPolicy: jsii.String("cfnExecutionPolicy"),
//   	Color: jsii.Boolean(false),
//   	Context: map[string]*string{
//   		"contextKey": jsii.String("context"),
//   	},
//   	CustomPermissionsBoundary: jsii.String("customPermissionsBoundary"),
//   	Debug: jsii.Boolean(false),
//   	Ec2Creds: jsii.Boolean(false),
//   	Environments: []*string{
//   		jsii.String("environments"),
//   	},
//   	ExamplePermissionsBoundary: jsii.Boolean(false),
//   	Execute: jsii.Boolean(false),
//   	Force: jsii.Boolean(false),
//   	IgnoreErrors: jsii.Boolean(false),
//   	Json: jsii.Boolean(false),
//   	Lookups: jsii.Boolean(false),
//   	Notices: jsii.Boolean(false),
//   	PathMetadata: jsii.Boolean(false),
//   	Profile: jsii.String("profile"),
//   	Proxy: jsii.String("proxy"),
//   	PublicAccessBlockConfiguration: jsii.String("publicAccessBlockConfiguration"),
//   	Qualifier: jsii.String("qualifier"),
//   	RoleArn: jsii.String("roleArn"),
//   	ShowTemplate: jsii.Boolean(false),
//   	Stacks: []*string{
//   		jsii.String("stacks"),
//   	},
//   	Staging: jsii.Boolean(false),
//   	Strict: jsii.Boolean(false),
//   	Template: jsii.String("template"),
//   	TerminationProtection: jsii.Boolean(false),
//   	ToolkitStackName: jsii.String("toolkitStackName"),
//   	Trace: jsii.Boolean(false),
//   	Trust: jsii.String("trust"),
//   	TrustForLookup: jsii.String("trustForLookup"),
//   	UsePreviousParameters: jsii.Boolean(false),
//   	Verbose: jsii.Boolean(false),
//   	VersionReporting: jsii.Boolean(false),
//   }
//
// Experimental.
type BootstrapOptions struct {
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
	// The name of the CDK toolkit bucket;
	//
	// bucket will be created and
	// must not exist.
	// Default: - auto-generated CloudFormation name.
	//
	// Experimental.
	BootstrapBucketName *string `field:"optional" json:"bootstrapBucketName" yaml:"bootstrapBucketName"`
	// Create a Customer Master Key (CMK) for the bootstrap bucket (you will be charged but can customize permissions, modern bootstrapping only).
	// Default: undefined.
	//
	// Experimental.
	BootstrapCustomerKey *string `field:"optional" json:"bootstrapCustomerKey" yaml:"bootstrapCustomerKey"`
	// AWS KMS master key ID used for the SSE-KMS encryption.
	// Default: undefined.
	//
	// Experimental.
	BootstrapKmsKeyId *string `field:"optional" json:"bootstrapKmsKeyId" yaml:"bootstrapKmsKeyId"`
	// The Managed Policy ARNs that should be attached to the role performing deployments into this environment (may be repeated, modern bootstrapping only).
	// Default: - none.
	//
	// Experimental.
	CfnExecutionPolicy *string `field:"optional" json:"cfnExecutionPolicy" yaml:"cfnExecutionPolicy"`
	// Use the permissions boundary specified by name.
	// Default: undefined.
	//
	// Experimental.
	CustomPermissionsBoundary *string `field:"optional" json:"customPermissionsBoundary" yaml:"customPermissionsBoundary"`
	// The target AWS environments to deploy the bootstrap stack to.
	//
	// Uses the following format: `aws://<account-id>/<region>`.
	//
	// Example:
	//   "aws://123456789012/us-east-1"
	//
	// Default: - Bootstrap all environments referenced in the CDK app or determine an environment from local configuration.
	//
	// Experimental.
	Environments *[]*string `field:"optional" json:"environments" yaml:"environments"`
	// Use the example permissions boundary.
	// Default: undefined.
	//
	// Experimental.
	ExamplePermissionsBoundary *bool `field:"optional" json:"examplePermissionsBoundary" yaml:"examplePermissionsBoundary"`
	// Whether to execute ChangeSet (--no-execute will NOT execute the ChangeSet).
	// Default: true.
	//
	// Experimental.
	Execute *bool `field:"optional" json:"execute" yaml:"execute"`
	// Always bootstrap even if it would downgrade template version.
	// Default: false.
	//
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Block public access configuration on CDK toolkit bucket (enabled by default).
	// Default: undefined.
	//
	// Experimental.
	PublicAccessBlockConfiguration *string `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// String which must be unique for each bootstrap stack.
	//
	// You
	// must configure it on your CDK app if you change this
	// from the default.
	// Default: undefined.
	//
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Instead of actual bootstrapping, print the current CLI\'s bootstrapping template to stdout for customization.
	// Default: false.
	//
	// Experimental.
	ShowTemplate *bool `field:"optional" json:"showTemplate" yaml:"showTemplate"`
	// Use the template from the given file instead of the built-in one (use --show-template to obtain an example).
	// Experimental.
	Template *string `field:"optional" json:"template" yaml:"template"`
	// Toggle CloudFormation termination protection on the bootstrap stacks.
	// Default: false.
	//
	// Experimental.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// The name of the CDK toolkit stack to create.
	// Experimental.
	ToolkitStackName *string `field:"optional" json:"toolkitStackName" yaml:"toolkitStackName"`
	// The AWS account IDs that should be trusted to perform deployments into this environment (may be repeated, modern bootstrapping only).
	// Default: undefined.
	//
	// Experimental.
	Trust *string `field:"optional" json:"trust" yaml:"trust"`
	// The AWS account IDs that should be trusted to look up values in this environment (may be repeated, modern bootstrapping only).
	// Default: undefined.
	//
	// Experimental.
	TrustForLookup *string `field:"optional" json:"trustForLookup" yaml:"trustForLookup"`
	// Use previous values for existing parameters (you must specify all parameters on every deployment if this is disabled).
	// Default: true.
	//
	// Experimental.
	UsePreviousParameters *bool `field:"optional" json:"usePreviousParameters" yaml:"usePreviousParameters"`
}

