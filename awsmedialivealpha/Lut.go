package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The S3 location of a 3D LUT (look-up table) file used by a color-correction rule.
//
// MediaLive
// reads the LUT from S3 at runtime, so the file must be in an S3 bucket — the URI must use the
// `s3://` or `s3ssl://` protocol. Unlike a `FileLocation`, a LUT has no credentials.
//
// Use the static factory methods to create one from an S3 bucket (which uses the secure
// `s3ssl://` form and auto-grants the channel role read access) or a raw S3 URL.
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
type Lut interface {
}

// The jsii proxy struct for Lut
type jsiiProxy_Lut struct {
	_ byte // padding
}

// Experimental.
func NewLut_Override(l Lut) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.Lut",
		nil, // no parameters
		l,
	)
}

// Reference a LUT file in an S3 bucket.
//
// Uses the secure `s3ssl://` form and automatically
// grants the channel role read access.
// Experimental.
func Lut_FromBucket(bucket awss3.IBucket, key *string) Lut {
	_init_.Initialize()

	if err := validateLut_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns Lut

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Lut",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Reference a LUT file by S3 URL.
//
// The URL must use the `s3://` or `s3ssl://` protocol.
// Experimental.
func Lut_Url(url *string) Lut {
	_init_.Initialize()

	if err := validateLut_UrlParameters(url); err != nil {
		panic(err)
	}
	var returns Lut

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Lut",
		"url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

