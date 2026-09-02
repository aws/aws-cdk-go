package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// A reference to a file MediaLive reads at runtime — for example an input-loss slate image, an avail-blanking image, a burn-in caption font, or a color-correction LUT.
//
// Use the static factory methods to create one from an S3 bucket (which auto-grants the
// channel role read access) or from a raw URL.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   	},
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("video"),
//   			HlsSettings: medialive.HlsSettings_Standard(&StandardHlsSettingsProps{
//   				M3u8Settings: medialive.M3u8Settings_Of(&M3u8SettingsProps{
//   					Scte35Behavior: medialive.M3u8Scte35Behavior_PASSTHROUGH(),
//   					ProgramNum: jsii.Number(1),
//   				}),
//   			}),
//   		},
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				audio,
//   			},
//   			OutputName: jsii.String("audio"),
//   			HlsSettings: medialive.HlsSettings_AudioOnly(&AudioOnlyHlsSettingsProps{
//   				AudioGroupId: jsii.String("program"),
//   				AudioOnlyImage: medialive.FileLocation_FromBucket(bucket, jsii.String("art/cover.png")),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type FileLocation interface {
}

// The jsii proxy struct for FileLocation
type jsiiProxy_FileLocation struct {
	_ byte // padding
}

// Experimental.
func NewFileLocation_Override(f FileLocation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.FileLocation",
		nil, // no parameters
		f,
	)
}

// Reference a file in an S3 bucket.
//
// Automatically grants the channel role read access.
// Experimental.
func FileLocation_FromBucket(bucket awss3.IBucket, key *string) FileLocation {
	_init_.Initialize()

	if err := validateFileLocation_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns FileLocation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FileLocation",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Reference a file by URL (e.g. an `https://` endpoint or an `s3ssl://` path).
// Experimental.
func FileLocation_Url(url *string, options *FileLocationOptions) FileLocation {
	_init_.Initialize()

	if err := validateFileLocation_UrlParameters(url, options); err != nil {
		panic(err)
	}
	var returns FileLocation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FileLocation",
		"url",
		[]interface{}{url, options},
		&returns,
	)

	return returns
}

