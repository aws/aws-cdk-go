package awsautoscaling


// `TotalLocalStorageGBRequest` is a property of the `InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property type that describes the minimum and maximum total local storage size for an instance type, in GB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   totalLocalStorageGBRequestProperty := &totalLocalStorageGBRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnAutoScalingGroup_TotalLocalStorageGBRequestProperty struct {
	// The storage maximum in GB.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The storage minimum in GB.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

