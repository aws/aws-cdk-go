// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class representing the assigment of a value to an attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   assign := appsync_alpha.NewAssign(jsii.String("attr"), jsii.String("arg"))
//
// Experimental.
type Assign interface {
	// Renders the assignment as a map element.
	// Experimental.
	PutInMap(map_ *string) *string
	// Renders the assignment as a VTL string.
	// Experimental.
	RenderAsAssignment() *string
}

// The jsii proxy struct for Assign
type jsiiProxy_Assign struct {
	_ byte // padding
}

// Experimental.
func NewAssign(attr *string, arg *string) Assign {
	_init_.Initialize()

	j := jsiiProxy_Assign{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Assign",
		[]interface{}{attr, arg},
		&j,
	)

	return &j
}

// Experimental.
func NewAssign_Override(a Assign, attr *string, arg *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Assign",
		[]interface{}{attr, arg},
		a,
	)
}

func (a *jsiiProxy_Assign) PutInMap(map_ *string) *string {
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

