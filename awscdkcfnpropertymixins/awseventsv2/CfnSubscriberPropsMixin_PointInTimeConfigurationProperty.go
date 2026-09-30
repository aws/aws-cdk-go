package awseventsv2


// The point in time to start delivering events from, used when StartingPosition is POINT_IN_TIME.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   pointInTimeConfigurationProperty := &PointInTimeConfigurationProperty{
//   	EndPoint: jsii.Number(123),
//   	PointType: jsii.String("pointType"),
//   	StartingPoint: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-pointintimeconfiguration.html
//
type CfnSubscriberPropsMixin_PointInTimeConfigurationProperty struct {
	// An optional time to stop delivering events at, in seconds since the Unix epoch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-pointintimeconfiguration.html#cfn-eventsv2-subscriber-pointintimeconfiguration-endpoint
	//
	EndPoint *float64 `field:"optional" json:"endPoint" yaml:"endPoint"`
	// Where to start: HORIZON starts from the earliest available event;
	//
	// TIMESTAMP starts from the StartingPoint timestamp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-pointintimeconfiguration.html#cfn-eventsv2-subscriber-pointintimeconfiguration-pointtype
	//
	PointType *string `field:"optional" json:"pointType" yaml:"pointType"`
	// The time to start delivering events from, in seconds since the Unix epoch.
	//
	// Required when PointType is TIMESTAMP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-pointintimeconfiguration.html#cfn-eventsv2-subscriber-pointintimeconfiguration-startingpoint
	//
	StartingPoint *float64 `field:"optional" json:"startingPoint" yaml:"startingPoint"`
}

