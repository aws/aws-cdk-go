package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS codec specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsCodecSpecification := medialive_alpha.HlsCodecSpecification_Of(jsii.String("value"))
//
// Experimental.
type HlsCodecSpecification interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsCodecSpecification
type jsiiProxy_HlsCodecSpecification struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsCodecSpecification) Value() *string {
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
func HlsCodecSpecification_Of(value *string) HlsCodecSpecification {
	_init_.Initialize()

	if err := validateHlsCodecSpecification_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsCodecSpecification

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCodecSpecification",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsCodecSpecification_RFC_4281() HlsCodecSpecification {
	_init_.Initialize()
	var returns HlsCodecSpecification
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsCodecSpecification",
		"RFC_4281",
		&returns,
	)
	return returns
}

func HlsCodecSpecification_RFC_6381() HlsCodecSpecification {
	_init_.Initialize()
	var returns HlsCodecSpecification
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsCodecSpecification",
		"RFC_6381",
		&returns,
	)
	return returns
}

