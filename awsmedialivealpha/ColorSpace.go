package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A color space supported for 3D-LUT color conversion in a color-correction rule.
//
// Example:
//   var stack Stack
//   var bucket IBucket
//   var input IInput
//   var video EncodeConfiguration
//   var destination OutputDestination
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	ColorCorrections: []ColorCorrection{
//   		&ColorCorrection{
//   			InputColorSpace: medialive.ColorSpace_REC_601(),
//   			OutputColorSpace: medialive.ColorSpace_REC_709(),
//   			Lut: medialive.Lut_FromBucket(bucket, jsii.String("luts/rec601-to-rec709.cube")),
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				destination,
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("video"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ColorSpace interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ColorSpace
type jsiiProxy_ColorSpace struct {
	_ byte // padding
}

func (j *jsiiProxy_ColorSpace) Value() *string {
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
func ColorSpace_Of(value *string) ColorSpace {
	_init_.Initialize()

	if err := validateColorSpace_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ColorSpace

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ColorSpace",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ColorSpace_HDR10() ColorSpace {
	_init_.Initialize()
	var returns ColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorSpace",
		"HDR10",
		&returns,
	)
	return returns
}

func ColorSpace_HLG_2020() ColorSpace {
	_init_.Initialize()
	var returns ColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorSpace",
		"HLG_2020",
		&returns,
	)
	return returns
}

func ColorSpace_REC_601() ColorSpace {
	_init_.Initialize()
	var returns ColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorSpace",
		"REC_601",
		&returns,
	)
	return returns
}

func ColorSpace_REC_709() ColorSpace {
	_init_.Initialize()
	var returns ColorSpace
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorSpace",
		"REC_709",
		&returns,
	)
	return returns
}

