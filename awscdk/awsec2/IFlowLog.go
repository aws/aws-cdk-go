package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A FlowLog.
type IFlowLog interface {
	IFlowLogRef
	awscdk.IResource
	// The Id of the VPC Flow Log.
	FlowLogId() *string
}

// The jsii proxy for IFlowLog
type jsiiProxy_IFlowLog struct {
	jsiiProxy_IFlowLogRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlowLog) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IFlowLog) FlowLogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLog) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLog) FlowLogRef() *FlowLogReference {
	var returns *FlowLogReference
	_jsii_.Get(
		j,
		"flowLogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLog) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

