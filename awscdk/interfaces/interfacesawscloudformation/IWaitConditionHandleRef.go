package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitConditionHandle.
// Experimental.
type IWaitConditionHandleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WaitConditionHandle resource.
	// Experimental.
	WaitConditionHandleRef() *WaitConditionHandleReference
}

// The jsii proxy for IWaitConditionHandleRef
type jsiiProxy_IWaitConditionHandleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWaitConditionHandleRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWaitConditionHandleRef) WaitConditionHandleRef() *WaitConditionHandleReference {
	var returns *WaitConditionHandleReference
	_jsii_.Get(
		j,
		"waitConditionHandleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaitConditionHandleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaitConditionHandleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

