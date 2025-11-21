package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The value of a constant in a component document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   componentConstantValue := imagebuilder_alpha.ComponentConstantValue_FromString(jsii.String("value"))
//
// Experimental.
type ComponentConstantValue interface {
	// The data type of the constant.
	// Experimental.
	Type() *string
	// The value of the constant.
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for ComponentConstantValue
type jsiiProxy_ComponentConstantValue struct {
	_ byte // padding
}

func (j *jsiiProxy_ComponentConstantValue) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComponentConstantValue) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponentConstantValue(type_ *string, value interface{}) ComponentConstantValue {
	_init_.Initialize()

	if err := validateNewComponentConstantValueParameters(type_, value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComponentConstantValue{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentConstantValue",
		[]interface{}{type_, value},
		&j,
	)

	return &j
}

// Experimental.
func NewComponentConstantValue_Override(c ComponentConstantValue, type_ *string, value interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentConstantValue",
		[]interface{}{type_, value},
		c,
	)
}

// Creates a string type constant in a component document.
// Experimental.
func ComponentConstantValue_FromString(value *string) ComponentConstantValue {
	_init_.Initialize()

	if err := validateComponentConstantValue_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ComponentConstantValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentConstantValue",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

