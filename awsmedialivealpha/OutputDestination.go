package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// A general URL-based output destination — an S3 bucket or an HTTP(S) endpoint.
//
// Used by HLS, MS Smooth, and CMAF Ingest output groups. Each output group's `destinations`
// prop is typed to the destination valid for its protocol, so an invalid pairing (e.g. a UDP
// destination on an HLS group) is a compile-time error rather than a deploy-time failure.
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
//   		},
//   	},
//   	AvailBlanking: &AvailBlanking{
//   		State: medialive.AvailBlankingState_ENABLED(),
//   		Image: medialive.FileLocation_FromBucket(bucket, jsii.String("slates/avail.png")),
//   	},
//   	BlackoutSlate: &BlackoutSlate{
//   		State: medialive.BlackoutSlateState_ENABLED(),
//   		Image: medialive.FileLocation_*FromBucket(bucket, jsii.String("slates/blackout.png")),
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
type OutputDestination interface {
}

// The jsii proxy struct for OutputDestination
type jsiiProxy_OutputDestination struct {
	_ byte // padding
}

// Deliver to an S3 bucket (auto-grants write to the channel role).
// Experimental.
func OutputDestination_ToBucket(bucket awss3.IBucket, prefix *string) OutputDestination {
	_init_.Initialize()

	if err := validateOutputDestination_ToBucketParameters(bucket); err != nil {
		panic(err)
	}
	var returns OutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputDestination",
		"toBucket",
		[]interface{}{bucket, prefix},
		&returns,
	)

	return returns
}

// Deliver to a raw URL — typically an `https://` origin (or an `s3ssl://` path).
// Experimental.
func OutputDestination_Url(url *string, options *OutputDestinationOptions) OutputDestination {
	_init_.Initialize()

	if err := validateOutputDestination_UrlParameters(url, options); err != nil {
		panic(err)
	}
	var returns OutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OutputDestination",
		"url",
		[]interface{}{url, options},
		&returns,
	)

	return returns
}

