package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitCondition.
// Experimental.
type IWaitConditionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WaitCondition resource.
	// Experimental.
	WaitConditionRef() *WaitConditionReference
}

// The jsii proxy for IWaitConditionRef
type jsiiProxy_IWaitConditionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWaitConditionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWaitConditionRef) WaitConditionRef() *WaitConditionReference {
	var returns *WaitConditionReference
	_jsii_.Get(
		j,
		"waitConditionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaitConditionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaitConditionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

