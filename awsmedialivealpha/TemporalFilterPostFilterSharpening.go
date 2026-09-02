package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Post-filter sharpening for temporal filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   temporalFilterPostFilterSharpening := medialive_alpha.TemporalFilterPostFilterSharpening_AUTO()
//
// Experimental.
type TemporalFilterPostFilterSharpening interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TemporalFilterPostFilterSharpening
type jsiiProxy_TemporalFilterPostFilterSharpening struct {
	_ byte // padding
}

func (j *jsiiProxy_TemporalFilterPostFilterSharpening) Value() *string {
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
func TemporalFilterPostFilterSharpening_Of(value *string) TemporalFilterPostFilterSharpening {
	_init_.Initialize()

	if err := validateTemporalFilterPostFilterSharpening_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TemporalFilterPostFilterSharpening

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterPostFilterSharpening",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TemporalFilterPostFilterSharpening_AUTO() TemporalFilterPostFilterSharpening {
	_init_.Initialize()
	var returns TemporalFilterPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterPostFilterSharpening",
		"AUTO",
		&returns,
	)
	return returns
}

func TemporalFilterPostFilterSharpening_DISABLED() TemporalFilterPostFilterSharpening {
	_init_.Initialize()
	var returns TemporalFilterPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterPostFilterSharpening",
		"DISABLED",
		&returns,
	)
	return returns
}

func TemporalFilterPostFilterSharpening_ENABLED() TemporalFilterPostFilterSharpening {
	_init_.Initialize()
	var returns TemporalFilterPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterPostFilterSharpening",
		"ENABLED",
		&returns,
	)
	return returns
}

