package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Metrics on the rate limiting performed on state machine execution.
//
// These rate limits are shared across all state machines.
//
// Example:
//   cloudwatch.NewAlarm(this, jsii.String("ThrottledAlarm"), &alarmProps{
//   	metric: sfn.stateTransitionMetric.metricThrottledEvents(),
//   	threshold: jsii.Number(10),
//   	evaluationPeriods: jsii.Number(2),
//   })
//
type StateTransitionMetric interface {
}

// The jsii proxy struct for StateTransitionMetric
type jsiiProxy_StateTransitionMetric struct {
	_ byte // padding
}

func NewStateTransitionMetric() StateTransitionMetric {
	_init_.Initialize()

	j := jsiiProxy_StateTransitionMetric{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStateTransitionMetric_Override(s StateTransitionMetric) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		nil, // no parameters
		s,
	)
}

// Return the given named metric for the service's state transition metrics.
func StateTransitionMetric_Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of available state transitions per second.
func StateTransitionMetric_MetricConsumedCapacity(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metricConsumedCapacity",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of available state transitions.
func StateTransitionMetric_MetricProvisionedBucketSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metricProvisionedBucketSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the provisioned steady-state execution rate.
func StateTransitionMetric_MetricProvisionedRefillRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metricProvisionedRefillRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled state transitions.
func StateTransitionMetric_MetricThrottledEvents(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metricThrottledEvents",
		[]interface{}{props},
		&returns,
	)

	return returns
}

