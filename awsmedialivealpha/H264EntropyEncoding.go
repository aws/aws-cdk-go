package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 entropy encoding mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264EntropyEncoding := medialive_alpha.H264EntropyEncoding_Of(jsii.String("value"))
//
// Experimental.
type H264EntropyEncoding interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264EntropyEncoding
type jsiiProxy_H264EntropyEncoding struct {
	_ byte // padding
}

func (j *jsiiProxy_H264EntropyEncoding) Value() *string {
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
func H264EntropyEncoding_Of(value *string) H264EntropyEncoding {
	_init_.Initialize()

	if err := validateH264EntropyEncoding_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264EntropyEncoding

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264EntropyEncoding",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264EntropyEncoding_CABAC() H264EntropyEncoding {
	_init_.Initialize()
	var returns H264EntropyEncoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264EntropyEncoding",
		"CABAC",
		&returns,
	)
	return returns
}

func H264EntropyEncoding_CAVLC() H264EntropyEncoding {
	_init_.Initialize()
	var returns H264EntropyEncoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264EntropyEncoding",
		"CAVLC",
		&returns,
	)
	return returns
}

