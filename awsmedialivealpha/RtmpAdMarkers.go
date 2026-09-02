package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Ad marker type for an RTMP output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpAdMarkers := medialive_alpha.RtmpAdMarkers_Of(jsii.String("value"))
//
// Experimental.
type RtmpAdMarkers interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpAdMarkers
type jsiiProxy_RtmpAdMarkers struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpAdMarkers) Value() *string {
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
func RtmpAdMarkers_Of(value *string) RtmpAdMarkers {
	_init_.Initialize()

	if err := validateRtmpAdMarkers_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpAdMarkers

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpAdMarkers",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpAdMarkers_ON_CUE_POINT_SCTE35() RtmpAdMarkers {
	_init_.Initialize()
	var returns RtmpAdMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpAdMarkers",
		"ON_CUE_POINT_SCTE35",
		&returns,
	)
	return returns
}

