package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A destination for an RTMP output group.
//
// Use the static factory methods to create.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Rtmp(&RtmpOutputGroupProps{
//   	Name: jsii.String("social"),
//   	Outputs: []RtmpOutputDefinition{
//   		&RtmpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("live"),
//   			Destinations: []RtmpDestination{
//   				medialive.RtmpDestination_Url(jsii.String("rtmp://rtmp.example.com/live"), jsii.String("your-stream-key")),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type RtmpDestination interface {
}

// The jsii proxy struct for RtmpDestination
type jsiiProxy_RtmpDestination struct {
	_ byte // padding
}

// Experimental.
func NewRtmpDestination_Override(r RtmpDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.RtmpDestination",
		nil, // no parameters
		r,
	)
}

// Create an RTMP destination.
// Experimental.
func RtmpDestination_Url(url *string, streamName *string, options *OutputDestinationOptions) RtmpDestination {
	_init_.Initialize()

	if err := validateRtmpDestination_UrlParameters(url, streamName, options); err != nil {
		panic(err)
	}
	var returns RtmpDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpDestination",
		"url",
		[]interface{}{url, streamName, options},
		&returns,
	)

	return returns
}

