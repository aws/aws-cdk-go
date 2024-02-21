package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an SNS topic.
type ITopic interface {
	awscodestarnotifications.INotificationRuleTarget
	awscdk.IResource
	// Subscribe some endpoint to this topic.
	AddSubscription(subscription ITopicSubscription) Subscription
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToResourcePolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant topic publishing permissions to the given identity.
	GrantPublish(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Enables content-based deduplication for FIFO topics.
	ContentBasedDeduplication() *bool
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	Fifo() *bool
	// The ARN of the topic.
	TopicArn() *string
	// The name of the topic.
	TopicName() *string
}

// The jsii proxy for ITopic
type jsiiProxy_ITopic struct {
	internal.Type__awscodestarnotificationsINotificationRuleTarget
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ITopic) AddSubscription(subscription ITopicSubscription) Subscription {
	if err := i.validateAddSubscriptionParameters(subscription); err != nil {
		panic(err)
	}
	var returns Subscription

	_jsii_.Invoke(
		i,
		"addSubscription",
		[]interface{}{subscription},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) GrantPublish(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPublishParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublish",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ITopic) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfMessagesPublishedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfMessagesPublished",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfNotificationsDeliveredParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsDelivered",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfNotificationsFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfNotificationsFilteredOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfNotificationsFilteredOutInvalidAttributesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOutInvalidAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfNotificationsFilteredOutNoMessageAttributesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfNotificationsFilteredOutNoMessageAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPublishSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPublishSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSMSMonthToDateSpentUSDParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSMSMonthToDateSpentUSD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSMSSuccessRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSMSSuccessRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITopic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ITopic) BindAsNotificationRuleTarget(scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	if err := i.validateBindAsNotificationRuleTargetParameters(scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITopic) ContentBasedDeduplication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

