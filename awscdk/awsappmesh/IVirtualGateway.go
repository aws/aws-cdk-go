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

// Interface which all Virtual Gateway based classes must implement.
type IVirtualGateway interface {
	awscdk.IResource
	interfacesawsappmesh.IVirtualGatewayRef
	// Utility method to add a new GatewayRoute to the VirtualGateway.
	AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	//
	// [disable-awslint:no-grants].
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualGateway belongs to.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualGateway.
	VirtualGatewayArn() *string
	// Name of the VirtualGateway.
	VirtualGatewayName() *string
}

// The jsii proxy for IVirtualGateway
type jsiiProxy_IVirtualGateway struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsappmeshIVirtualGatewayRef
}

func (i *jsiiProxy_IVirtualGateway) AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute {
	if err := i.validateAddGatewayRouteParameters(id, route); err != nil {
		panic(err)
	}
	var returns GatewayRoute

	_jsii_.Invoke(
		i,
		"addGatewayRoute",
		[]interface{}{id, route},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVirtualGateway) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IVirtualGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IVirtualGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVirtualGateway) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayRef() *interfacesawsappmesh.VirtualGatewayReference {
	var returns *interfacesawsappmesh.VirtualGatewayReference
	_jsii_.Get(
		j,
		"virtualGatewayRef",
		&returns,
	)
	return returns
}

