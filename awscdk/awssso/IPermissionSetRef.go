package awssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PermissionSet.
// Experimental.
type IPermissionSetRef interface {
	constructs.IConstruct
	// A reference to a PermissionSet resource.
	// Experimental.
	PermissionSetRef() *PermissionSetReference
}

// The jsii proxy for IPermissionSetRef
type jsiiProxy_IPermissionSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPermissionSetRef) PermissionSetRef() *PermissionSetReference {
	var returns *PermissionSetReference
	_jsii_.Get(
		j,
		"permissionSetRef",
		&returns,
	)
	return returns
}

