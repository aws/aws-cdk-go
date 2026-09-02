package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls which output groups receive SCTE-35 segmentation cues.
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
type Scte35SegmentationScope interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Scte35SegmentationScope
type jsiiProxy_Scte35SegmentationScope struct {
	_ byte // padding
}

func (j *jsiiProxy_Scte35SegmentationScope) Value() *string {
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
func Scte35SegmentationScope_Of(value *string) Scte35SegmentationScope {
	_init_.Initialize()

	if err := validateScte35SegmentationScope_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Scte35SegmentationScope

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Scte35SegmentationScope",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Scte35SegmentationScope_ALL_OUTPUT_GROUPS() Scte35SegmentationScope {
	_init_.Initialize()
	var returns Scte35SegmentationScope
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35SegmentationScope",
		"ALL_OUTPUT_GROUPS",
		&returns,
	)
	return returns
}

func Scte35SegmentationScope_SCTE35_ENABLED_OUTPUT_GROUPS() Scte35SegmentationScope {
	_init_.Initialize()
	var returns Scte35SegmentationScope
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35SegmentationScope",
		"SCTE35_ENABLED_OUTPUT_GROUPS",
		&returns,
	)
	return returns
}

