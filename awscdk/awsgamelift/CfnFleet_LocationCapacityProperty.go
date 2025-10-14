package awsgamelift


// Current resource capacity settings for managed EC2 fleets and managed container fleets.
//
// For multi-location fleets, location values might refer to a fleet's remote location or its home Region.
//
// *Returned by:* [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) , [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) , [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationCapacityProperty := &LocationCapacityProperty{
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//
//   	// the properties below are optional
//   	DesiredEc2Instances: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationcapacity.html
//
type CfnFleet_LocationCapacityProperty struct {
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationcapacity.html#cfn-gamelift-fleet-locationcapacity-maxsize
	//
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationcapacity.html#cfn-gamelift-fleet-locationcapacity-minsize
	//
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	// The number of Amazon EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits. Changes in desired instance value can take up to 1 minute to be reflected when viewing the fleet's capacity settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationcapacity.html#cfn-gamelift-fleet-locationcapacity-desiredec2instances
	//
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
}

