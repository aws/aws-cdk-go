package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2"
)

// A MediaPackage V2 destination for a MediaLive output group.
//
// Use the static factory method to create.
//
// Example:
//   var primary IChannel
//   var hdVideo EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_MediaPackageV2PerPipeline(&MediaPackageV2PerPipelineOutputGroupProps{
//   	Name: jsii.String("emp"),
//   	Destinations: []MediaPackageV2Destination{
//   		medialive.MediaPackageV2Destination_Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_2()),
//   		medialive.MediaPackageV2Destination_*Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_1()),
//   	},
//   	Outputs: []MediaPackageV2OutputDefinition{
//   		&MediaPackageV2OutputDefinition{
//   			Encode: hdVideo,
//   			OutputName: jsii.String("hd"),
//   		},
//   	},
//   })
//
// Experimental.
type MediaPackageV2Destination interface {
}

// The jsii proxy struct for MediaPackageV2Destination
type jsiiProxy_MediaPackageV2Destination struct {
	_ byte // padding
}

// Experimental.
func NewMediaPackageV2Destination_Override(m MediaPackageV2Destination) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2Destination",
		nil, // no parameters
		m,
	)
}

// Create a MediaPackage V2 destination.
//
// The region is resolved from the channel's `region` property. Import the MediaPackage V2
// channel with its region (e.g. via `fromChannelAttributes`) for cross-region delivery.
// Experimental.
func MediaPackageV2Destination_Channel(channel awsmediapackagev2alpha.IChannel, endpointId MediaPackageV2EndpointId) MediaPackageV2Destination {
	_init_.Initialize()

	if err := validateMediaPackageV2Destination_ChannelParameters(channel); err != nil {
		panic(err)
	}
	var returns MediaPackageV2Destination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2Destination",
		"channel",
		[]interface{}{channel, endpointId},
		&returns,
	)

	return returns
}

