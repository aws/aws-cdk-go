package interfacesawsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespace.
// Experimental.
type IIdNamespaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdNamespace resource.
	// Experimental.
	IdNamespaceRef() *IdNamespaceReference
}

// The jsii proxy for IIdNamespaceRef
type jsiiProxy_IIdNamespaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIdNamespaceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIdNamespaceRef) IdNamespaceRef() *IdNamespaceReference {
	var returns *IdNamespaceReference
	_jsii_.Get(
		j,
		"idNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

