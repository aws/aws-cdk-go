package interfacesawschime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppInstanceBot.
// Experimental.
type IAppInstanceBotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AppInstanceBot resource.
	// Experimental.
	AppInstanceBotRef() *AppInstanceBotReference
}

// The jsii proxy for IAppInstanceBotRef
type jsiiProxy_IAppInstanceBotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAppInstanceBotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAppInstanceBotRef) AppInstanceBotRef() *AppInstanceBotReference {
	var returns *AppInstanceBotReference
	_jsii_.Get(
		j,
		"appInstanceBotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppInstanceBotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppInstanceBotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

