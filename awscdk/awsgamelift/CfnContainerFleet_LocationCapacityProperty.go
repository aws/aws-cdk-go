package awsgamelift


// Current resource capacity settings in a specified fleet or location.
//
// The location value might refer to a fleet's remote location or its home Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationCapacityProperty := &LocationCapacityProperty{
//   	DesiredEc2Instances: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationcapacity.html
//
type CfnContainerFleet_LocationCapacityProperty struct {
	// The number of EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationcapacity.html#cfn-gamelift-containerfleet-locationcapacity-desiredec2instances
	//
	DesiredEc2Instances *float64 `field:"required" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The maximum value that is allowed for the fleet's instance count for a location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationcapacity.html#cfn-gamelift-containerfleet-locationcapacity-maxsize
	//
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum value allowed for the fleet's instance count for a location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationcapacity.html#cfn-gamelift-containerfleet-locationcapacity-minsize
	//
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
}

