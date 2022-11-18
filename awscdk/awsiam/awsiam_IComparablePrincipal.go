package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for principals that can be compared.
//
// This only needs to be implemented for principals that could potentially be value-equal.
// Identity-equal principals will be handled correctly by default.
type IComparablePrincipal interface {
	IPrincipal
	// Return a string format of this principal which should be identical if the two principals are the same.
	DedupeString() *string
}

// The jsii proxy for IComparablePrincipal
type jsiiProxy_IComparablePrincipal struct {
	jsiiProxy_IPrincipal
}

func (i *jsiiProxy_IComparablePrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

