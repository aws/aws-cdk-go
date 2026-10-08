package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Batch Consumable Resource.
type IConsumableResource interface {
	interfacesawsbatch.IConsumableResourceRef
	awscdk.IResource
	// The ARN of this consumable resource.
	ConsumableResourceArn() *string
	// The name of this consumable resource.
	ConsumableResourceName() *string
	// Collection of grant methods for this consumable resource.
	Grants() ConsumableResourceGrants
}

// The jsii proxy for IConsumableResource
type jsiiProxy_IConsumableResource struct {
	internal.Type__interfacesawsbatchIConsumableResourceRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IConsumableResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IConsumableResource) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConsumableResource) ConsumableResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumableResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) ConsumableResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumableResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) Grants() ConsumableResourceGrants {
	var returns ConsumableResourceGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) ConsumableResourceRef() *interfacesawsbatch.ConsumableResourceReference {
	var returns *interfacesawsbatch.ConsumableResourceReference
	_jsii_.Get(
		j,
		"consumableResourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

