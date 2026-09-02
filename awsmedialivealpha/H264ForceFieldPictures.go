package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 force field pictures.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264ForceFieldPictures := medialive_alpha.H264ForceFieldPictures_Of(jsii.String("value"))
//
// Experimental.
type H264ForceFieldPictures interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264ForceFieldPictures
type jsiiProxy_H264ForceFieldPictures struct {
	_ byte // padding
}

func (j *jsiiProxy_H264ForceFieldPictures) Value() *string {
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
func H264ForceFieldPictures_Of(value *string) H264ForceFieldPictures {
	_init_.Initialize()

	if err := validateH264ForceFieldPictures_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264ForceFieldPictures

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264ForceFieldPictures",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264ForceFieldPictures_DISABLED() H264ForceFieldPictures {
	_init_.Initialize()
	var returns H264ForceFieldPictures
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264ForceFieldPictures",
		"DISABLED",
		&returns,
	)
	return returns
}

func H264ForceFieldPictures_ENABLED() H264ForceFieldPictures {
	_init_.Initialize()
	var returns H264ForceFieldPictures
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264ForceFieldPictures",
		"ENABLED",
		&returns,
	)
	return returns
}

