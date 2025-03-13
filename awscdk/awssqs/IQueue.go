package awssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
)

// Represents an SQS queue.
type IQueue interface {
	awscdk.IResource
	// Adds a statement to the IAM resource policy associated with this queue.
	//
	// If this queue was created in this stack (`new Queue`), a queue policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the queue is imported (`Queue.import`), then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant the actions defined in queueActions to the identity Principal given on this SQS queue resource.
	Grant(grantee awsiam.IGrantable, queueActions ...*string) awsiam.Grant
	// Grant permissions to consume messages from a queue.
	//
	// This will grant the following permissions:
	//
	//   - sqs:ChangeMessageVisibility
	//   - sqs:DeleteMessage
	//   - sqs:ReceiveMessage
	//   - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant
	// Grant an IAM principal permissions to purge all messages from the queue.
	//
	// This will grant the following permissions:
	//
	//  - sqs:PurgeQueue
	//  - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantPurge(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to send messages to a queue to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - sqs:SendMessage
	//  - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantSendMessages(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Queue.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The approximate age of the oldest non-deleted message in the queue.
	//
	// Maximum over 5 minutes.
	MetricApproximateAgeOfOldestMessage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages in the queue that are delayed and not available for reading immediately.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesDelayed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that are in flight.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesNotVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages available for retrieval from the queue.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of ReceiveMessage API calls that did not return a message.
	//
	// Sum over 5 minutes.
	MetricNumberOfEmptyReceives(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages deleted from the queue.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesDeleted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages returned by calls to the ReceiveMessage action.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesReceived(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages added to a queue.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesSent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The size of messages added to a queue.
	//
	// Average over 5 minutes.
	MetricSentMessageSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// If this queue is server-side encrypted, this is the KMS encryption key.
	EncryptionMasterKey() awskms.IKey
	// Whether the contents of the queue are encrypted, and by what type of key.
	EncryptionType() QueueEncryption
	// Whether this queue is an Amazon SQS FIFO queue.
	//
	// If false, this is a standard queue.
	Fifo() *bool
	// The ARN of this queue.
	QueueArn() *string
	// The name of this queue.
	QueueName() *string
	// The URL of this queue.
	QueueUrl() *string
}

// The jsii proxy for IQueue
type jsiiProxy_IQueue struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IQueue) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
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

func (i *jsiiProxy_IQueue) Grant(grantee awsiam.IGrantable, queueActions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range queueActions {
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

func (i *jsiiProxy_IQueue) GrantConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConsumeMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConsumeMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) GrantPurge(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPurgeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPurge",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) GrantSendMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantSendMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantSendMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IQueue) MetricApproximateAgeOfOldestMessage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricApproximateAgeOfOldestMessageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricApproximateAgeOfOldestMessage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricApproximateNumberOfMessagesDelayed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricApproximateNumberOfMessagesDelayedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricApproximateNumberOfMessagesDelayed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricApproximateNumberOfMessagesNotVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricApproximateNumberOfMessagesNotVisibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricApproximateNumberOfMessagesNotVisible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricApproximateNumberOfMessagesVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricApproximateNumberOfMessagesVisibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricApproximateNumberOfMessagesVisible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricNumberOfEmptyReceives(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfEmptyReceivesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfEmptyReceives",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricNumberOfMessagesDeleted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfMessagesDeletedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfMessagesDeleted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricNumberOfMessagesReceived(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfMessagesReceivedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfMessagesReceived",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricNumberOfMessagesSent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfMessagesSentParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfMessagesSent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IQueue) MetricSentMessageSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSentMessageSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSentMessageSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IQueue) EncryptionMasterKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionMasterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueue) EncryptionType() QueueEncryption {
	var returns QueueEncryption
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueue) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueue) QueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueue) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueue) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

