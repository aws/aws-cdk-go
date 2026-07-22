package cloudassemblyschema


// A violation that was acknowledged/suppressed and excluded from the active violation set.
//
// Example:
//   import "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//
//   suppressed := &SuppressedViolationJson{
//   	RuleName: jsii.String("no-public-access"),
//   	Description: jsii.String("S3 bucket should not allow public access"),
//   	Severity: jsii.String("warning"),
//   	ViolatingConstructs: []ViolatingConstructJson{
//   		&ViolatingConstructJson{
//   			ConstructPath: jsii.String("MyStack/MyBucket"),
//   		},
//   	},
//   	AcknowledgedId: jsii.String("my-plugin::no-public-access"),
//   }
//
type SuppressedViolationJson struct {
	// A description of the violation.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the rule that was violated.
	//
	// This may include a namespace: `'<namespace>::<rule-name>'`. If it does not,
	// the plugin name will be used as the namespace.
	//
	// The namespace must be included when acknowledging the violation.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The severity of the violation.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Constructs that violated the rule.
	ViolatingConstructs *[]*ViolatingConstructJson `field:"required" json:"violatingConstructs" yaml:"violatingConstructs"`
	// If the plugin wants to report using a non-standard severity, put it here.
	CustomSeverity *string `field:"optional" json:"customSeverity" yaml:"customSeverity"`
	// Additional rule-specific metadata.
	// Default: - no metadata.
	//
	RuleMetadata *map[string]*string `field:"optional" json:"ruleMetadata" yaml:"ruleMetadata"`
	// How to fix the violation.
	// Default: - no fix provided.
	//
	SuggestedFix *string `field:"optional" json:"suggestedFix" yaml:"suggestedFix"`
	// The acknowledgement ID that caused this violation to be suppressed.
	//
	// Format: `<plugin-name>::<rule-name>` (spaces replaced with hyphens).
	AcknowledgedId *string `field:"required" json:"acknowledgedId" yaml:"acknowledgedId"`
	// The construct path where the acknowledgement was declared.
	// Default: - unknown.
	//
	AcknowledgedAt *string `field:"optional" json:"acknowledgedAt" yaml:"acknowledgedAt"`
	// Stack trace showing where the acknowledgement was declared.
	//
	// A `\n`-delimited string of stack frames.
	// Default: - no stack trace.
	//
	AcknowledgedStackTrace *string `field:"optional" json:"acknowledgedStackTrace" yaml:"acknowledgedStackTrace"`
	// The reason given for the acknowledgement, if provided.
	// Default: - no reason given.
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

