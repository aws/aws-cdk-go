package mixinsawsgamelift


// A remote location where a multi-location fleet can deploy game servers for game hosting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationConfigurationProperty := &LocationConfigurationProperty{
//   	Location: jsii.String("location"),
//   	LocationCapacity: &LocationCapacityProperty{
//   		DesiredEc2Instances: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		MinSize: jsii.Number(123),
//   	},
//   	StoppedActions: []*string{
//   		jsii.String("stoppedActions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html
//
type CfnContainerFleetPropsMixin_LocationConfigurationProperty struct {
	// An AWS Region code, such as `us-west-2` .
	//
	// For a list of supported Regions and Local Zones, see [Amazon GameLift Servers service locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-locationconfiguration.html#cfn-gamelift-containerfleet-locationconfiguration-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
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

