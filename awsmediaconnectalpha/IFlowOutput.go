package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Flow Output.
// Experimental.
type IFlowOutput interface {
	interfacesawsmediaconnect.IFlowOutputRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the flow output.
	// Experimental.
	FlowOutputArn() *string
}

// The jsii proxy for IFlowOutput
type jsiiProxy_IFlowOutput struct {
	internal.Type__interfacesawsmediaconnectIFlowOutputRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlowOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IFlowOutput) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFlowOutput) FlowOutputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowOutputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutput) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutput) FlowOutputRef() *interfacesawsmediaconnect.FlowOutputReference {
	var returns *interfacesawsmediaconnect.FlowOutputReference
	_jsii_.Get(
		j,
		"flowOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

