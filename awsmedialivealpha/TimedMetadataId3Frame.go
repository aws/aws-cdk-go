package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CMAF Ingest timed metadata ID3 frame.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   timedMetadataId3Frame := medialive_alpha.TimedMetadataId3Frame_NONE()
//
// Experimental.
type TimedMetadataId3Frame interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimedMetadataId3Frame
type jsiiProxy_TimedMetadataId3Frame struct {
	_ byte // padding
}

func (j *jsiiProxy_TimedMetadataId3Frame) Value() *string {
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
func TimedMetadataId3Frame_Of(value *string) TimedMetadataId3Frame {
	_init_.Initialize()

	if err := validateTimedMetadataId3Frame_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimedMetadataId3Frame

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataId3Frame",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimedMetadataId3Frame_NONE() TimedMetadataId3Frame {
	_init_.Initialize()
	var returns TimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataId3Frame",
		"NONE",
		&returns,
	)
	return returns
}

func TimedMetadataId3Frame_PRIV() TimedMetadataId3Frame {
	_init_.Initialize()
	var returns TimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataId3Frame",
		"PRIV",
		&returns,
	)
	return returns
}

func TimedMetadataId3Frame_TDRL() TimedMetadataId3Frame {
	_init_.Initialize()
	var returns TimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimedMetadataId3Frame",
		"TDRL",
		&returns,
	)
	return returns
}

