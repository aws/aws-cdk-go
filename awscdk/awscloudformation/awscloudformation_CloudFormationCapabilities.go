package awscloudformation


// Capabilities that affect whether CloudFormation is allowed to change IAM resources.
// Deprecated: use `core.CfnCapabilities`
type CloudFormationCapabilities string

const (
	// No IAM Capabilities.
	//
	// Pass this capability if you wish to block the creation IAM resources.
	// Deprecated: use `core.CfnCapabilities`
	CloudFormationCapabilities_NONE CloudFormationCapabilities = "NONE"
	// Capability to create anonymous IAM resources.
	//
	// Pass this capability if you're only creating anonymous resources.
	// Deprecated: use `core.CfnCapabilities`
	CloudFormationCapabilities_ANONYMOUS_IAM CloudFormationCapabilities = "ANONYMOUS_IAM"
	// Capability to create named IAM resources.
	//
	// Pass this capability if you're creating IAM resources that have physical
	// names.
	//
	// `CloudFormationCapabilities.NamedIAM` implies `CloudFormationCapabilities.IAM`; you don't have to pass both.
	// Deprecated: use `core.CfnCapabilities`
	CloudFormationCapabilities_NAMED_IAM CloudFormationCapabilities = "NAMED_IAM"
	// Capability to run CloudFormation macros.
	//
	// Pass this capability if your template includes macros, for example AWS::Include or AWS::Serverless.
	// Deprecated: use `core.CfnCapabilities`
	CloudFormationCapabilities_AUTO_EXPAND CloudFormationCapabilities = "AUTO_EXPAND"
)

