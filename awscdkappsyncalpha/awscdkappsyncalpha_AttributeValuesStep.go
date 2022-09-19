// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class to allow assigning a value to an attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   var assign assign
//
//   attributeValuesStep := appsync_alpha.NewAttributeValuesStep(jsii.String("attr"), jsii.String("container"), []assign{
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

	if err := validateNewAttributeValuesStepParameters(attr, container, assignments); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttributeValuesStep{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeValuesStep_Override(a AttributeValuesStep, attr *string, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		a,
	)
}

func (a *jsiiProxy_AttributeValuesStep) Is(val *string) AttributeValues {
	if err := a.validateIsParameters(val); err != nil {
		panic(err)
	}
	var returns AttributeValues

	_jsii_.Invoke(
		a,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}

