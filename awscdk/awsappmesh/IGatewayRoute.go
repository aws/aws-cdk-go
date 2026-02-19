package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for which all GatewayRoute based classes MUST implement.
type IGatewayRoute interface {
	interfacesawsappmesh.IGatewayRouteRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the GatewayRoute.
	GatewayRouteArn() *string
	// The name of the GatewayRoute.
	GatewayRouteName() *string
	// The VirtualGateway the GatewayRoute belongs to.
	VirtualGateway() IVirtualGateway
}

// The jsii proxy for IGatewayRoute
type jsiiProxy_IGatewayRoute struct {
	internal.Type__interfacesawsappmeshIGatewayRouteRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGatewayRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGatewayRoute) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGatewayRoute) GatewayRouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) VirtualGateway() IVirtualGateway {
	var returns IVirtualGateway
	_jsii_.Get(
		j,
		"virtualGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteRef() *interfacesawsappmesh.GatewayRouteReference {
	var returns *interfacesawsappmesh.GatewayRouteReference
	_jsii_.Get(
		j,
		"gatewayRouteRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

