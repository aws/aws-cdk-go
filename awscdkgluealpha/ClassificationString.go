package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Classification string given to tables with this data format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   classificationString := glue_alpha.ClassificationString_AVRO()
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html#classifier-built-in
//
// Experimental.
type ClassificationString interface {
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ClassificationString
type jsiiProxy_ClassificationString struct {
	_ byte // padding
}

func (j *jsiiProxy_ClassificationString) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewClassificationString(value *string) ClassificationString {
	_init_.Initialize()

	if err := validateNewClassificationStringParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClassificationString{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewClassificationString_Override(c ClassificationString, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		[]interface{}{value},
		c,
	)
}

func ClassificationString_AVRO() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"AVRO",
		&returns,
	)
	return returns
}

func ClassificationString_CSV() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"CSV",
		&returns,
	)
	return returns
}

func ClassificationString_JSON() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"JSON",
		&returns,
	)
	return returns
}

func ClassificationString_ORC() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"ORC",
		&returns,
	)
	return returns
}

func ClassificationString_PARQUET() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"PARQUET",
		&returns,
	)
	return returns
}

func ClassificationString_XML() ClassificationString {
	_init_.Initialize()
	var returns ClassificationString
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ClassificationString",
		"XML",
		&returns,
	)
	return returns
}

