package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// TODO: EXAMPLE
//
// Experimental.
type ApiEventSource interface {
	awslambda.IEventSource
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for ApiEventSource
type jsiiProxy_ApiEventSource struct {
	internal.Type__awslambdaIEventSource
}

// Experimental.
func NewApiEventSource(method *string, path *string, options *awsapigateway.MethodOptions) ApiEventSource {
	_init_.Initialize()

	j := jsiiProxy_ApiEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewApiEventSource_Override(a ApiEventSource, method *string, path *string, options *awsapigateway.MethodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		a,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (a *jsiiProxy_ApiEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{target},
	)
}

// The authentication method to use with SelfManagedKafkaEventSource.
// Experimental.
type AuthenticationMethod string

const (
	AuthenticationMethod_SASL_SCRAM_512_AUTH AuthenticationMethod = "SASL_SCRAM_512_AUTH"
	AuthenticationMethod_SASL_SCRAM_256_AUTH AuthenticationMethod = "SASL_SCRAM_256_AUTH"
	AuthenticationMethod_BASIC_AUTH AuthenticationMethod = "BASIC_AUTH"
)

// Use an Amazon DynamoDB stream as an event source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type DynamoEventSource interface {
	StreamEventSource
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for DynamoEventSource
type jsiiProxy_DynamoEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_DynamoEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoEventSource(table awsdynamodb.ITable, props *DynamoEventSourceProps) DynamoEventSource {
	_init_.Initialize()

	j := jsiiProxy_DynamoEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		[]interface{}{table, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoEventSource_Override(d DynamoEventSource, table awsdynamodb.ITable, props *DynamoEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		[]interface{}{table, props},
		d,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (d *jsiiProxy_DynamoEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		d,
		"bind",
		[]interface{}{target},
	)
}

// Experimental.
func (d *jsiiProxy_DynamoEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		d,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
// Experimental.
type DynamoEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
}

// Properties for a Kafka event source.
//
// TODO: EXAMPLE
//
// Experimental.
type KafkaEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret"`
	// The Kafka topic to subscribe to.
	// Experimental.
	Topic *string `json:"topic"`
}

// Use an Amazon Kinesis stream as an event source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type KinesisEventSource interface {
	StreamEventSource
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	Stream() awskinesis.IStream
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for KinesisEventSource
type jsiiProxy_KinesisEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_KinesisEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisEventSource) Stream() awskinesis.IStream {
	var returns awskinesis.IStream
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisEventSource(stream awskinesis.IStream, props *KinesisEventSourceProps) KinesisEventSource {
	_init_.Initialize()

	j := jsiiProxy_KinesisEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisEventSource",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisEventSource_Override(k KinesisEventSource, stream awskinesis.IStream, props *KinesisEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisEventSource",
		[]interface{}{stream, props},
		k,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (k *jsiiProxy_KinesisEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		k,
		"bind",
		[]interface{}{target},
	)
}

// Experimental.
func (k *jsiiProxy_KinesisEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		k,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
// Experimental.
type KinesisEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
}

// Use a MSK cluster as a streaming source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type ManagedKafkaEventSource interface {
	StreamEventSource
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for ManagedKafkaEventSource
type jsiiProxy_ManagedKafkaEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_ManagedKafkaEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewManagedKafkaEventSource(props *ManagedKafkaEventSourceProps) ManagedKafkaEventSource {
	_init_.Initialize()

	j := jsiiProxy_ManagedKafkaEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewManagedKafkaEventSource_Override(m ManagedKafkaEventSource, props *ManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		[]interface{}{props},
		m,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (m *jsiiProxy_ManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		m,
		"bind",
		[]interface{}{target},
	)
}

// Experimental.
func (m *jsiiProxy_ManagedKafkaEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		m,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties for a MSK event source.
//
// TODO: EXAMPLE
//
// Experimental.
type ManagedKafkaEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret"`
	// The Kafka topic to subscribe to.
	// Experimental.
	Topic *string `json:"topic"`
	// An MSK cluster construct.
	// Experimental.
	ClusterArn *string `json:"clusterArn"`
}

// Use S3 bucket notifications as an event source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type S3EventSource interface {
	awslambda.IEventSource
	Bucket() awss3.Bucket
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for S3EventSource
type jsiiProxy_S3EventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_S3EventSource) Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3EventSource(bucket awss3.Bucket, props *S3EventSourceProps) S3EventSource {
	_init_.Initialize()

	j := jsiiProxy_S3EventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3EventSource_Override(s S3EventSource, bucket awss3.Bucket, props *S3EventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		[]interface{}{bucket, props},
		s,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (s *jsiiProxy_S3EventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// TODO: EXAMPLE
//
// Experimental.
type S3EventSourceProps struct {
	// The s3 event types that will trigger the notification.
	// Experimental.
	Events *[]awss3.EventType `json:"events"`
	// S3 object key filter rules to determine which objects trigger this event.
	//
	// Each filter must include a `prefix` and/or `suffix` that will be matched
	// against the s3 object key. Refer to the S3 Developer Guide for details
	// about allowed filter rules.
	// Experimental.
	Filters *[]*awss3.NotificationKeyFilter `json:"filters"`
}

// Use a self hosted Kafka installation as a streaming source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type SelfManagedKafkaEventSource interface {
	StreamEventSource
	Props() *StreamEventSourceProps
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for SelfManagedKafkaEventSource
type jsiiProxy_SelfManagedKafkaEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_SelfManagedKafkaEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewSelfManagedKafkaEventSource(props *SelfManagedKafkaEventSourceProps) SelfManagedKafkaEventSource {
	_init_.Initialize()

	j := jsiiProxy_SelfManagedKafkaEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSelfManagedKafkaEventSource_Override(s SelfManagedKafkaEventSource, props *SelfManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		[]interface{}{props},
		s,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (s *jsiiProxy_SelfManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// Experimental.
func (s *jsiiProxy_SelfManagedKafkaEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		s,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Properties for a self managed Kafka cluster event source.
//
// If your Kafka cluster is only reachable via VPC make sure to configure it.
//
// TODO: EXAMPLE
//
// Experimental.
type SelfManagedKafkaEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret"`
	// The Kafka topic to subscribe to.
	// Experimental.
	Topic *string `json:"topic"`
	// The list of host and port pairs that are the addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.xyz.com:xxxx`.
	// Experimental.
	BootstrapServers *[]*string `json:"bootstrapServers"`
	// The authentication method for your Kafka cluster.
	// Experimental.
	AuthenticationMethod AuthenticationMethod `json:"authenticationMethod"`
	// If your Kafka brokers are only reachable via VPC, provide the security group here.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// If your Kafka brokers are only reachable via VPC provide the VPC here.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// If your Kafka brokers are only reachable via VPC, provide the subnets selection here.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// An SNS dead letter queue destination configuration for a Lambda event source.
//
// TODO: EXAMPLE
//
// Experimental.
type SnsDlq interface {
	awslambda.IEventSourceDlq
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SnsDlq
type jsiiProxy_SnsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

// Experimental.
func NewSnsDlq(topic awssns.ITopic) SnsDlq {
	_init_.Initialize()

	j := jsiiProxy_SnsDlq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsDlq_Override(s SnsDlq, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		[]interface{}{topic},
		s,
	)
}

// Returns a destination configuration for the DLQ.
// Experimental.
func (s *jsiiProxy_SnsDlq) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

// Use an Amazon SNS topic as an event source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type SnsEventSource interface {
	awslambda.IEventSource
	Topic() awssns.ITopic
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for SnsEventSource
type jsiiProxy_SnsEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_SnsEventSource) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsEventSource(topic awssns.ITopic, props *SnsEventSourceProps) SnsEventSource {
	_init_.Initialize()

	j := jsiiProxy_SnsEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsEventSource_Override(s SnsEventSource, topic awssns.ITopic, props *SnsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		[]interface{}{topic, props},
		s,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (s *jsiiProxy_SnsEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// Properties forwarded to the Lambda Subscription.
//
// TODO: EXAMPLE
//
// Experimental.
type SnsEventSourceProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `json:"filterPolicy"`
}

// An SQS dead letter queue destination configuration for a Lambda event source.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsDlq interface {
	awslambda.IEventSourceDlq
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SqsDlq
type jsiiProxy_SqsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

// Experimental.
func NewSqsDlq(queue awssqs.IQueue) SqsDlq {
	_init_.Initialize()

	j := jsiiProxy_SqsDlq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsDlq_Override(s SqsDlq, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		[]interface{}{queue},
		s,
	)
}

// Returns a destination configuration for the DLQ.
// Experimental.
func (s *jsiiProxy_SqsDlq) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

// Use an Amazon SQS queue as an event source for AWS Lambda.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsEventSource interface {
	awslambda.IEventSource
	EventSourceMappingId() *string
	Queue() awssqs.IQueue
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for SqsEventSource
type jsiiProxy_SqsEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_SqsEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsEventSource) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsEventSource(queue awssqs.IQueue, props *SqsEventSourceProps) SqsEventSource {
	_init_.Initialize()

	j := jsiiProxy_SqsEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsEventSource_Override(s SqsEventSource, queue awssqs.IQueue, props *SqsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		[]interface{}{queue, props},
		s,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (s *jsiiProxy_SqsEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// TODO: EXAMPLE
//
// Experimental.
type SqsEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10.
	// If `maxBatchingWindow` is configured, this value can go up to 10,000.
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the SQS event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Valid Range: Minimum value of 0 minutes. Maximum value of 5 minutes.
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
}

// Use an stream as an event source for AWS Lambda.
// Experimental.
type StreamEventSource interface {
	awslambda.IEventSource
	Props() *StreamEventSourceProps
	Bind(_target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for StreamEventSource
type jsiiProxy_StreamEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_StreamEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewStreamEventSource_Override(s StreamEventSource, props *StreamEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.StreamEventSource",
		[]interface{}{props},
		s,
	)
}

// Called by `lambda.addEventSource` to allow the event source to bind to this function.
// Experimental.
func (s *jsiiProxy_StreamEventSource) Bind(_target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{_target},
	)
}

// Experimental.
func (s *jsiiProxy_StreamEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		s,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// The set of properties for event sources that follow the streaming model, such as, Dynamo, Kinesis and Kafka.
//
// TODO: EXAMPLE
//
// Experimental.
type StreamEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	//    * 10000 for {@link KinesisEventSource}
	// Experimental.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	// Experimental.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// If the stream event source mapping should be enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	// Experimental.
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	// Experimental.
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	// Experimental.
	OnFailure awslambda.IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	// Experimental.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	// Experimental.
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Where to begin consuming the stream.
	// Experimental.
	StartingPosition awslambda.StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	// Experimental.
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
}

