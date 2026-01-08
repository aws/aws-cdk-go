package previewawsgameliftstreamsmixins


// Configuration settings that define a stream group's stream capacity for a location.
//
// When configuring a location for the first time, you must specify a numeric value for at least one of the two capacity types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationConfigurationProperty := &LocationConfigurationProperty{
//   	AlwaysOnCapacity: jsii.Number(123),
//   	LocationName: jsii.String("locationName"),
//   	MaximumCapacity: jsii.Number(123),
//   	OnDemandCapacity: jsii.Number(123),
//   	TargetIdleCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html
//
type CfnStreamGroupPropsMixin_LocationConfigurationProperty struct {
	// This setting, if non-zero, indicates minimum streaming capacity which is allocated to you and is never released back to the service.
	//
	// You pay for this base level of capacity at all times, whether used or idle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-alwaysoncapacity
	//
	AlwaysOnCapacity *float64 `field:"optional" json:"alwaysOnCapacity" yaml:"alwaysOnCapacity"`
	// A location's name.
	//
	// For example, `us-east-1` . For a complete list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html) in the *Amazon GameLift Streams Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-locationname
	//
	LocationName *string `field:"optional" json:"locationName" yaml:"locationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-maximumcapacity
	//
	MaximumCapacity *float64 `field:"optional" json:"maximumCapacity" yaml:"maximumCapacity"`
	// This field is deprecated. Use MaximumCapacity instead. This parameter is ignored when MaximumCapacity is specified.
	//
	// The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated. This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes). Default is 0 when you create a stream group or add a location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-ondemandcapacity
	//
	OnDemandCapacity *float64 `field:"optional" json:"onDemandCapacity" yaml:"onDemandCapacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html#cfn-gameliftstreams-streamgroup-locationconfiguration-targetidlecapacity
	//
	TargetIdleCapacity *float64 `field:"optional" json:"targetIdleCapacity" yaml:"targetIdleCapacity"`
}

