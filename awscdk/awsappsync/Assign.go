package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class representing the assigment of a value to an attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assign := awscdk.Aws_appsync.NewAssign(jsii.String("attr"), jsii.String("arg"))
//
type Assign interface {
	// Renders the assignment as a map element.
	PutInMap(map_ *string) *string
	// Renders the assignment as a VTL string.
	RenderAsAssignment() *string
}

// The jsii proxy struct for Assign
type jsiiProxy_Assign struct {
	_ byte // padding
}

func NewAssign(attr *string, arg *string) Assign {
	_init_.Initialize()

	if err := validateNewAssignParameters(attr, arg); err != nil {
		panic(err)
	}
	j := jsiiProxy_Assign{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Assign",
		[]interface{}{attr, arg},
		&j,
	)

	return &j
}

func NewAssign_Override(a Assign, attr *string, arg *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Assign",
		[]interface{}{attr, arg},
		a,
	)
}

func (a *jsiiProxy_Assign) PutInMap(map_ *string) *string {
	if err := a.validatePutInMapParameters(map_); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"putInMap",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Assign) RenderAsAssignment() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderAsAssignment",
		nil, // no parameters
		&returns,
	)

	return returns
}

