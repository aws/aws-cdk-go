package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Any object that has an associated principal that a permission can be granted to.
type IGrantable interface {
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
}

// The jsii proxy for IGrantable
type jsiiProxy_IGrantable struct {
	_ byte // padding
}

func (j *jsiiProxy_IGrantable) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

