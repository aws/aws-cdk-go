package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RolePolicy.
// Experimental.
type IRolePolicyRef interface {
	constructs.IConstruct
	// A reference to a RolePolicy resource.
	// Experimental.
	RolePolicyRef() *RolePolicyReference
}

// The jsii proxy for IRolePolicyRef
type jsiiProxy_IRolePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRolePolicyRef) RolePolicyRef() *RolePolicyReference {
	var returns *RolePolicyReference
	_jsii_.Get(
		j,
		"rolePolicyRef",
		&returns,
	)
	return returns
}

