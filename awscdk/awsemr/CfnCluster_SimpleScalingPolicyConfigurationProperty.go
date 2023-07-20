package awsemr


// `SimpleScalingPolicyConfiguration` is a subproperty of the `ScalingAction` property type.
//
// `SimpleScalingPolicyConfiguration` determines how an automatic scaling action adds or removes instances, the cooldown period, and the number of EC2 instances that are added each time the CloudWatch metric alarm condition is satisfied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleScalingPolicyConfigurationProperty := &SimpleScalingPolicyConfigurationProperty{
//   	ScalingAdjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	AdjustmentType: jsii.String("adjustmentType"),
//   	CoolDown: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-simplescalingpolicyconfiguration.html
//
type CfnCluster_SimpleScalingPolicyConfigurationProperty struct {
	// The amount by which to scale in or scale out, based on the specified `AdjustmentType` .
	//
	// A positive value adds to the instance group's Amazon EC2 instance count while a negative number removes instances. If `AdjustmentType` is set to `EXACT_CAPACITY` , the number should only be a positive integer. If `AdjustmentType` is set to `PERCENT_CHANGE_IN_CAPACITY` , the value should express the percentage as an integer. For example, -20 indicates a decrease in 20% increments of cluster capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-simplescalingpolicyconfiguration.html#cfn-emr-cluster-simplescalingpolicyconfiguration-scalingadjustment
	//
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The way in which Amazon EC2 instances are added (if `ScalingAdjustment` is a positive number) or terminated (if `ScalingAdjustment` is a negative number) each time the scaling activity is triggered.
	//
	// `CHANGE_IN_CAPACITY` is the default. `CHANGE_IN_CAPACITY` indicates that the Amazon EC2 instance count increments or decrements by `ScalingAdjustment` , which should be expressed as an integer. `PERCENT_CHANGE_IN_CAPACITY` indicates the instance count increments or decrements by the percentage specified by `ScalingAdjustment` , which should be expressed as an integer. For example, 20 indicates an increase in 20% increments of cluster capacity. `EXACT_CAPACITY` indicates the scaling activity results in an instance group with the number of Amazon EC2 instances specified by `ScalingAdjustment` , which should be expressed as a positive integer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-simplescalingpolicyconfiguration.html#cfn-emr-cluster-simplescalingpolicyconfiguration-adjustmenttype
	//
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
	//
	// The default value is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-simplescalingpolicyconfiguration.html#cfn-emr-cluster-simplescalingpolicyconfiguration-cooldown
	//
	CoolDown *float64 `field:"optional" json:"coolDown" yaml:"coolDown"`
}

