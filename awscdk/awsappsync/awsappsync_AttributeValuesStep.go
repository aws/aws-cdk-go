package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class to allow assigning a value to an attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign assign
//
//   attributeValuesStep := awscdk.Aws_appsync.NewAttributeValuesStep(jsii.String("attr"), jsii.String("container"), []assign{
//   	assign,
//   })
//
// Experimental.
type AttributeValuesStep interface {
	// Assign the value to the current attribute.
	// Experimental.
	Is(val *string) AttributeValues
}

// The jsii proxy struct for AttributeValuesStep
type jsiiProxy_AttributeValuesStep struct {
	_ byte // padding
}

// Experimental.
func NewAttributeValuesStep(attr *string, container *string, assignments *[]Assign) AttributeValuesStep {
	_init_.Initialize()

	j := jsiiProxy_AttributeValuesStep{}

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeValuesStep_Override(a AttributeValuesStep, attr *string, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		a,
	)
}

func (a *jsiiProxy_AttributeValuesStep) Is(val *string) AttributeValues {
	var returns AttributeValues

	_jsii_.Invoke(
		a,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}

