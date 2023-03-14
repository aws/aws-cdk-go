package awscontroltower


// Properties for defining a `CfnEnabledControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnabledControlProps := &cfnEnabledControlProps{
//   	controlIdentifier: jsii.String("controlIdentifier"),
//   	targetIdentifier: jsii.String("targetIdentifier"),
//   }
//
type CfnEnabledControlProps struct {
	// The ARN of the control.
	//
	// Only *Strongly recommended* and *Elective* controls are permitted, with the exception of the *Region deny* guardrail.
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// The ARN of the organizational unit.
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

