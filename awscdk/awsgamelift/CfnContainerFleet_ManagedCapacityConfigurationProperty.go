package awsgamelift


// Configuration options for GameLift-managed capacity behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedCapacityConfigurationProperty := &ManagedCapacityConfigurationProperty{
//   	ZeroCapacityStrategy: jsii.String("zeroCapacityStrategy"),
//
//   	// the properties below are optional
//   	ScaleInAfterInactivityMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-managedcapacityconfiguration.html
//
type CfnContainerFleet_ManagedCapacityConfigurationProperty struct {
	// The strategy Amazon GameLift Servers will use to automatically scale your capacity to and from zero in response to game session activity.
	//
	// Game session activity refers to any active running sessions or game session requests. When set to SCALE_TO_AND_FROM_ZERO, MinSize must not be specified and will be managed automatically. When set to MANUAL, MinSize is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-managedcapacityconfiguration.html#cfn-gamelift-containerfleet-managedcapacityconfiguration-zerocapacitystrategy
	//
	ZeroCapacityStrategy *string `field:"required" json:"zeroCapacityStrategy" yaml:"zeroCapacityStrategy"`
	// Length of time, in minutes, that Amazon GameLift Servers will wait before scaling in your MinSize and DesiredInstances to 0 after a period with no game session activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-managedcapacityconfiguration.html#cfn-gamelift-containerfleet-managedcapacityconfiguration-scaleinafterinactivityminutes
	//
	ScaleInAfterInactivityMinutes *float64 `field:"optional" json:"scaleInAfterInactivityMinutes" yaml:"scaleInAfterInactivityMinutes"`
}

