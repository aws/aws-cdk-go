package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Compression format for DecompressionProcessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decompressionProcessorCompressionFormat := awscdk.Aws_kinesisfirehose.DecompressionProcessorCompressionFormat_Of(jsii.String("compressionFormat"))
//
type DecompressionProcessorCompressionFormat interface {
	// The compression format string.
	CompressionFormat() *string
}

// The jsii proxy struct for DecompressionProcessorCompressionFormat
type jsiiProxy_DecompressionProcessorCompressionFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_DecompressionProcessorCompressionFormat) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}


// A custom compression format.
func DecompressionProcessorCompressionFormat_Of(compressionFormat *string) DecompressionProcessorCompressionFormat {
	_init_.Initialize()

	if err := validateDecompressionProcessorCompressionFormat_OfParameters(compressionFormat); err != nil {
		panic(err)
	}
	var returns DecompressionProcessorCompressionFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.DecompressionProcessorCompressionFormat",
		"of",
		[]interface{}{compressionFormat},
		&returns,
	)

	return returns
}

func DecompressionProcessorCompressionFormat_GZIP() DecompressionProcessorCompressionFormat {
	_init_.Initialize()
	var returns DecompressionProcessorCompressionFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DecompressionProcessorCompressionFormat",
		"GZIP",
		&returns,
	)
	return returns
}

