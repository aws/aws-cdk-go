// AWS CDK Programmatic CLI library
package awscdkclilibalpha


// Options for cdk list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cli_lib_alpha "github.com/aws/aws-cdk-go/awscdkclilibalpha"
//
//   listOptions := &ListOptions{
//   	AssetMetadata: jsii.Boolean(false),
//   	CaBundlePath: jsii.String("caBundlePath"),
//   	Color: jsii.Boolean(false),
//   	Context: map[string]*string{
//   		"contextKey": jsii.String("context"),
//   	},
//   	Debug: jsii.Boolean(false),
//   	Ec2Creds: jsii.Boolean(false),
//   	IgnoreErrors: jsii.Boolean(false),
//   	Json: jsii.Boolean(false),
//   	Long: jsii.Boolean(false),
//   	Lookups: jsii.Boolean(false),
//   	Notices: jsii.Boolean(false),
//   	PathMetadata: jsii.Boolean(false),
//   	Profile: jsii.String("profile"),
//   	Proxy: jsii.String("proxy"),
//   	RoleArn: jsii.String("roleArn"),
//   	Stacks: []*string{
//   		jsii.String("stacks"),
//   	},
//   	Staging: jsii.Boolean(false),
//   	Strict: jsii.Boolean(false),
//   	Trace: jsii.Boolean(false),
//   	Verbose: jsii.Boolean(false),
//   	VersionReporting: jsii.Boolean(false),
//   }
//
// Experimental.
type ListOptions struct {
	// Include "aws:asset:*" CloudFormation metadata for resources that use assets.
	// Experimental.
	AssetMetadata *bool `field:"optional" json:"assetMetadata" yaml:"assetMetadata"`
	// Path to CA certificate to use when validating HTTPS requests.
	// Experimental.
	CaBundlePath *string `field:"optional" json:"caBundlePath" yaml:"caBundlePath"`
	// Show colors and other style from console output.
	// Experimental.
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Additional context.
	// Experimental.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// enable emission of additional debugging information, such as creation stack traces of tokens.
	// Experimental.
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
	// Force trying to fetch EC2 instance credentials.
	// Experimental.
	Ec2Creds *bool `field:"optional" json:"ec2Creds" yaml:"ec2Creds"`
	// Ignores synthesis errors, which will likely produce an invalid output.
	// Experimental.
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Use JSON output instead of YAML when templates are printed to STDOUT.
	// Experimental.
	Json *bool `field:"optional" json:"json" yaml:"json"`
	// Perform context lookups.
	//
	// Synthesis fails if this is disabled and context lookups need
	// to be performed.
	// Experimental.
	Lookups *bool `field:"optional" json:"lookups" yaml:"lookups"`
	// Show relevant notices.
	// Experimental.
	Notices *bool `field:"optional" json:"notices" yaml:"notices"`
	// Include "aws:cdk:path" CloudFormation metadata for each resource.
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// Use the indicated AWS profile as the default environment.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Use the indicated proxy.
	//
	// Will read from
	// HTTPS_PROXY environment if specified.
	// Experimental.
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Role to pass to CloudFormation for deployment.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of stacks to deploy.
	// Experimental.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
	// Copy assets to the output directory.
	//
	// Needed for local debugging the source files with SAM CLI.
	// Experimental.
	Staging *bool `field:"optional" json:"staging" yaml:"staging"`
	// Do not construct stacks with warnings.
	// Experimental.
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// Print trace for stack warnings.
	// Experimental.
	Trace *bool `field:"optional" json:"trace" yaml:"trace"`
	// show debug logs.
	// Experimental.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Include "AWS::CDK::Metadata" resource in synthesized templates.
	// Experimental.
	VersionReporting *bool `field:"optional" json:"versionReporting" yaml:"versionReporting"`
	// Display environment information for each stack.
	// Experimental.
	Long *bool `field:"optional" json:"long" yaml:"long"`
}

