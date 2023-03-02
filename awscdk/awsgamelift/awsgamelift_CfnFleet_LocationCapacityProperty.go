package awsgamelift


// Current resource capacity settings in a specified fleet or location.
//
// The location value might refer to a fleet's remote location or its home Region.
//
// *Related actions*
//
// [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) | [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) | [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationCapacityProperty := &locationCapacityProperty{
//   	desiredEc2Instances: jsii.Number(123),
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   }
//
type CfnFleet_LocationCapacityProperty struct {
	// The number of Amazon EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	DesiredEc2Instances *float64 `field:"required" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
}

