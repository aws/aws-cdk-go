package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Flow Source.
// Experimental.
type IFlowSource interface {
	interfacesawsmediaconnect.IFlowSourceRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the flow source.
	// Experimental.
	FlowSourceArn() *string
	// The name of the flow source.
	// Experimental.
	FlowSourceName() *string
	// The IP address that the flow will be listening on for incoming content.
	// Experimental.
	IngestIp() *string
	// The port that the flow will be listening on for incoming content.
	// Experimental.
	SourceIngestPort() *string
}

// The jsii proxy for IFlowSource
type jsiiProxy_IFlowSource struct {
	internal.Type__interfacesawsmediaconnectIFlowSourceRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlowSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IFlowSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFlowSource) FlowSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) FlowSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) IngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) SourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) FlowSourceRef() *interfacesawsmediaconnect.FlowSourceReference {
	var returns *interfacesawsmediaconnect.FlowSourceReference
	_jsii_.Get(
		j,
		"flowSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

