package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitCondition.
// Experimental.
type IWaitConditionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WaitCondition resource.
	// Experimental.
	WaitConditionRef() *WaitConditionReference
}

// The jsii proxy for IWaitConditionRef
type jsiiProxy_IWaitConditionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IWaitConditionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

