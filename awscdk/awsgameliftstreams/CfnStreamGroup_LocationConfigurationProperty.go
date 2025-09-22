package awsgameliftstreams


// Configuration settings that define a stream group's stream capacity for a location.
//
// When configuring a location for the first time, you must specify a numeric value for at least one of the two capacity types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationConfigurationProperty := &LocationConfigurationProperty{
//   	LocationName: jsii.String("locationName"),
//
//   	// the properties below are optional
//   	AlwaysOnCapacity: jsii.Number(123),
//   	OnDemandCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html
//
type CfnStreamGroup_LocationConfigurationProperty struct {
	// A location's name.
	//
	// For example, `us-east-1` . For a complete list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html) in the *Amazon GameLift Streams Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-locationname
	//
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// The streaming capacity that is allocated and ready to handle stream requests without delay.
	//
	// You pay for this capacity whether it's in use or not. Best for quickest time from streaming request to streaming session. Default is 1 (2 for high stream classes) when creating a stream group or adding a location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-alwaysoncapacity
	//
	AlwaysOnCapacity *float64 `field:"optional" json:"alwaysOnCapacity" yaml:"alwaysOnCapacity"`
	// The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated.
	//
	// This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes). Default is 0 when creating a stream group or adding a location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-ondemandcapacity
	//
	OnDemandCapacity *float64 `field:"optional" json:"onDemandCapacity" yaml:"onDemandCapacity"`
}

