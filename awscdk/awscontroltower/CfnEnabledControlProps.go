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
	// Only *Strongly recommended* and *Elective* controls are permitted, with the exception of the *Region deny* control. For information on how to find the `controlIdentifier` , see [the overview page](https://docs.aws.amazon.com//controltower/latest/APIReference/Welcome.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-controlidentifier
	//
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// The ARN of the organizational unit.
	//
	// For information on how to find the `targetIdentifier` , see [the overview page](https://docs.aws.amazon.com//controltower/latest/APIReference/Welcome.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-targetidentifier
	//
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

