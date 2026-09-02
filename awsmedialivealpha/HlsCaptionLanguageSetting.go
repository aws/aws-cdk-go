package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS caption language setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsCaptionLanguageSetting := medialive_alpha.HlsCaptionLanguageSetting_INSERT()
//
// Experimental.
type HlsCaptionLanguageSetting interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsCaptionLanguageSetting
type jsiiProxy_HlsCaptionLanguageSetting struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsCaptionLanguageSetting) Value() *string {
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
func HlsCaptionLanguageSetting_Of(value *string) HlsCaptionLanguageSetting {
	_init_.Initialize()

	if err := validateHlsCaptionLanguageSetting_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsCaptionLanguageSetting

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCaptionLanguageSetting",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsCaptionLanguageSetting_INSERT() HlsCaptionLanguageSetting {
	_init_.Initialize()
	var returns HlsCaptionLanguageSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsCaptionLanguageSetting",
		"INSERT",
		&returns,
	)
	return returns
}

func HlsCaptionLanguageSetting_NONE() HlsCaptionLanguageSetting {
	_init_.Initialize()
	var returns HlsCaptionLanguageSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsCaptionLanguageSetting",
		"NONE",
		&returns,
	)
	return returns
}

func HlsCaptionLanguageSetting_OMIT() HlsCaptionLanguageSetting {
	_init_.Initialize()
	var returns HlsCaptionLanguageSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsCaptionLanguageSetting",
		"OMIT",
		&returns,
	)
	return returns
}

