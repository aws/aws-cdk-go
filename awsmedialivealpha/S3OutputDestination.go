package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// A destination for an Archive or Frame Capture output group — always an S3 bucket.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Archive(&ArchiveOutputGroupProps{
//   	Name: jsii.String("archive"),
//   	Destinations: []S3OutputDestination{
//   		medialive.S3OutputDestination_ToBucket(bucket, jsii.String("archive/recording")),
//   	},
//   	RolloverInterval: awscdk.Duration_Seconds(jsii.Number(600)),
//   	Outputs: []ArchiveOutputDefinition{
//   		&ArchiveOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("archive_out"),
//   		},
//   	},
//   })
//
// Experimental.
type S3OutputDestination interface {
}

// The jsii proxy struct for S3OutputDestination
type jsiiProxy_S3OutputDestination struct {
	_ byte // padding
}

// Deliver to an S3 bucket (auto-grants write to the channel role).
// Experimental.
func S3OutputDestination_ToBucket(bucket awss3.IBucket, prefix *string) S3OutputDestination {
	_init_.Initialize()

	if err := validateS3OutputDestination_ToBucketParameters(bucket); err != nil {
		panic(err)
	}
	var returns S3OutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.S3OutputDestination",
		"toBucket",
		[]interface{}{bucket, prefix},
		&returns,
	)

	return returns
}

// Deliver to a raw `s3ssl://` path (escape hatch when you don't have a bucket construct).
// Experimental.
func S3OutputDestination_Url(url *string) S3OutputDestination {
	_init_.Initialize()

	if err := validateS3OutputDestination_UrlParameters(url); err != nil {
		panic(err)
	}
	var returns S3OutputDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.S3OutputDestination",
		"url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

