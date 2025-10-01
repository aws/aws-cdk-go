package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilePermission.
// Experimental.
type IProfilePermissionRef interface {
	constructs.IConstruct
	// A reference to a ProfilePermission resource.
	// Experimental.
	ProfilePermissionRef() *ProfilePermissionReference
}

// The jsii proxy for IProfilePermissionRef
type jsiiProxy_IProfilePermissionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProfilePermissionRef) ProfilePermissionRef() *ProfilePermissionReference {
	var returns *ProfilePermissionReference
	_jsii_.Get(
		j,
		"profilePermissionRef",
		&returns,
	)
	return returns
}

