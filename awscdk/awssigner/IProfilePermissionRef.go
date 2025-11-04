package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilePermission.
// Experimental.
type IProfilePermissionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ProfilePermission resource.
	// Experimental.
	ProfilePermissionRef() *ProfilePermissionReference
}

// The jsii proxy for IProfilePermissionRef
type jsiiProxy_IProfilePermissionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IProfilePermissionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilePermissionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

