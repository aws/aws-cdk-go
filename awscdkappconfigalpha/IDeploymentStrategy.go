package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
)

// Experimental.
type IDeploymentStrategy interface {
	awscdk.IResource
	// The deployment duration in minutes.
	// Experimental.
	DeploymentDurationInMinutes() *float64
	// The Amazon Resource Name (ARN) of the deployment strategy.
	// Experimental.
	DeploymentStrategyArn() *string
	// The ID of the deployment strategy.
	// Experimental.
	DeploymentStrategyId() *string
	// The description of the deployment strategy.
	// Experimental.
	Description() *string
	// The final bake time in minutes.
	// Experimental.
	FinalBakeTimeInMinutes() *float64
	// The growth factor of the deployment strategy.
	// Experimental.
	GrowthFactor() *float64
	// The growth type of the deployment strategy.
	// Experimental.
	GrowthType() GrowthType
	// The name of the deployment strategy.
	// Experimental.
	Name() *string
}

// The jsii proxy for IDeploymentStrategy
type jsiiProxy_IDeploymentStrategy struct {
	internal.Type__awscdkIResource
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

