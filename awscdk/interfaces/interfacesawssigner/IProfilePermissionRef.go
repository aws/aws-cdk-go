package interfacesawssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilePermission.
// Experimental.
type IProfilePermissionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProfilePermission resource.
	// Experimental.
	ProfilePermissionRef() *ProfilePermissionReference
}

// The jsii proxy for IProfilePermissionRef
type jsiiProxy_IProfilePermissionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IProfilePermissionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IProfilePermissionRef) ProfilePermissionRef() *ProfilePermissionReference {
	var returns *ProfilePermissionReference
	_jsii_.Get(
		j,
		"profilePermissionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilePermissionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

