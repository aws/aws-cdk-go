package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Video color space.
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
type VideoColorSpace interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for VideoColorSpace
type jsiiProxy_VideoColorSpace struct {
	_ byte // padding
}

func (j *jsiiProxy_VideoColorSpace) Value() *string {
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
func VideoColorSpace_Of(value *string) VideoColorSpace {
	_init_.Initialize()

	if err := validateVideoColorSpace_OfParameters(value); err != nil {
		panic(err)
	}
	var returns VideoColorSpace

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VideoColorSpace_FOLLOW() VideoColorSpace {
	_init_.Initialize()
	var returns VideoColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"FOLLOW",
		&returns,
	)
	return returns
}

func VideoColorSpace_HDR10() VideoColorSpace {
	_init_.Initialize()
	var returns VideoColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"HDR10",
		&returns,
	)
	return returns
}

func VideoColorSpace_HLG_2020() VideoColorSpace {
	_init_.Initialize()
	var returns VideoColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"HLG_2020",
		&returns,
	)
	return returns
}

func VideoColorSpace_REC_601() VideoColorSpace {
	_init_.Initialize()
	var returns VideoColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"REC_601",
		&returns,
	)
	return returns
}

func VideoColorSpace_REC_709() VideoColorSpace {
	_init_.Initialize()
	var returns VideoColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.VideoColorSpace",
		"REC_709",
		&returns,
	)
	return returns
}

