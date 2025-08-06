package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface representing a new or an imported Glue Job.
// Experimental.
type IJob interface {
	awsiam.IGrantable
	awscdk.IResource
	// Create a CloudWatch metric.
	// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
	//
	// Experimental.
	Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job failure.
	// Experimental.
	MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job success.
	// Experimental.
	MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Metric indicating job timeout.
	// Experimental.
	MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule triggered when something happens with this job.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the FAILED state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the SUCCEEDED state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when this job moves to the TIMEOUT state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the job.
	// Experimental.
	JobArn() *string
	// The name of the job.
	// Experimental.
	JobName() *string
}

// The jsii proxy for IJob
type jsiiProxy_IJob struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IJob) Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, type_, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, type_, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTimeoutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTimeout",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnFailureParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onFailure",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnSuccessParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onSuccess",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnTimeoutParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onTimeout",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

