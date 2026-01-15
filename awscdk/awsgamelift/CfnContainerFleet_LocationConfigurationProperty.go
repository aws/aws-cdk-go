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
//   		MaxSize: jsii.Number(123),
//
//   		// the properties below are optional
//   		DesiredEc2Instances: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   	StoppedActions: []*string{
//   		jsii.String("stoppedActions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html
//
type CfnContainerFleet_LocationConfigurationProperty struct {
	// An AWS Region code, such as `us-west-2` .
	//
	// For a list of supported Regions and Local Zones, see [Amazon GameLift Servers service locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html#cfn-gamelift-containerfleet-locationconfiguration-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html#cfn-gamelift-containerfleet-locationconfiguration-locationcapacity
	//
	LocationCapacity interface{} `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
	// A list of fleet actions that have been suspended in the fleet location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html#cfn-gamelift-containerfleet-locationconfiguration-stoppedactions
	//
	StoppedActions *[]*string `field:"optional" json:"stoppedActions" yaml:"stoppedActions"`
}

