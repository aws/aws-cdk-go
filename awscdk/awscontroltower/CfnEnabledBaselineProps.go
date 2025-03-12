package awscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnabledBaseline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnEnabledBaselineProps := &CfnEnabledBaselineProps{
//   	BaselineIdentifier: jsii.String("baselineIdentifier"),
//   	BaselineVersion: jsii.String("baselineVersion"),
//   	TargetIdentifier: jsii.String("targetIdentifier"),
//
//   	// the properties below are optional
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: value,
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html
//
type CfnEnabledBaselineProps struct {
	// The specific `Baseline` enabled as part of the `EnabledBaseline` resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-baselineidentifier
	//
	BaselineIdentifier *string `field:"required" json:"baselineIdentifier" yaml:"baselineIdentifier"`
	// The enabled version of the `Baseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-baselineversion
	//
	BaselineVersion *string `field:"required" json:"baselineVersion" yaml:"baselineVersion"`
	// The target on which to enable the `Baseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-targetidentifier
	//
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
	// Parameters that are applied when enabling this `Baseline` .
	//
	// These parameters configure the behavior of the baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Tags associated with input to `EnableBaseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

