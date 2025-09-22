package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLinkedRole.
// Experimental.
type IServiceLinkedRoleRef interface {
	constructs.IConstruct
	// A reference to a ServiceLinkedRole resource.
	// Experimental.
	ServiceLinkedRoleRef() *ServiceLinkedRoleReference
}

// The jsii proxy for IServiceLinkedRoleRef
type jsiiProxy_IServiceLinkedRoleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceLinkedRoleRef) ServiceLinkedRoleRef() *ServiceLinkedRoleReference {
	var returns *ServiceLinkedRoleReference
	_jsii_.Get(
		j,
		"serviceLinkedRoleRef",
		&returns,
	)
	return returns
}

