package mixinsawscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEnabledBaselinePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var value interface{}
//
//   cfnEnabledBaselineMixinProps := &CfnEnabledBaselineMixinProps{
//   	BaselineIdentifier: jsii.String("baselineIdentifier"),
//   	BaselineVersion: jsii.String("baselineVersion"),
//   	Parameters: []interface{}{
//   		&ParameterProperty{
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
//   	TargetIdentifier: jsii.String("targetIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html
//
type CfnEnabledBaselineMixinProps struct {
	// The specific `Baseline` enabled as part of the `EnabledBaseline` resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-baselineidentifier
	//
	BaselineIdentifier *string `field:"optional" json:"baselineIdentifier" yaml:"baselineIdentifier"`
	// The enabled version of the `Baseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-baselineversion
	//
	BaselineVersion *string `field:"optional" json:"baselineVersion" yaml:"baselineVersion"`
	// Shows the parameters that are applied when enabling this `Baseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The target on which to enable the `Baseline` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controltower-enabledbaseline.html#cfn-controltower-enabledbaseline-targetidentifier
	//
	TargetIdentifier *string `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

