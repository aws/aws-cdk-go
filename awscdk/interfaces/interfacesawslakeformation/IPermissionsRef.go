package interfacesawslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permissions.
// Experimental.
type IPermissionsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Permissions resource.
	// Experimental.
	PermissionsRef() *PermissionsReference
}

// The jsii proxy for IPermissionsRef
type jsiiProxy_IPermissionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPermissionsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPermissionsRef) PermissionsRef() *PermissionsReference {
	var returns *PermissionsReference
	_jsii_.Get(
		j,
		"permissionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPermissionsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

