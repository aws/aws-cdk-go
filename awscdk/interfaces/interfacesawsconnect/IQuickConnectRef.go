package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickConnect.
// Experimental.
type IQuickConnectRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QuickConnect resource.
	// Experimental.
	QuickConnectRef() *QuickConnectReference
}

// The jsii proxy for IQuickConnectRef
type jsiiProxy_IQuickConnectRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQuickConnectRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQuickConnectRef) QuickConnectRef() *QuickConnectReference {
	var returns *QuickConnectReference
	_jsii_.Get(
		j,
		"quickConnectRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuickConnectRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuickConnectRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

