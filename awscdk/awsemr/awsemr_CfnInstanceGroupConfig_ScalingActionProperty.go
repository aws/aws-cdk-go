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
//   scalingActionProperty := &scalingActionProperty{
//   	simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   		scalingAdjustment: jsii.Number(123),
//
//   		// the properties below are optional
//   		adjustmentType: jsii.String("adjustmentType"),
//   		coolDown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	market: jsii.String("market"),
//   }
//
type CfnInstanceGroupConfig_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	SimpleScalingPolicyConfiguration interface{} `field:"required" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	Market *string `field:"optional" json:"market" yaml:"market"`
}

