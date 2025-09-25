package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMultiplex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiplexProps := &CfnMultiplexProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	MultiplexSettings: &MultiplexSettingsProperty{
//   		TransportStreamBitrate: jsii.Number(123),
//   		TransportStreamId: jsii.Number(123),
//
//   		// the properties below are optional
//   		MaximumVideoBufferDelayMilliseconds: jsii.Number(123),
//   		TransportStreamReservedBitrate: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Destinations: []interface{}{
//   		&MultiplexOutputDestinationProperty{
//   			MultiplexMediaConnectOutputDestinationSettings: &MultiplexMediaConnectOutputDestinationSettingsProperty{
//   				EntitlementArn: jsii.String("entitlementArn"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html
//
type CfnMultiplexProps struct {
	// A list of availability zones for the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-availabilityzones
	//
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Configuration for a multiplex event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-multiplexsettings
	//
	MultiplexSettings interface{} `field:"required" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The name of the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of the multiplex output destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// A collection of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

