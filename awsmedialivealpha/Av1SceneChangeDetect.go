package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 scene change detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1SceneChangeDetect := medialive_alpha.Av1SceneChangeDetect_Of(jsii.String("value"))
//
// Experimental.
type Av1SceneChangeDetect interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1SceneChangeDetect
type jsiiProxy_Av1SceneChangeDetect struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1SceneChangeDetect) Value() *string {
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
func Av1SceneChangeDetect_Of(value *string) Av1SceneChangeDetect {
	_init_.Initialize()

	if err := validateAv1SceneChangeDetect_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1SceneChangeDetect

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1SceneChangeDetect",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1SceneChangeDetect_DISABLED() Av1SceneChangeDetect {
	_init_.Initialize()
	var returns Av1SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1SceneChangeDetect",
		"DISABLED",
		&returns,
	)
	return returns
}

func Av1SceneChangeDetect_ENABLED() Av1SceneChangeDetect {
	_init_.Initialize()
	var returns Av1SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1SceneChangeDetect",
		"ENABLED",
		&returns,
	)
	return returns
}

