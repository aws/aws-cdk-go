package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper class for working with `IComparablePrincipal`s.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comparablePrincipal := awscdk.Aws_iam.NewComparablePrincipal()
//
// Experimental.
type ComparablePrincipal interface {
}

// The jsii proxy struct for ComparablePrincipal
type jsiiProxy_ComparablePrincipal struct {
	_ byte // padding
}

// Experimental.
func NewComparablePrincipal() ComparablePrincipal {
	_init_.Initialize()

	j := jsiiProxy_ComparablePrincipal{}

	_jsii_.Create(
		"monocdk.aws_iam.ComparablePrincipal",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewComparablePrincipal_Override(c ComparablePrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.ComparablePrincipal",
		nil, // no parameters
		c,
	)
}

// Return the dedupeString of the given principal, if available.
// Experimental.
func ComparablePrincipal_DedupeStringFor(x IPrincipal) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ComparablePrincipal",
		"dedupeStringFor",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Whether or not the given principal is a comparable principal.
// Experimental.
func ComparablePrincipal_IsComparablePrincipal(x IPrincipal) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ComparablePrincipal",
		"isComparablePrincipal",
		[]interface{}{x},
		&returns,
	)

	return returns
}

