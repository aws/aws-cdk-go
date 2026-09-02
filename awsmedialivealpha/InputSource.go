package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// A source for a pull-type input.
//
// Use the static factory methods to create.
//
// Example:
//   var stack Stack
//   var bucket IBucket
//
//
//   medialive.NewInput(stack, jsii.String("FileInput"), &InputProps{
//   	InputName: jsii.String("mp4-file"),
//   	Input: medialive.InputConfiguration_Mp4File([]InputSource{
//   		medialive.InputSource_FromBucket(bucket, jsii.String("media/input.mp4")),
//   	}),
//   })
//
// Experimental.
type InputSource interface {
}

// The jsii proxy struct for InputSource
type jsiiProxy_InputSource struct {
	_ byte // padding
}

// Experimental.
func NewInputSource_Override(i InputSource) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.InputSource",
		nil, // no parameters
		i,
	)
}

// Create a source from an S3 bucket.
//
// Automatically grants read access
// to the channel's role.
// Experimental.
func InputSource_FromBucket(bucket awss3.IBucket, key *string) InputSource {
	_init_.Initialize()

	if err := validateInputSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns InputSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Create a source from a URL.
// Experimental.
func InputSource_Url(url *string, options *InputSourceOptions) InputSource {
	_init_.Initialize()

	if err := validateInputSource_UrlParameters(url, options); err != nil {
		panic(err)
	}
	var returns InputSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputSource",
		"url",
		[]interface{}{url, options},
		&returns,
	)

	return returns
}

