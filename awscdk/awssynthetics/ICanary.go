package awssynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssynthetics"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a CloudWatch Synthetics Canary.
type ICanary interface {
	interfacesawssynthetics.ICanaryRef
	awscdk.IResource
	// The ARN of the canary.
	CanaryArn() *string
	// The ID of the canary.
	//
	// For imported canaries, this may be the canary name as a fallback,
	// since the actual ID (a UUID) is not available when importing by name.
	CanaryId() *string
	// The name of the canary.
	CanaryName() *string
}

// The jsii proxy for ICanary
type jsiiProxy_ICanary struct {
	internal.Type__interfacesawssyntheticsICanaryRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICanary) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ICanary) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICanary) CanaryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) CanaryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) CanaryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) CanaryRef() *interfacesawssynthetics.CanaryReference {
	var returns *interfacesawssynthetics.CanaryReference
	_jsii_.Get(
		j,
		"canaryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanary) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

