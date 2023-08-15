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
//   cloudwatch.NewAlarm(this, jsii.String("ThrottledAlarm"), &AlarmProps{
//   	Metric: sfn.StateTransitionMetric_MetricThrottledEvents(),
//   	Threshold: jsii.Number(10),
//   	EvaluationPeriods: jsii.Number(2),
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
// Default: average over 5 minutes.
//
func StateTransitionMetric_Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricParameters(metricName, props); err != nil {
		panic(err)
	}
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
// Default: average over 5 minutes.
//
func StateTransitionMetric_MetricConsumedCapacity(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricConsumedCapacityParameters(props); err != nil {
		panic(err)
	}
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
// Default: average over 5 minutes.
//
func StateTransitionMetric_MetricProvisionedBucketSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricProvisionedBucketSizeParameters(props); err != nil {
		panic(err)
	}
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
// Default: average over 5 minutes.
//
func StateTransitionMetric_MetricProvisionedRefillRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricProvisionedRefillRateParameters(props); err != nil {
		panic(err)
	}
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
// Default: sum over 5 minutes.
//
func StateTransitionMetric_MetricThrottledEvents(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricThrottledEventsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateTransitionMetric",
		"metricThrottledEvents",
		[]interface{}{props},
		&returns,
	)

	return returns
}

