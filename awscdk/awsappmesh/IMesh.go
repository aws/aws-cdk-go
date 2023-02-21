package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
)

// Interface which all Mesh based classes MUST implement.
type IMesh interface {
	awscdk.IResource
	// Creates a new VirtualGateway in this Mesh.
	//
	// Note that the Gateway is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	// Creates a new VirtualNode in this Mesh.
	//
	// Note that the Node is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	// Creates a new VirtualRouter in this Mesh.
	//
	// Note that the Router is created in the same Stack that this Mesh belongs to,
	// which might be different than the current stack.
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
	// The Amazon Resource Name (ARN) of the AppMesh mesh.
	MeshArn() *string
	// The name of the AppMesh mesh.
	MeshName() *string
}

// The jsii proxy for IMesh
type jsiiProxy_IMesh struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IMesh) AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway {
	if err := i.validateAddVirtualGatewayParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualGateway

	_jsii_.Invoke(
		i,
		"addVirtualGateway",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMesh) AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode {
	if err := i.validateAddVirtualNodeParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualNode

	_jsii_.Invoke(
		i,
		"addVirtualNode",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMesh) AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter {
	if err := i.validateAddVirtualRouterParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualRouter

	_jsii_.Invoke(
		i,
		"addVirtualRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMesh) MeshArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMesh) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

