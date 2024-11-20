package awsgamelift


// A remote location where a multi-location fleet can deploy game servers for game hosting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationConfigurationProperty := &LocationConfigurationProperty{
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	LocationCapacity: &LocationCapacityProperty{
//   		DesiredEc2Instances: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationconfiguration.html
//
type CfnFleet_LocationConfigurationProperty struct {
	// An AWS Region code, such as `us-west-2` .
	//
	// For a list of supported Regions and Local Zones, see [Amazon GameLift service locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationconfiguration.html#cfn-gamelift-fleet-locationconfiguration-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// Current resource capacity settings for managed EC2 fleets and managed container fleets.
	//
	// For multi-location fleets, location values might refer to a fleet's remote location or its home Region.
	//
	// *Returned by:* [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) , [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) , [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-locationconfiguration.html#cfn-gamelift-fleet-locationconfiguration-locationcapacity
	//
	LocationCapacity interface{} `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
}

