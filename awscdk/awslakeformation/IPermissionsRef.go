package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permissions.
// Experimental.
type IPermissionsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Permissions resource.
	// Experimental.
	PermissionsRef() *PermissionsReference
}

// The jsii proxy for IPermissionsRef
type jsiiProxy_IPermissionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IPermissionsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPermissionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

