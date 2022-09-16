package awsstepfunctionstasks


// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// And an automatic scaling configuration, which describes how the policy adds or removes instances, the cooldown period,
// and the number of EC2 instances that will be added each time the CloudWatch metric alarm condition is satisfied.
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
//   		adjustmentType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   		coolDown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.instanceMarket_ON_DEMAND,
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingAction.html
//
type EmrCreateCluster_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	SimpleScalingPolicyConfiguration *EmrCreateCluster_SimpleScalingPolicyConfigurationProperty `field:"required" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	Market EmrCreateCluster_InstanceMarket `field:"optional" json:"market" yaml:"market"`
}

