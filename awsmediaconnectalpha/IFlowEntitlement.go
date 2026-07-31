package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for FlowEntitlement.
// Experimental.
type IFlowEntitlement interface {
	interfacesawsmediaconnect.IFlowEntitlementRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the entitlement.
	// Experimental.
	EntitlementArn() *string
}

// The jsii proxy for IFlowEntitlement
type jsiiProxy_IFlowEntitlement struct {
	internal.Type__interfacesawsmediaconnectIFlowEntitlementRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlowEntitlement) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IFlowEntitlement) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFlowEntitlement) EntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlement) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlement) FlowEntitlementRef() *interfacesawsmediaconnect.FlowEntitlementReference {
	var returns *interfacesawsmediaconnect.FlowEntitlementReference
	_jsii_.Get(
		j,
		"flowEntitlementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlement) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

