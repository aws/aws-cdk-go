package awscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnabledControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnEnabledControlProps := &CfnEnabledControlProps{
//   	ControlIdentifier: jsii.String("controlIdentifier"),
//   	TargetIdentifier: jsii.String("targetIdentifier"),
//
//   	// the properties below are optional
//   	Parameters: []interface{}{
//   		&EnabledControlParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: value,
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
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
	// Array of `EnabledControlParameter` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A set of tags to assign to the enabled control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledcontrol.html#cfn-controltower-enabledcontrol-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

