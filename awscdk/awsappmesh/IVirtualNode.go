package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface which all VirtualNode based classes must implement.
type IVirtualNode interface {
	awscdk.IResource
	interfacesawsappmesh.IVirtualNodeRef
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualNode belongs to.
	Mesh() IMesh
	// The Amazon Resource Name belonging to the VirtualNode.
	//
	// Set this value as the APPMESH_VIRTUAL_NODE_NAME environment variable for
	// your task group's Envoy proxy container in your task definition or pod
	// spec.
	VirtualNodeArn() *string
	// The name of the VirtualNode.
	VirtualNodeName() *string
}

// The jsii proxy for IVirtualNode
type jsiiProxy_IVirtualNode struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsappmeshIVirtualNodeRef
}

func (i *jsiiProxy_IVirtualNode) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStreamAggregatedResourcesParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVirtualNode) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IVirtualNode) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVirtualNode) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) VirtualNodeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) VirtualNodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualNode) VirtualNodeRef() *interfacesawsappmesh.VirtualNodeReference {
	var returns *interfacesawsappmesh.VirtualNodeReference
	_jsii_.Get(
		j,
		"virtualNodeRef",
		&returns,
	)
	return returns
}

