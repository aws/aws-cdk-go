package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permission.
// Experimental.
type IPermissionRef interface {
	constructs.IConstruct
	// A reference to a Permission resource.
	// Experimental.
	PermissionRef() *PermissionReference
}

// The jsii proxy for IPermissionRef
type jsiiProxy_IPermissionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPermissionRef) PermissionRef() *PermissionReference {
	var returns *PermissionReference
	_jsii_.Get(
		j,
		"permissionRef",
		&returns,
	)
	return returns
}

