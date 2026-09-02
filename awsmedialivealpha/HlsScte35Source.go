package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The source MediaLive ingests SCTE-35 messages from for an HLS input.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			NetworkInputSettings: &NetworkInputSettings{
//   				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
//   				HlsInputSettings: &HlsInputSettings{
//   					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
//   					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
//   				},
//   			},
//   			LogicalInterfaceNames: []*string{
//   				jsii.String("eth0"),
//   				jsii.String("eth1"),
//   			},
//   		},
//   	},
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
type HlsScte35Source interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsScte35Source
type jsiiProxy_HlsScte35Source struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsScte35Source) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func HlsScte35Source_Of(value *string) HlsScte35Source {
	_init_.Initialize()

	if err := validateHlsScte35Source_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsScte35Source

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsScte35Source",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsScte35Source_MANIFEST() HlsScte35Source {
	_init_.Initialize()
	var returns HlsScte35Source
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsScte35Source",
		"MANIFEST",
		&returns,
	)
	return returns
}

func HlsScte35Source_SEGMENTS() HlsScte35Source {
	_init_.Initialize()
	var returns HlsScte35Source
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsScte35Source",
		"SEGMENTS",
		&returns,
	)
	return returns
}

