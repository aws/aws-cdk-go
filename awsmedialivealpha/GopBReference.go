package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// GOP B-frame reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   gopBReference := medialive_alpha.GopBReference_Of(jsii.String("value"))
//
// Experimental.
type GopBReference interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for GopBReference
type jsiiProxy_GopBReference struct {
	_ byte // padding
}

func (j *jsiiProxy_GopBReference) Value() *string {
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
func GopBReference_Of(value *string) GopBReference {
	_init_.Initialize()

	if err := validateGopBReference_OfParameters(value); err != nil {
		panic(err)
	}
	var returns GopBReference

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.GopBReference",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GopBReference_DISABLED() GopBReference {
	_init_.Initialize()
	var returns GopBReference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.GopBReference",
		"DISABLED",
		&returns,
	)
	return returns
}

func GopBReference_ENABLED() GopBReference {
	_init_.Initialize()
	var returns GopBReference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.GopBReference",
		"ENABLED",
		&returns,
	)
	return returns
}

