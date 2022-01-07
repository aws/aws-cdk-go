package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// The action to send data to Amazon CloudWatch Logs.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchLogsAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchLogsAction
type jsiiProxy_CloudWatchLogsAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchLogsAction(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) CloudWatchLogsAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchLogsAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogsAction_Override(c CloudWatchLogsAction, logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		c,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (c *jsiiProxy_CloudWatchLogsAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for CloudWatch Logs.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchLogsActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// The action to capture an Amazon CloudWatch metric.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchPutMetricAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchPutMetricAction
type jsiiProxy_CloudWatchPutMetricAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchPutMetricAction(props *CloudWatchPutMetricActionProps) CloudWatchPutMetricAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchPutMetricAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchPutMetricAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchPutMetricAction_Override(c CloudWatchPutMetricAction, props *CloudWatchPutMetricActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchPutMetricAction",
		[]interface{}{props},
		c,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (c *jsiiProxy_CloudWatchPutMetricAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for CloudWatch metric.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchPutMetricActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The CloudWatch metric name.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricName *string `json:"metricName"`
	// The CloudWatch metric namespace name.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricNamespace *string `json:"metricNamespace"`
	// The metric unit supported by CloudWatch.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricUnit *string `json:"metricUnit"`
	// A string that contains the CloudWatch metric value.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricValue *string `json:"metricValue"`
	// A string that contains the timestamp, expressed in seconds in Unix epoch time.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricTimestamp *string `json:"metricTimestamp"`
}

// The action to change the state of an Amazon CloudWatch alarm.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchSetAlarmStateAction interface {
	awsiot.IAction
	Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchSetAlarmStateAction
type jsiiProxy_CloudWatchSetAlarmStateAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchSetAlarmStateAction(alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) CloudWatchSetAlarmStateAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchSetAlarmStateAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchSetAlarmStateAction_Override(c CloudWatchSetAlarmStateAction, alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		c,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (c *jsiiProxy_CloudWatchSetAlarmStateAction) Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for CloudWatch alarm.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchSetAlarmStateActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The value of the alarm state to set.
	// Experimental.
	AlarmStateToSet awscloudwatch.AlarmState `json:"alarmStateToSet"`
	// The reason for the alarm change.
	// Experimental.
	Reason *string `json:"reason"`
}

// Common properties shared by Actions it access to AWS service.
//
// TODO: EXAMPLE
//
// Experimental.
type CommonActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// The action to put the record from an MQTT message to the Kinesis Data Firehose stream.
//
// TODO: EXAMPLE
//
// Experimental.
type FirehoseStreamAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for FirehoseStreamAction
type jsiiProxy_FirehoseStreamAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewFirehoseStreamAction(stream awskinesisfirehose.IDeliveryStream, props *FirehoseStreamActionProps) FirehoseStreamAction {
	_init_.Initialize()

	j := jsiiProxy_FirehoseStreamAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.FirehoseStreamAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehoseStreamAction_Override(f FirehoseStreamAction, stream awskinesisfirehose.IDeliveryStream, props *FirehoseStreamActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.FirehoseStreamAction",
		[]interface{}{stream, props},
		f,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (f *jsiiProxy_FirehoseStreamAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for the Kinesis Data Firehose stream.
//
// TODO: EXAMPLE
//
// Experimental.
type FirehoseStreamActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using `PutRecordBatch`.
	//
	// When batchMode is true and the rule's SQL statement evaluates to an Array, each Array
	// element forms one record in the PutRecordBatch request. The resulting array can't have
	// more than 500 records.
	// Experimental.
	BatchMode *bool `json:"batchMode"`
	// A character separator that will be used to separate records written to the Kinesis Data Firehose stream.
	// Experimental.
	RecordSeparator FirehoseStreamRecordSeparator `json:"recordSeparator"`
}

// Record Separator to be used to separate records.
//
// TODO: EXAMPLE
//
// Experimental.
type FirehoseStreamRecordSeparator string

const (
	FirehoseStreamRecordSeparator_NEWLINE FirehoseStreamRecordSeparator = "NEWLINE"
	FirehoseStreamRecordSeparator_TAB FirehoseStreamRecordSeparator = "TAB"
	FirehoseStreamRecordSeparator_WINDOWS_NEWLINE FirehoseStreamRecordSeparator = "WINDOWS_NEWLINE"
	FirehoseStreamRecordSeparator_COMMA FirehoseStreamRecordSeparator = "COMMA"
)

// The action to invoke an AWS Lambda function, passing in an MQTT message.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaFunctionAction interface {
	awsiot.IAction
	Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for LambdaFunctionAction
type jsiiProxy_LambdaFunctionAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewLambdaFunctionAction(func_ awslambda.IFunction) LambdaFunctionAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionAction_Override(l LambdaFunctionAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		[]interface{}{func_},
		l,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (l *jsiiProxy_LambdaFunctionAction) Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

// The action to write the data from an MQTT message to an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
// Experimental.
type S3PutObjectAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for S3PutObjectAction
type jsiiProxy_S3PutObjectAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewS3PutObjectAction(bucket awss3.IBucket, props *S3PutObjectActionProps) S3PutObjectAction {
	_init_.Initialize()

	j := jsiiProxy_S3PutObjectAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.S3PutObjectAction",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3PutObjectAction_Override(s S3PutObjectAction, bucket awss3.IBucket, props *S3PutObjectActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.S3PutObjectAction",
		[]interface{}{bucket, props},
		s,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (s *jsiiProxy_S3PutObjectAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for s3.
//
// TODO: EXAMPLE
//
// Experimental.
type S3PutObjectActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Amazon S3 canned ACL that controls access to the object identified by the object key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	//
	// Experimental.
	AccessControl awss3.BucketAccessControl `json:"accessControl"`
	// The path to the file where the data is written.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	Key *string `json:"key"`
}

// The action to write the data from an MQTT message to an Amazon SQS queue.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsQueueAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for SqsQueueAction
type jsiiProxy_SqsQueueAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewSqsQueueAction(queue awssqs.IQueue, props *SqsQueueActionProps) SqsQueueAction {
	_init_.Initialize()

	j := jsiiProxy_SqsQueueAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.SqsQueueAction",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueueAction_Override(s SqsQueueAction, queue awssqs.IQueue, props *SqsQueueActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.SqsQueueAction",
		[]interface{}{queue, props},
		s,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (s *jsiiProxy_SqsQueueAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for SQS.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsQueueActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Specifies whether to use Base64 encoding.
	// Experimental.
	UseBase64 *bool `json:"useBase64"`
}

