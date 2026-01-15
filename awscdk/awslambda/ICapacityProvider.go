package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Lambda capacity provider.
type ICapacityProvider interface {
	interfacesawslambda.ICapacityProviderRef
	awscdk.IResource
	// The ARN of the capacity provider.
	CapacityProviderArn() *string
	// The name of the capacity provider.
	CapacityProviderName() *string
}

// The jsii proxy for ICapacityProvider
type jsiiProxy_ICapacityProvider struct {
	internal.Type__interfacesawslambdaICapacityProviderRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICapacityProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICapacityProvider) CapacityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) CapacityProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) CapacityProviderRef() *interfacesawslambda.CapacityProviderReference {
	var returns *interfacesawslambda.CapacityProviderReference
	_jsii_.Get(
		j,
		"capacityProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

