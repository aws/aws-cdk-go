package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for which all Route based classes MUST implement.
type IRoute interface {
	awscdk.IResource
	interfacesawsappmesh.IRouteRef
	// The Amazon Resource Name (ARN) for the route.
	RouteArn() *string
	// The name of the route.
	RouteName() *string
	// The VirtualRouter the Route belongs to.
	VirtualRouter() IVirtualRouter
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsappmeshIRouteRef
}

func (i *jsiiProxy_IRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRoute) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRoute) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) VirtualRouter() IVirtualRouter {
	var returns IVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) RouteRef() *interfacesawsappmesh.RouteReference {
	var returns *interfacesawsappmesh.RouteReference
	_jsii_.Get(
		j,
		"routeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

