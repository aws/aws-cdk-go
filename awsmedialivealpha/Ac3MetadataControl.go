package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AC3 metadata control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ac3MetadataControl := medialive_alpha.Ac3MetadataControl_Of(jsii.String("value"))
//
// Experimental.
type Ac3MetadataControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Ac3MetadataControl
type jsiiProxy_Ac3MetadataControl struct {
	_ byte // padding
}

func (j *jsiiProxy_Ac3MetadataControl) Value() *string {
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
func Ac3MetadataControl_Of(value *string) Ac3MetadataControl {
	_init_.Initialize()

	if err := validateAc3MetadataControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Ac3MetadataControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Ac3MetadataControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Ac3MetadataControl_FOLLOW_INPUT() Ac3MetadataControl {
	_init_.Initialize()
	var returns Ac3MetadataControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3MetadataControl",
		"FOLLOW_INPUT",
		&returns,
	)
	return returns
}

func Ac3MetadataControl_USE_CONFIGURED() Ac3MetadataControl {
	_init_.Initialize()
	var returns Ac3MetadataControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3MetadataControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

