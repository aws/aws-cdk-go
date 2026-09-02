package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP caption data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpCaptionData := medialive_alpha.RtmpCaptionData_ALL()
//
// Experimental.
type RtmpCaptionData interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpCaptionData
type jsiiProxy_RtmpCaptionData struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpCaptionData) Value() *string {
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
func RtmpCaptionData_Of(value *string) RtmpCaptionData {
	_init_.Initialize()

	if err := validateRtmpCaptionData_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpCaptionData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpCaptionData",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpCaptionData_ALL() RtmpCaptionData {
	_init_.Initialize()
	var returns RtmpCaptionData
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCaptionData",
		"ALL",
		&returns,
	)
	return returns
}

func RtmpCaptionData_FIELD1_608() RtmpCaptionData {
	_init_.Initialize()
	var returns RtmpCaptionData
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCaptionData",
		"FIELD1_608",
		&returns,
	)
	return returns
}

func RtmpCaptionData_FIELD1_AND_FIELD2_608() RtmpCaptionData {
	_init_.Initialize()
	var returns RtmpCaptionData
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCaptionData",
		"FIELD1_AND_FIELD2_608",
		&returns,
	)
	return returns
}

