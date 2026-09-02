package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Avail settings — how SCTE-35 ad avail markers are handled.
//
// Use the static factory methods to create.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//   var poisPassword StringParameter
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	AvailSettings: medialive.AvailSettings_Esam(&EsamSettings{
//   		Pois: &PoisEndpoint{
//   			Url: jsii.String("https://pois.example.com/esam"),
//   			Username: jsii.String("pois-user"),
//   			Password: poisPassword,
//   		},
//   		AcquisitionPointId: jsii.String("acquisition-point-1"),
//   		AdAvailOffset: awscdk.Duration_Millis(jsii.Number(200)),
//   	}),
//   	Scte35SegmentationScope: medialive.Scte35SegmentationScope_SCTE35_ENABLED_OUTPUT_GROUPS(),
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type AvailSettings interface {
}

// The jsii proxy struct for AvailSettings
type jsiiProxy_AvailSettings struct {
	_ byte // padding
}

// Experimental.
func NewAvailSettings_Override(a AvailSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.AvailSettings",
		nil, // no parameters
		a,
	)
}

// Use ESAM (Event Signaling and Management) mode, signaling ad avails to an external POIS.
// Experimental.
func AvailSettings_Esam(props *EsamSettings) AvailSettings {
	_init_.Initialize()

	if err := validateAvailSettings_EsamParameters(props); err != nil {
		panic(err)
	}
	var returns AvailSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AvailSettings",
		"esam",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use SCTE-35 splice insert mode for ad avail handling.
// Experimental.
func AvailSettings_SpliceInsert(props *Scte35SpliceInsertSettings) AvailSettings {
	_init_.Initialize()

	if err := validateAvailSettings_SpliceInsertParameters(props); err != nil {
		panic(err)
	}
	var returns AvailSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AvailSettings",
		"spliceInsert",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use SCTE-35 time signal APOS mode for ad avail handling.
// Experimental.
func AvailSettings_TimeSignalApos(props *Scte35TimeSignalAposSettings) AvailSettings {
	_init_.Initialize()

	if err := validateAvailSettings_TimeSignalAposParameters(props); err != nil {
		panic(err)
	}
	var returns AvailSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AvailSettings",
		"timeSignalApos",
		[]interface{}{props},
		&returns,
	)

	return returns
}

