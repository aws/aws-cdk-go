package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Common job exit reasons.
//
// Example:
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
type Reason interface {
}

// The jsii proxy struct for Reason
type jsiiProxy_Reason struct {
	_ byte // padding
}

func NewReason() Reason {
	_init_.Initialize()

	j := jsiiProxy_Reason{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.Reason",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewReason_Override(r Reason) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.Reason",
		nil, // no parameters
		r,
	)
}

// A custom Reason that can match on multiple conditions.
//
// Note that all specified conditions must be met for this reason to match.
func Reason_Custom(customReasonProps *CustomReason) Reason {
	_init_.Initialize()

	if err := validateReason_CustomParameters(customReasonProps); err != nil {
		panic(err)
	}
	var returns Reason

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.Reason",
		"custom",
		[]interface{}{customReasonProps},
		&returns,
	)

	return returns
}

func Reason_CANNOT_PULL_CONTAINER() Reason {
	_init_.Initialize()
	var returns Reason
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.Reason",
		"CANNOT_PULL_CONTAINER",
		&returns,
	)
	return returns
}

func Reason_NON_ZERO_EXIT_CODE() Reason {
	_init_.Initialize()
	var returns Reason
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.Reason",
		"NON_ZERO_EXIT_CODE",
		&returns,
	)
	return returns
}

func Reason_SPOT_INSTANCE_RECLAIMED() Reason {
	_init_.Initialize()
	var returns Reason
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.Reason",
		"SPOT_INSTANCE_RECLAIMED",
		&returns,
	)
	return returns
}

