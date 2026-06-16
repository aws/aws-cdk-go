package interfacesawsresiliencehubv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsresiliencehubv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a System.
// Experimental.
type ISystemRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a System resource.
	// Experimental.
	SystemRef() *SystemReference
}

// The jsii proxy for ISystemRef
type jsiiProxy_ISystemRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISystemRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISystemRef) SystemRef() *SystemReference {
	var returns *SystemReference
	_jsii_.Get(
		j,
		"systemRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISystemRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISystemRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

