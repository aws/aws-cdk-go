package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Thumbnail state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   thumbnailState := medialive_alpha.ThumbnailState_Of(jsii.String("value"))
//
// Experimental.
type ThumbnailState interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ThumbnailState
type jsiiProxy_ThumbnailState struct {
	_ byte // padding
}

func (j *jsiiProxy_ThumbnailState) Value() *string {
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
func ThumbnailState_Of(value *string) ThumbnailState {
	_init_.Initialize()

	if err := validateThumbnailState_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ThumbnailState

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ThumbnailState",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ThumbnailState_AUTO() ThumbnailState {
	_init_.Initialize()
	var returns ThumbnailState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ThumbnailState",
		"AUTO",
		&returns,
	)
	return returns
}

func ThumbnailState_DISABLED() ThumbnailState {
	_init_.Initialize()
	var returns ThumbnailState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ThumbnailState",
		"DISABLED",
		&returns,
	)
	return returns
}

