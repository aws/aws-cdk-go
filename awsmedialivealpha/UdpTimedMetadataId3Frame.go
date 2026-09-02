package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// UDP timed metadata ID3 frame.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   udpTimedMetadataId3Frame := medialive_alpha.UdpTimedMetadataId3Frame_NONE()
//
// Experimental.
type UdpTimedMetadataId3Frame interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for UdpTimedMetadataId3Frame
type jsiiProxy_UdpTimedMetadataId3Frame struct {
	_ byte // padding
}

func (j *jsiiProxy_UdpTimedMetadataId3Frame) Value() *string {
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
func UdpTimedMetadataId3Frame_Of(value *string) UdpTimedMetadataId3Frame {
	_init_.Initialize()

	if err := validateUdpTimedMetadataId3Frame_OfParameters(value); err != nil {
		panic(err)
	}
	var returns UdpTimedMetadataId3Frame

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.UdpTimedMetadataId3Frame",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func UdpTimedMetadataId3Frame_NONE() UdpTimedMetadataId3Frame {
	_init_.Initialize()
	var returns UdpTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpTimedMetadataId3Frame",
		"NONE",
		&returns,
	)
	return returns
}

func UdpTimedMetadataId3Frame_PRIV() UdpTimedMetadataId3Frame {
	_init_.Initialize()
	var returns UdpTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpTimedMetadataId3Frame",
		"PRIV",
		&returns,
	)
	return returns
}

func UdpTimedMetadataId3Frame_TDRL() UdpTimedMetadataId3Frame {
	_init_.Initialize()
	var returns UdpTimedMetadataId3Frame
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.UdpTimedMetadataId3Frame",
		"TDRL",
		&returns,
	)
	return returns
}

