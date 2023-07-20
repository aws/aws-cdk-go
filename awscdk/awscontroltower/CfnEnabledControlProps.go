package awscontroltower


// Properties for defining a `CfnEnabledControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnabledControlProps := &CfnEnabledControlProps{
//   	ControlIdentifier: jsii.String("controlIdentifier"),
//   	TargetIdentifier: jsii.String("targetIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html
//
type CfnEnabledControlProps struct {
	// The ARN of the control.
	//
	// Only *Strongly recommended* and *Elective* controls are permitted, with the exception of the *Region deny* guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-controlidentifier
	//
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// The ARN of the organizational unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-targetidentifier
	//
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

