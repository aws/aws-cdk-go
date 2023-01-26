package cloudassemblyschema


// Options to use with cdk destroy.
//
// Example:
//   app := awscdk.NewApp()
//
//   stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))
//
//   stack := awscdk.NewStack(app, jsii.String("stack"))
//
//   testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &integTestProps{
//   	testCases: []stack{
//   		stackUnderTest,
//   	},
//   	diffAssets: jsii.Boolean(true),
//   	stackUpdateWorkflow: jsii.Boolean(true),
//   	cdkCommandOptions: &cdkCommands{
//   		deploy: &deployCommand{
//   			args: &deployOptions{
//   				requireApproval: awscdk.RequireApproval_NEVER,
//   				json: jsii.Boolean(true),
//   			},
//   		},
//   		destroy: &destroyCommand{
//   			args: &destroyOptions{
//   				force: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   })
//
type DestroyOptions struct {
	// Deploy all stacks.
	//
	// Requried if `stacks` is not set.
	All *bool `field:"optional" json:"all" yaml:"all"`
	// command-line for executing your app or a cloud assembly directory e.g. "node bin/my-app.js" or "cdk.out".
	App *string `field:"optional" json:"app" yaml:"app"`
	// Include "aws:asset:*" CloudFormation metadata for resources that use assets.
	AssetMetadata *bool `field:"optional" json:"assetMetadata" yaml:"assetMetadata"`
	// Path to CA certificate to use when validating HTTPS requests.
	CaBundlePath *string `field:"optional" json:"caBundlePath" yaml:"caBundlePath"`
	// Show colors and other style from console output.
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Additional context.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// enable emission of additional debugging information, such as creation stack traces of tokens.
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
	// Force trying to fetch EC2 instance credentials.
	Ec2Creds *bool `field:"optional" json:"ec2Creds" yaml:"ec2Creds"`
	// Ignores synthesis errors, which will likely produce an invalid output.
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Use JSON output instead of YAML when templates are printed to STDOUT.
	Json *bool `field:"optional" json:"json" yaml:"json"`
	// Perform context lookups.
	//
	// Synthesis fails if this is disabled and context lookups need
	// to be performed.
	Lookups *bool `field:"optional" json:"lookups" yaml:"lookups"`
	// Show relevant notices.
	Notices *bool `field:"optional" json:"notices" yaml:"notices"`
	// Emits the synthesized cloud assembly into a directory.
	Output *string `field:"optional" json:"output" yaml:"output"`
	// Include "aws:cdk:path" CloudFormation metadata for each resource.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// Use the indicated AWS profile as the default environment.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Use the indicated proxy.
	//
	// Will read from
	// HTTPS_PROXY environment if specified.
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Role to pass to CloudFormation for deployment.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of stacks to deploy.
	//
	// Requried if `all` is not set.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
	// Copy assets to the output directory.
	//
	// Needed for local debugging the source files with SAM CLI.
	Staging *bool `field:"optional" json:"staging" yaml:"staging"`
	// Do not construct stacks with warnings.
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// Print trace for stack warnings.
	Trace *bool `field:"optional" json:"trace" yaml:"trace"`
	// show debug logs.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Include "AWS::CDK::Metadata" resource in synthesized templates.
	VersionReporting *bool `field:"optional" json:"versionReporting" yaml:"versionReporting"`
	// Only destroy the given stack.
	Exclusively *bool `field:"optional" json:"exclusively" yaml:"exclusively"`
	// Do not ask for permission before destroying stacks.
	Force *bool `field:"optional" json:"force" yaml:"force"`
}

