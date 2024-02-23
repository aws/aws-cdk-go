package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines a JSON Schema validator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   jsonSchemaValidator := appconfig_alpha.JsonSchemaValidator_FromFile(jsii.String("inputPath"))
//
// Deprecated.
type JsonSchemaValidator interface {
	IValidator
	// The content of the validator.
	// Deprecated.
	Content() *string
	// The type of validator.
	// Deprecated.
	Type() ValidatorType
}

// The jsii proxy struct for JsonSchemaValidator
type jsiiProxy_JsonSchemaValidator struct {
	jsiiProxy_IValidator
}

func (j *jsiiProxy_JsonSchemaValidator) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonSchemaValidator) Type() ValidatorType {
	var returns ValidatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Deprecated.
func NewJsonSchemaValidator_Override(j JsonSchemaValidator) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		nil, // no parameters
		j,
	)
}

// Defines a JSON Schema validator from a file.
// Deprecated.
func JsonSchemaValidator_FromFile(inputPath *string) JsonSchemaValidator {
	_init_.Initialize()

	if err := validateJsonSchemaValidator_FromFileParameters(inputPath); err != nil {
		panic(err)
	}
	var returns JsonSchemaValidator

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		"fromFile",
		[]interface{}{inputPath},
		&returns,
	)

	return returns
}

// Defines a JSON Schema validator from inline code.
// Deprecated.
func JsonSchemaValidator_FromInline(code *string) JsonSchemaValidator {
	_init_.Initialize()

	if err := validateJsonSchemaValidator_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns JsonSchemaValidator

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

