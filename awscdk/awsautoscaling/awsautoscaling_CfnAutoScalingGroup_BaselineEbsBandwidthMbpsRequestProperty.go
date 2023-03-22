package awsautoscaling


// `BaselineEbsBandwidthMbpsRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum baseline bandwidth performance for an instance type, in Mbps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineEbsBandwidthMbpsRequestProperty := &baselineEbsBandwidthMbpsRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnAutoScalingGroup_BaselineEbsBandwidthMbpsRequestProperty struct {
	// The maximum value in Mbps.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum value in Mbps.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

