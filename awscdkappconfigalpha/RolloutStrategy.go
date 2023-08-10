package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Defines the rollout strategy for a deployment strategy and includes the growth factor, deployment duration, growth type, and optionally final bake time.
//
// Example:
//   appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
//   	RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
//   		GrowthFactor: jsii.Number(20),
//   		DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   		FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(30)),
//   	}),
//   })
//
// Experimental.
type RolloutStrategy interface {
	// The deployment duration of the rollout strategy.
	// Experimental.
	DeploymentDuration() awscdk.Duration
	// The final bake time of the deployment strategy.
	// Experimental.
	FinalBakeTime() awscdk.Duration
	// The growth factor of the rollout strategy.
	// Experimental.
	GrowthFactor() *float64
	// The growth type of the rollout strategy.
	// Experimental.
	GrowthType() GrowthType
}

// The jsii proxy struct for RolloutStrategy
type jsiiProxy_RolloutStrategy struct {
	_ byte // padding
}

func (j *jsiiProxy_RolloutStrategy) DeploymentDuration() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"deploymentDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolloutStrategy) FinalBakeTime() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"finalBakeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolloutStrategy) GrowthFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"growthFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RolloutStrategy) GrowthType() GrowthType {
	var returns GrowthType
	_jsii_.Get(
		j,
		"growthType",
		&returns,
	)
	return returns
}


// Experimental.
func NewRolloutStrategy_Override(r RolloutStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		nil, // no parameters
		r,
	)
}

// Returns: An exponential rollout strategy.
// Experimental.
func RolloutStrategy_Exponential(props *RolloutStrategyProps) RolloutStrategy {
	_init_.Initialize()

	if err := validateRolloutStrategy_ExponentialParameters(props); err != nil {
		panic(err)
	}
	var returns RolloutStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"exponential",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns: A linear rollout strategy.
// Experimental.
func RolloutStrategy_Linear(props *RolloutStrategyProps) RolloutStrategy {
	_init_.Initialize()

	if err := validateRolloutStrategy_LinearParameters(props); err != nil {
		panic(err)
	}
	var returns RolloutStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"linear",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func RolloutStrategy_ALL_AT_ONCE() RolloutStrategy {
	_init_.Initialize()
	var returns RolloutStrategy
	_jsii_.StaticGet(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func RolloutStrategy_CANARY_10_PERCENT_20_MINUTES() RolloutStrategy {
	_init_.Initialize()
	var returns RolloutStrategy
	_jsii_.StaticGet(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"CANARY_10_PERCENT_20_MINUTES",
		&returns,
	)
	return returns
}

func RolloutStrategy_LINEAR_20_PERCENT_EVERY_6_MINUTES() RolloutStrategy {
	_init_.Initialize()
	var returns RolloutStrategy
	_jsii_.StaticGet(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"LINEAR_20_PERCENT_EVERY_6_MINUTES",
		&returns,
	)
	return returns
}

func RolloutStrategy_LINEAR_50_PERCENT_EVERY_30_SECONDS() RolloutStrategy {
	_init_.Initialize()
	var returns RolloutStrategy
	_jsii_.StaticGet(
		"@aws-cdk/aws-appconfig-alpha.RolloutStrategy",
		"LINEAR_50_PERCENT_EVERY_30_SECONDS",
		&returns,
	)
	return returns
}

