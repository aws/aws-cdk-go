package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MPEG-2 transport stream (M2TS) container settings for an MPEG-TS output.
//
// Use `M2tsSettings.of()` to configure the transport stream produced by a UDP, Archive, SRT, or
// MediaConnect Router output. Omitting it entirely uses MediaLive's service defaults.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp_out"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("ts"),
//   			M2tsSettings: medialive.M2tsSettings_Of(&M2tsSettingsProps{
//   				Bitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   				RateMode: medialive.M2tsRateMode_VBR(),
//   				ProgramNum: jsii.Number(1),
//   				PatInterval: awscdk.Duration_Millis(jsii.Number(100)),
//   				PmtInterval: awscdk.Duration_*Millis(jsii.Number(100)),
//   				Scte35Control: medialive.M2tsScte35Control_PASSTHROUGH(),
//   				DvbSdtSettings: &DvbSdtSettings{
//   					OutputSdt: medialive.DvbSdtOutputMode_SDT_MANUAL(),
//   					ServiceName: jsii.String("My Service"),
//   					RepInterval: awscdk.Duration_*Millis(jsii.Number(2000)),
//   				},
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type M2tsSettings interface {
}

// The jsii proxy struct for M2tsSettings
type jsiiProxy_M2tsSettings struct {
	_ byte // padding
}

// Create M2TS container settings.
// Experimental.
func M2tsSettings_Of(props *M2tsSettingsProps) M2tsSettings {
	_init_.Initialize()

	if err := validateM2tsSettings_OfParameters(props); err != nil {
		panic(err)
	}
	var returns M2tsSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsSettings",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

