package previewawsgameliftmixins


// Configuration options for Amazon GameLift Servers-managed capacity behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedCapacityConfigurationProperty := &ManagedCapacityConfigurationProperty{
//   	ScaleInAfterInactivityMinutes: jsii.Number(123),
//   	ZeroCapacityStrategy: jsii.String("zeroCapacityStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-managedcapacityconfiguration.html
//
type CfnFleetPropsMixin_ManagedCapacityConfigurationProperty struct {
	// Length of time, in minutes, that Amazon GameLift Servers will wait before scaling in your MinSize and DesiredInstances to 0 after a period with no game session activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-managedcapacityconfiguration.html#cfn-gamelift-fleet-managedcapacityconfiguration-scaleinafterinactivityminutes
	//
	ScaleInAfterInactivityMinutes *float64 `field:"optional" json:"scaleInAfterInactivityMinutes" yaml:"scaleInAfterInactivityMinutes"`
	// The strategy Amazon GameLift Servers will use to automatically scale your capacity to and from zero in response to game session activity.
	//
	// Game session activity refers to any active running sessions or game session requests. When set to SCALE_TO_AND_FROM_ZERO, MinSize must not be specified and will be managed automatically. When set to MANUAL, MinSize is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-managedcapacityconfiguration.html#cfn-gamelift-fleet-managedcapacityconfiguration-zerocapacitystrategy
	//
	ZeroCapacityStrategy *string `field:"optional" json:"zeroCapacityStrategy" yaml:"zeroCapacityStrategy"`
}

