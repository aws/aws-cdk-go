package awscdkscheduleralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/internal"
)

// Interface representing a created or an imported `ScheduleGroup`.
// Experimental.
type IScheduleGroup interface {
	awscdk.IResource
	// Grant the indicated permissions on this group to the given principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant delete schedule permission for schedules in this group to the given principal.
	// Experimental.
	GrantDeleteSchedules(identity awsiam.IGrantable) awsiam.Grant
	// Grant list and get schedule permissions for schedules in this group to the given principal.
	// Experimental.
	GrantReadSchedules(identity awsiam.IGrantable) awsiam.Grant
	// Grant create and update schedule permissions for schedules in this group to the given principal.
	// Experimental.
	GrantWriteSchedules(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this group schedules.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for all invocation attempts.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricAttempts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for dropped invocations when EventBridge Scheduler stops attempting to invoke the target after a schedule's retry policy has been exhausted.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for failed invocations that also failed to deliver to DLQ.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricFailedToBeSentToDLQ(errorCode *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for invocations delivered to the DLQ.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricSentToDLQ(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for delivery of failed invocations to DLQ when the payload of the event sent to the DLQ exceeds the maximum size allowed by Amazon SQS.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricSentToDLQTruncated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Emitted when the target returns an exception after EventBridge Scheduler calls the target API.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricTargetErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for invocation failures due to API throttling by the target.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricTargetThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations that were throttled because it exceeds your service quotas.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/scheduler-quotas.html
	//
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The arn of the schedule group.
	// Experimental.
	ScheduleGroupArn() *string
	// The name of the schedule group.
	// Experimental.
	ScheduleGroupName() *string
}

// The jsii proxy for IScheduleGroup
type jsiiProxy_IScheduleGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IScheduleGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) GrantDeleteSchedules(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDeleteSchedulesParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDeleteSchedules",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) GrantReadSchedules(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadSchedulesParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadSchedules",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) GrantWriteSchedules(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteSchedulesParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWriteSchedules",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricAttempts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAttemptsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAttempts",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDroppedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDropped",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricFailedToBeSentToDLQ(errorCode *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFailedToBeSentToDLQParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailedToBeSentToDLQ",
		[]interface{}{errorCode, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricSentToDLQ(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSentToDLQParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSentToDLQ",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricSentToDLQTruncated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSentToDLQTruncatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSentToDLQTruncated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricTargetErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTargetErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTargetErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricTargetThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTargetThrottledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTargetThrottled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IScheduleGroup) MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IScheduleGroup) ScheduleGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduleGroup) ScheduleGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleGroupName",
		&returns,
	)
	return returns
}

