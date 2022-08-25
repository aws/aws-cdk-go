package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awslambda/internal"
)

// Interface for scalable attributes.
// Experimental.
type IScalableFunctionAttribute interface {
	awscdk.IConstruct
	// Scale out or in based on schedule.
	// Experimental.
	ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in to keep utilization at a given level.
	//
	// The utilization is tracked by the
	// LambdaProvisionedConcurrencyUtilization metric, emitted by lambda. See:
	// https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics.html#monitoring-metrics-concurrency
	// Experimental.
	ScaleOnUtilization(options *UtilizationScalingOptions)
}

// The jsii proxy for IScalableFunctionAttribute
type jsiiProxy_IScalableFunctionAttribute struct {
	internal.Type__awscdkIConstruct
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnSchedule",
		[]interface{}{id, actions},
	)
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnUtilization(options *UtilizationScalingOptions) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnUtilization",
		[]interface{}{options},
	)
}

