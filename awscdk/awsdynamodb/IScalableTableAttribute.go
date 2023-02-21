package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
)

// Interface for scalable attributes.
type IScalableTableAttribute interface {
	// Add scheduled scaling for this scaling attribute.
	ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in to keep utilization at a given level.
	ScaleOnUtilization(props *UtilizationScalingProps)
}

// The jsii proxy for IScalableTableAttribute
type jsiiProxy_IScalableTableAttribute struct {
	_ byte // padding
}

func (i *jsiiProxy_IScalableTableAttribute) ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule) {
	if err := i.validateScaleOnScheduleParameters(id, actions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"scaleOnSchedule",
		[]interface{}{id, actions},
	)
}

func (i *jsiiProxy_IScalableTableAttribute) ScaleOnUtilization(props *UtilizationScalingProps) {
	if err := i.validateScaleOnUtilizationParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"scaleOnUtilization",
		[]interface{}{props},
	)
}

