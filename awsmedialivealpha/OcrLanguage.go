package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The OCR language to use when converting an image-based caption source to text.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ocrLanguage := medialive_alpha.OcrLanguage_DEU()
//
// Experimental.
type OcrLanguage interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for OcrLanguage
type jsiiProxy_OcrLanguage struct {
	_ byte // padding
}

func (j *jsiiProxy_OcrLanguage) Value() *string {
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
func OcrLanguage_Of(value *string) OcrLanguage {
	_init_.Initialize()

	if err := validateOcrLanguage_OfParameters(value); err != nil {
		panic(err)
	}
	var returns OcrLanguage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func OcrLanguage_DEU() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"DEU",
		&returns,
	)
	return returns
}

func OcrLanguage_ENG() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"ENG",
		&returns,
	)
	return returns
}

func OcrLanguage_FRA() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"FRA",
		&returns,
	)
	return returns
}

func OcrLanguage_NLD() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"NLD",
		&returns,
	)
	return returns
}

func OcrLanguage_POR() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"POR",
		&returns,
	)
	return returns
}

func OcrLanguage_SPA() OcrLanguage {
	_init_.Initialize()
	var returns OcrLanguage
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.OcrLanguage",
		"SPA",
		&returns,
	)
	return returns
}

