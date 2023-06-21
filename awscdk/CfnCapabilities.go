package awscdk


// Capabilities that affect whether CloudFormation is allowed to change IAM resources.
type CfnCapabilities string

const (
	// No IAM Capabilities.
	//
	// Pass this capability if you wish to block the creation IAM resources.
	CfnCapabilities_NONE CfnCapabilities = "NONE"
	// Capability to create anonymous IAM resources.
	//
	// Pass this capability if you're only creating anonymous resources.
	CfnCapabilities_ANONYMOUS_IAM CfnCapabilities = "ANONYMOUS_IAM"
	// Capability to create named IAM resources.
	//
	// Pass this capability if you're creating IAM resources that have physical
	// names.
	//
	// `CloudFormationCapabilities.NamedIAM` implies `CloudFormationCapabilities.IAM`; you don't have to pass both.
	CfnCapabilities_NAMED_IAM CfnCapabilities = "NAMED_IAM"
	// Capability to run CloudFormation macros.
	//
	// Pass this capability if your template includes macros, for example AWS::Include or AWS::Serverless.
	CfnCapabilities_AUTO_EXPAND CfnCapabilities = "AUTO_EXPAND"
)

