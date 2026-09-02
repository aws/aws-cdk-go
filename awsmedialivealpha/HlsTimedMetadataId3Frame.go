package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS timed metadata ID3 frame.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsTimedMetadataId3Frame := medialive_alpha.HlsTimedMetadataId3Frame_NONE()
//
// Experimental.
type HlsTimedMetadataId3Frame interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsTimedMetadataId3Frame
type jsiiProxy_HlsTimedMetadataId3Frame struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsTimedMetadataId3Frame) Value() *string {
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
func HlsTimedMetadataId3Frame_Of(value *string) HlsTimedMetadataId3Frame {
	_init_.Initialize()

	if err := validateHlsTimedMetadataId3Frame_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsTimedMetadataId3Frame

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsTimedMetadataId3Frame",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsTimedMetadataId3Frame_NONE() HlsTimedMetadataId3Frame {
	_init_.Initialize()
	var returns HlsTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsTimedMetadataId3Frame",
		"NONE",
		&returns,
	)
	return returns
}

func HlsTimedMetadataId3Frame_PRIV() HlsTimedMetadataId3Frame {
	_init_.Initialize()
	var returns HlsTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsTimedMetadataId3Frame",
		"PRIV",
		&returns,
	)
	return returns
}

func HlsTimedMetadataId3Frame_TDRL() HlsTimedMetadataId3Frame {
	_init_.Initialize()
	var returns HlsTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsTimedMetadataId3Frame",
		"TDRL",
		&returns,
	)
	return returns
}

