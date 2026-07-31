package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Router Network Interface.
// Experimental.
type IRouterNetworkInterface interface {
	awscdk.IResource
	interfacesawsmediaconnect.IRouterNetworkInterfaceRef
	// The date and time the router network interface was created.
	// Experimental.
	CreatedAt() *string
	// The Amazon Resource Name (ARN) of the router network interface.
	// Experimental.
	RouterNetworkInterfaceArn() *string
	// The unique identifier of the router network interface.
	// Experimental.
	RouterNetworkInterfaceId() *string
	// The date and time the router network interface was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IRouterNetworkInterface
type jsiiProxy_IRouterNetworkInterface struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsmediaconnectIRouterNetworkInterfaceRef
}

func (i *jsiiProxy_IRouterNetworkInterface) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRouterNetworkInterface) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRouterNetworkInterface) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) RouterNetworkInterfaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerNetworkInterfaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) RouterNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) RouterNetworkInterfaceRef() *interfacesawsmediaconnect.RouterNetworkInterfaceReference {
	var returns *interfacesawsmediaconnect.RouterNetworkInterfaceReference
	_jsii_.Get(
		j,
		"routerNetworkInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterNetworkInterface) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

