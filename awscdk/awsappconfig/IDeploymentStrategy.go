package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappconfig"
	"github.com/aws/constructs-go/constructs/v10"
)

type IDeploymentStrategy interface {
	interfacesawsappconfig.IDeploymentStrategyRef
	awscdk.IResource
	// The deployment duration in minutes.
	DeploymentDurationInMinutes() *float64
	// The Amazon Resource Name (ARN) of the deployment strategy.
	DeploymentStrategyArn() *string
	// The ID of the deployment strategy.
	DeploymentStrategyId() *string
	// The description of the deployment strategy.
	Description() *string
	// The final bake time in minutes.
	FinalBakeTimeInMinutes() *float64
	// The growth factor of the deployment strategy.
	GrowthFactor() *float64
	// The growth type of the deployment strategy.
	GrowthType() GrowthType
	// The name of the deployment strategy.
	Name() *string
}

// The jsii proxy for IDeploymentStrategy
type jsiiProxy_IDeploymentStrategy struct {
	internal.Type__interfacesawsappconfigIDeploymentStrategyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDeploymentStrategy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDeploymentStrategy) DeploymentDurationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentDurationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) DeploymentStrategyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) DeploymentStrategyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) FinalBakeTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalBakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) GrowthFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"growthFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) GrowthType() GrowthType {
	var returns GrowthType
	_jsii_.Get(
		j,
		"growthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) DeploymentStrategyRef() *interfacesawsappconfig.DeploymentStrategyReference {
	var returns *interfacesawsappconfig.DeploymentStrategyReference
	_jsii_.Get(
		j,
		"deploymentStrategyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeploymentStrategy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

