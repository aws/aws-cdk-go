package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls how the `colorSpace` value is used when it is not `FOLLOW`.
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
//   			AudioSelectors: []AudioSelector{
//   				medialive.AudioSelector_ByLanguage(jsii.String("eng"), jsii.String("eng"), medialive.AudioLanguageSelectionPolicy_STRICT()),
//   			},
//   			CaptionSelectors: []CaptionSelector{
//   				medialive.CaptionSelector_Embedded(jsii.String("embedded")),
//   			},
//   			VideoSelector: &VideoSelectorSettings{
//   				ColorSpace: medialive.VideoColorSpace_HDR10(),
//   				ColorSpaceUsage: medialive.VideoColorSpaceUsage_FORCE(),
//   				SelectBy: medialive.VideoSelection_ByProgramId(jsii.Number(1)),
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
type VideoColorSpaceUsage interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for VideoColorSpaceUsage
type jsiiProxy_VideoColorSpaceUsage struct {
	_ byte // padding
}

func (j *jsiiProxy_VideoColorSpaceUsage) Value() *string {
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
func VideoColorSpaceUsage_Of(value *string) VideoColorSpaceUsage {
	_init_.Initialize()

	if err := validateVideoColorSpaceUsage_OfParameters(value); err != nil {
		panic(err)
	}
	var returns VideoColorSpaceUsage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpaceUsage",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VideoColorSpaceUsage_FALLBACK() VideoColorSpaceUsage {
	_init_.Initialize()
	var returns VideoColorSpaceUsage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpaceUsage",
		"FALLBACK",
		&returns,
	)
	return returns
}

func VideoColorSpaceUsage_FORCE() VideoColorSpaceUsage {
	_init_.Initialize()
	var returns VideoColorSpaceUsage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpaceUsage",
		"FORCE",
		&returns,
	)
	return returns
}

