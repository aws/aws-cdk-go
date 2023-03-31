package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
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
// Experimental.
type StateTransitionMetric interface {
}

// The jsii proxy struct for StateTransitionMetric
type jsiiProxy_StateTransitionMetric struct {
	_ byte // padding
}

// Experimental.
func NewStateTransitionMetric() StateTransitionMetric {
	_init_.Initialize()

	j := jsiiProxy_StateTransitionMetric{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStateTransitionMetric_Override(s StateTransitionMetric) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		nil, // no parameters
		s,
	)
}

// Return the given named metric for the service's state transition metrics.
// Experimental.
func StateTransitionMetric_Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of available state transitions per second.
// Experimental.
func StateTransitionMetric_MetricConsumedCapacity(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricConsumedCapacityParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		"metricConsumedCapacity",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of available state transitions.
// Experimental.
func StateTransitionMetric_MetricProvisionedBucketSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricProvisionedBucketSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		"metricProvisionedBucketSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the provisioned steady-state execution rate.
// Experimental.
func StateTransitionMetric_MetricProvisionedRefillRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricProvisionedRefillRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		"metricProvisionedRefillRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled state transitions.
// Experimental.
func StateTransitionMetric_MetricThrottledEvents(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateStateTransitionMetric_MetricThrottledEventsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateTransitionMetric",
		"metricThrottledEvents",
		[]interface{}{props},
		&returns,
	)

	return returns
}

