package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitCondition.
// Experimental.
type IWaitConditionRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a WaitCondition resource.
	// Experimental.
	WaitConditionRef() *WaitConditionReference
}

// The jsii proxy for IWaitConditionRef
type jsiiProxy_IWaitConditionRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
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

func (j *jsiiProxy_IWaitConditionRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
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

