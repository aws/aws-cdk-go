package interfacesawsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WarmPool.
// Experimental.
type IWarmPoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WarmPool resource.
	// Experimental.
	WarmPoolRef() *WarmPoolReference
}

// The jsii proxy for IWarmPoolRef
type jsiiProxy_IWarmPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWarmPoolRef) WarmPoolRef() *WarmPoolReference {
	var returns *WarmPoolReference
	_jsii_.Get(
		j,
		"warmPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWarmPoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWarmPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

