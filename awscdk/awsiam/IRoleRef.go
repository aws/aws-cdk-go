package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Role.
// Experimental.
type IRoleRef interface {
	constructs.IConstruct
	// A reference to a Role resource.
	// Experimental.
	RoleRef() *RoleReference
}

// The jsii proxy for IRoleRef
type jsiiProxy_IRoleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRoleRef) RoleRef() *RoleReference {
	var returns *RoleReference
	_jsii_.Get(
		j,
		"roleRef",
		&returns,
	)
	return returns
}

