package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SCTE-20 detection mode for an embedded caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   scte20Detection := medialive_alpha.Scte20Detection_Of(jsii.String("value"))
//
// Experimental.
type Scte20Detection interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Scte20Detection
type jsiiProxy_Scte20Detection struct {
	_ byte // padding
}

func (j *jsiiProxy_Scte20Detection) Value() *string {
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
func Scte20Detection_Of(value *string) Scte20Detection {
	_init_.Initialize()

	if err := validateScte20Detection_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Scte20Detection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Scte20Detection",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Scte20Detection_AUTO() Scte20Detection {
	_init_.Initialize()
	var returns Scte20Detection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte20Detection",
		"AUTO",
		&returns,
	)
	return returns
}

func Scte20Detection_OFF() Scte20Detection {
	_init_.Initialize()
	var returns Scte20Detection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte20Detection",
		"OFF",
		&returns,
	)
	return returns
}

