package mixinsawsemr


// `ScalingAction` is a subproperty of the `ScalingRule` property type.
//
// `ScalingAction` determines the type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingActionProperty := &ScalingActionProperty{
//   	Market: jsii.String("market"),
//   	SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   		AdjustmentType: jsii.String("adjustmentType"),
//   		CoolDown: jsii.Number(123),
//   		ScalingAdjustment: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingaction.html
//
type CfnClusterPropsMixin_ScalingActionProperty struct {
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingaction.html#cfn-emr-cluster-scalingaction-market
	//
	Market *string `field:"optional" json:"market" yaml:"market"`
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scalingaction.html#cfn-emr-cluster-scalingaction-simplescalingpolicyconfiguration
	//
	SimpleScalingPolicyConfiguration interface{} `field:"optional" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
}

