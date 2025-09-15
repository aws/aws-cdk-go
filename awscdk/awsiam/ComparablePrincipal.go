package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
type ComparablePrincipal interface {
}

// The jsii proxy struct for ComparablePrincipal
type jsiiProxy_ComparablePrincipal struct {
	_ byte // padding
}

func NewComparablePrincipal() ComparablePrincipal {
	_init_.Initialize()

	j := jsiiProxy_ComparablePrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ComparablePrincipal",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewComparablePrincipal_Override(c ComparablePrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ComparablePrincipal",
		nil, // no parameters
		c,
	)
}

// Return the dedupeString of the given principal, if available.
func ComparablePrincipal_DedupeStringFor(x IPrincipal) *string {
	_init_.Initialize()

	if err := validateComparablePrincipal_DedupeStringForParameters(x); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ComparablePrincipal",
		"dedupeStringFor",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Whether or not the given principal is a comparable principal.
func ComparablePrincipal_IsComparablePrincipal(x IPrincipal) *bool {
	_init_.Initialize()

	if err := validateComparablePrincipal_IsComparablePrincipalParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ComparablePrincipal",
		"isComparablePrincipal",
		[]interface{}{x},
		&returns,
	)

	return returns
}

