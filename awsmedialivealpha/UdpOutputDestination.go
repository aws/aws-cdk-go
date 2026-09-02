package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A destination for a UDP output group — a UDP or RTP transport endpoint.
//
// Example:
//   var video EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Rtp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("ts"),
//   			Fec: &FecOutputSettings{
//   				Mode: medialive.FecMode_COLUMN_AND_ROW(),
//   				ColumnDepth: jsii.Number(10),
//   				RowLength: jsii.Number(10),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type UdpOutputDestination interface {
}

// The jsii proxy struct for UdpOutputDestination
type jsiiProxy_UdpOutputDestination struct {
	_ byte // padding
}

// Deliver over RTP — builds `rtp://address:port`.
// Experimental.
func UdpOutputDestination_Rtp(props *TransportOutputDestinationProps) UdpOutputDestination {
	_init_.Initialize()

	if err := validateUdpOutputDestination_RtpParameters(props); err != nil {
		panic(err)
	}
	var returns UdpOutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.UdpOutputDestination",
		"rtp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Deliver over UDP — builds `udp://address:port`.
// Experimental.
func UdpOutputDestination_Udp(props *TransportOutputDestinationProps) UdpOutputDestination {
	_init_.Initialize()

	if err := validateUdpOutputDestination_UdpParameters(props); err != nil {
		panic(err)
	}
	var returns UdpOutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.UdpOutputDestination",
		"udp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Deliver to a raw transport URL (escape hatch).
// Experimental.
func UdpOutputDestination_Url(url *string) UdpOutputDestination {
	_init_.Initialize()

	if err := validateUdpOutputDestination_UrlParameters(url); err != nil {
		panic(err)
	}
	var returns UdpOutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.UdpOutputDestination",
		"url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

