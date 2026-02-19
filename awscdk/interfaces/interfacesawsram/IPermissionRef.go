package interfacesawsram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permission.
// Experimental.
type IPermissionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Permission resource.
	// Experimental.
	PermissionRef() *PermissionReference
}

// The jsii proxy for IPermissionRef
type jsiiProxy_IPermissionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPermissionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IPermissionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPermissionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

