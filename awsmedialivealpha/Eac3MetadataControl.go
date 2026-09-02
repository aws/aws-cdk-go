package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 metadata control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3MetadataControl := medialive_alpha.Eac3MetadataControl_Of(jsii.String("value"))
//
// Experimental.
type Eac3MetadataControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3MetadataControl
type jsiiProxy_Eac3MetadataControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3MetadataControl) Value() *string {
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
func Eac3MetadataControl_Of(value *string) Eac3MetadataControl {
	_init_.Initialize()

	if err := validateEac3MetadataControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3MetadataControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3MetadataControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3MetadataControl_FOLLOW_INPUT() Eac3MetadataControl {
	_init_.Initialize()
	var returns Eac3MetadataControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3MetadataControl",
		"FOLLOW_INPUT",
		&returns,
	)
	return returns
}

func Eac3MetadataControl_USE_CONFIGURED() Eac3MetadataControl {
	_init_.Initialize()
	var returns Eac3MetadataControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3MetadataControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

