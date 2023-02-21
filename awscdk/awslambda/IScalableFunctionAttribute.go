package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for scalable attributes.
type IScalableFunctionAttribute interface {
	constructs.IConstruct
	// Scale out or in based on schedule.
	ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in to keep utilization at a given level.
	//
	// The utilization is tracked by the
	// LambdaProvisionedConcurrencyUtilization metric, emitted by lambda. See:
	// https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics.html#monitoring-metrics-concurrency
	ScaleOnUtilization(options *UtilizationScalingOptions)
}

// The jsii proxy for IScalableFunctionAttribute
type jsiiProxy_IScalableFunctionAttribute struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule) {
	if err := i.validateScaleOnScheduleParameters(id, actions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"scaleOnSchedule",
		[]interface{}{id, actions},
	)
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnUtilization(options *UtilizationScalingOptions) {
	if err := i.validateScaleOnUtilizationParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"scaleOnUtilization",
		[]interface{}{options},
	)
}

