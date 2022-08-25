package awsstepfunctionstasks


// An automatic scaling configuration, which describes how the policy adds or removes instances, the cooldown period, and the number of EC2 instances that will be added each time the CloudWatch metric alarm condition is satisfied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleScalingPolicyConfigurationProperty := &simpleScalingPolicyConfigurationProperty{
//   	scalingAdjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	adjustmentType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   	coolDown: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SimpleScalingPolicyConfiguration.html
//
type EmrCreateCluster_SimpleScalingPolicyConfigurationProperty struct {
	// The amount by which to scale in or scale out, based on the specified AdjustmentType.
	//
	// A positive value adds to the instance group's
	// EC2 instance count while a negative number removes instances. If AdjustmentType is set to EXACT_CAPACITY, the number should only be
	// a positive integer.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The way in which EC2 instances are added (if ScalingAdjustment is a positive number) or terminated (if ScalingAdjustment is a negative number) each time the scaling activity is triggered.
	AdjustmentType EmrCreateCluster_ScalingAdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
	CoolDown *float64 `field:"optional" json:"coolDown" yaml:"coolDown"`
}

