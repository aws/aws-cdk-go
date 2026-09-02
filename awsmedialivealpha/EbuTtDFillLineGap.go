package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether EBU-TT-D fills the gap between multi-line captions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ebuTtDFillLineGap := medialive_alpha.EbuTtDFillLineGap_Of(jsii.String("value"))
//
// Experimental.
type EbuTtDFillLineGap interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for EbuTtDFillLineGap
type jsiiProxy_EbuTtDFillLineGap struct {
	_ byte // padding
}

func (j *jsiiProxy_EbuTtDFillLineGap) Value() *string {
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
func EbuTtDFillLineGap_Of(value *string) EbuTtDFillLineGap {
	_init_.Initialize()

	if err := validateEbuTtDFillLineGap_OfParameters(value); err != nil {
		panic(err)
	}
	var returns EbuTtDFillLineGap

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.EbuTtDFillLineGap",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func EbuTtDFillLineGap_DISABLED() EbuTtDFillLineGap {
	_init_.Initialize()
	var returns EbuTtDFillLineGap
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.EbuTtDFillLineGap",
		"DISABLED",
		&returns,
	)
	return returns
}

func EbuTtDFillLineGap_ENABLED() EbuTtDFillLineGap {
	_init_.Initialize()
	var returns EbuTtDFillLineGap
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.EbuTtDFillLineGap",
		"ENABLED",
		&returns,
	)
	return returns
}

