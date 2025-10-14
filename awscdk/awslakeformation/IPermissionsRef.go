package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permissions.
// Experimental.
type IPermissionsRef interface {
	constructs.IConstruct
	// A reference to a Permissions resource.
	// Experimental.
	PermissionsRef() *PermissionsReference
}

// The jsii proxy for IPermissionsRef
type jsiiProxy_IPermissionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPermissionsRef) PermissionsRef() *PermissionsReference {
	var returns *PermissionsReference
	_jsii_.Get(
		j,
		"permissionsRef",
		&returns,
	)
	return returns
}

