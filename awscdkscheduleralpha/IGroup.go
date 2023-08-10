package awscdkscheduleralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/internal"
)

// Experimental.
type IGroup interface {
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
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for all invocation attempts.
	// Experimental.
	MetricAttempts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for dropped invocations when EventBridge Scheduler stops attempting to invoke the target after a schedule's retry policy has been exhausted.
	// Experimental.
	MetricDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for failed invocations that also failed to deliver to DLQ.
	// Experimental.
	MetricFailedToBeSentToDLQ(errorCode *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for invocations delivered to the DLQ.
	// Experimental.
	MetricSentToDLQ(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for delivery of failed invocations to DLQ when the payload of the event sent to the DLQ exceeds the maximum size allowed by Amazon SQS.
	// Experimental.
	MetricSentToDLQTrunacted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Emitted when the target returns an exception after EventBridge Scheduler calls the target API.
	// Experimental.
	MetricTargetErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for invocation failures due to API throttling by the target.
	// Experimental.
	MetricTargetThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations that were throttled because it exceeds your service quotas.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/scheduler-quotas.html
	//
	// Experimental.
	MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The arn of the schedule group.
	// Experimental.
	GroupArn() *string
	// The name of the schedule group.
	// Experimental.
	GroupName() *string
}

// The jsii proxy for IGroup
type jsiiProxy_IGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IGroup) GrantDeleteSchedules(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IGroup) GrantReadSchedules(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IGroup) GrantWriteSchedules(identity awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IGroup) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricAttempts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricFailedToBeSentToDLQ(errorCode *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricSentToDLQ(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricSentToDLQTrunacted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSentToDLQTrunactedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSentToDLQTrunacted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGroup) MetricTargetErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricTargetThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGroup) MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IGroup) GroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

