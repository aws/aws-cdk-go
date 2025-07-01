package awsemr


// `ScalingAction` is a subproperty of the `ScalingRule` property type.
//
// `ScalingAction` determines the type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingActionProperty := &ScalingActionProperty{
//   	SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   		ScalingAdjustment: jsii.Number(123),
//
//   		// the properties below are optional
//   		AdjustmentType: jsii.String("adjustmentType"),
//   		CoolDown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Market: jsii.String("market"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingaction.html
//
type CfnInstanceGroupConfig_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingaction.html#cfn-emr-instancegroupconfig-scalingaction-simplescalingpolicyconfiguration
	//
	SimpleScalingPolicyConfiguration interface{} `field:"required" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-scalingaction.html#cfn-emr-instancegroupconfig-scalingaction-market
	//
	Market *string `field:"optional" json:"market" yaml:"market"`
}

