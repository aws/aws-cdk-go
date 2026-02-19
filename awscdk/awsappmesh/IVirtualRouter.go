package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface which all VirtualRouter based classes MUST implement.
type IVirtualRouter interface {
	awscdk.IResource
	interfacesawsappmesh.IVirtualRouterRef
	// Add a single route to the router.
	AddRoute(id *string, props *RouteBaseProps) Route
	// The Mesh which the VirtualRouter belongs to.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualRouter.
	VirtualRouterArn() *string
	// The name of the VirtualRouter.
	VirtualRouterName() *string
}

// The jsii proxy for IVirtualRouter
type jsiiProxy_IVirtualRouter struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsappmeshIVirtualRouterRef
}

func (i *jsiiProxy_IVirtualRouter) AddRoute(id *string, props *RouteBaseProps) Route {
	if err := i.validateAddRouteParameters(id, props); err != nil {
		panic(err)
	}
	var returns Route

	_jsii_.Invoke(
		i,
		"addRoute",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVirtualRouter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IVirtualRouter) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVirtualRouter) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) VirtualRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualRouter) VirtualRouterRef() *interfacesawsappmesh.VirtualRouterReference {
	var returns *interfacesawsappmesh.VirtualRouterReference
	_jsii_.Get(
		j,
		"virtualRouterRef",
		&returns,
	)
	return returns
}

