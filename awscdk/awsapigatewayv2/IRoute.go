package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a route.
type IRoute interface {
	awscdk.IResource
	interfacesawsapigatewayv2.IRouteRef
	// Id of the Route.
	RouteId() *string
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsapigatewayv2IRouteRef
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

func (j *jsiiProxy_IRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
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

func (j *jsiiProxy_IRoute) RouteRef() *interfacesawsapigatewayv2.RouteReference {
	var returns *interfacesawsapigatewayv2.RouteReference
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

