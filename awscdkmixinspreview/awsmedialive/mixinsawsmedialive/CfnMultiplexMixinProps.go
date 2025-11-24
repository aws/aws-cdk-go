package mixinsawsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMultiplexPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMultiplexMixinProps := &CfnMultiplexMixinProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	Destinations: []interface{}{
//   		&MultiplexOutputDestinationProperty{
//   			MultiplexMediaConnectOutputDestinationSettings: &MultiplexMediaConnectOutputDestinationSettingsProperty{
//   				EntitlementArn: jsii.String("entitlementArn"),
//   			},
//   		},
//   	},
//   	MultiplexSettings: &MultiplexSettingsProperty{
//   		MaximumVideoBufferDelayMilliseconds: jsii.Number(123),
//   		TransportStreamBitrate: jsii.Number(123),
//   		TransportStreamId: jsii.Number(123),
//   		TransportStreamReservedBitrate: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html
//
type CfnMultiplexMixinProps struct {
	// A list of availability zones for the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// A list of the multiplex output destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Configuration for a multiplex event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-multiplexsettings
	//
	MultiplexSettings interface{} `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The name of the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A collection of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplex.html#cfn-medialive-multiplex-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

