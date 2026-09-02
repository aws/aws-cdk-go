package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 quality level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264QualityLevel := medialive_alpha.H264QualityLevel_Of(jsii.String("value"))
//
// Experimental.
type H264QualityLevel interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264QualityLevel
type jsiiProxy_H264QualityLevel struct {
	_ byte // padding
}

func (j *jsiiProxy_H264QualityLevel) Value() *string {
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
func H264QualityLevel_Of(value *string) H264QualityLevel {
	_init_.Initialize()

	if err := validateH264QualityLevel_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264QualityLevel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264QualityLevel",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264QualityLevel_ENHANCED_QUALITY() H264QualityLevel {
	_init_.Initialize()
	var returns H264QualityLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264QualityLevel",
		"ENHANCED_QUALITY",
		&returns,
	)
	return returns
}

func H264QualityLevel_STANDARD_QUALITY() H264QualityLevel {
	_init_.Initialize()
	var returns H264QualityLevel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264QualityLevel",
		"STANDARD_QUALITY",
		&returns,
	)
	return returns
}

