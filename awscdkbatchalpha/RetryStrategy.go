// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Define how Jobs using this JobDefinition respond to different exit conditions.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   	}),
//   	RetryAttempts: jsii.Number(5),
//   	RetryStrategies: []retryStrategy{
//   		batch.*retryStrategy_Of(batch.Action_EXIT, batch.Reason_CANNOT_PULL_CONTAINER()),
//   	},
//   })
//   jobDefn.addRetryStrategy(batch.retryStrategy_Of(batch.Action_EXIT, batch.Reason_SPOT_INSTANCE_RECLAIMED()))
//   jobDefn.addRetryStrategy(batch.retryStrategy_Of(batch.Action_EXIT, batch.Reason_CANNOT_PULL_CONTAINER()))
//   jobDefn.addRetryStrategy(batch.retryStrategy_Of(batch.Action_EXIT, batch.Reason_Custom(&CustomReason{
//   	OnExitCode: jsii.String("40*"),
//   	OnReason: jsii.String("some reason"),
//   })))
//
// Experimental.
type RetryStrategy interface {
	// The action to take when the job exits with the Reason specified.
	// Experimental.
	Action() Action
	// If the job exits with this Reason it will trigger the specified Action.
	// Experimental.
	On() Reason
}

// The jsii proxy struct for RetryStrategy
type jsiiProxy_RetryStrategy struct {
	_ byte // padding
}

func (j *jsiiProxy_RetryStrategy) Action() Action {
	var returns Action
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RetryStrategy) On() Reason {
	var returns Reason
	_jsii_.Get(
		j,
		"on",
		&returns,
	)
	return returns
}


// Experimental.
func NewRetryStrategy(action Action, on Reason) RetryStrategy {
	_init_.Initialize()

	if err := validateNewRetryStrategyParameters(action, on); err != nil {
		panic(err)
	}
	j := jsiiProxy_RetryStrategy{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.RetryStrategy",
		[]interface{}{action, on},
		&j,
	)

	return &j
}

// Experimental.
func NewRetryStrategy_Override(r RetryStrategy, action Action, on Reason) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.RetryStrategy",
		[]interface{}{action, on},
		r,
	)
}

// Create a new RetryStrategy.
// Experimental.
func RetryStrategy_Of(action Action, on Reason) RetryStrategy {
	_init_.Initialize()

	if err := validateRetryStrategy_OfParameters(action, on); err != nil {
		panic(err)
	}
	var returns RetryStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.RetryStrategy",
		"of",
		[]interface{}{action, on},
		&returns,
	)

	return returns
}

