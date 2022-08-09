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

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var model model
//   var requestValidator requestValidator
//
//   apiEventSource := awscdk.Aws_lambda_event_sources.NewApiEventSource(jsii.String("method"), jsii.String("path"), &methodOptions{
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   	authorizer: authorizer,
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			responseModels: map[string]iModel{
//   				"responseModelsKey": model,
//   			},
//   			responseParameters: map[string]*bool{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	operationName: jsii.String("operationName"),
//   	requestModels: map[string]*iModel{
//   		"requestModelsKey": model,
//   	},
//   	requestParameters: map[string]*bool{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	requestValidator: requestValidator,
//   	requestValidatorOptions: &requestValidatorOptions{
//   		requestValidatorName: jsii.String("requestValidatorName"),
//   		validateRequestBody: jsii.Boolean(false),
//   		validateRequestParameters: jsii.Boolean(false),
//   	},
//   })
//
type ApiEventSource interface {
	awslambda.IEventSource
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for ApiEventSource
type jsiiProxy_ApiEventSource struct {
	internal.Type__awslambdaIEventSource
}

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

func NewApiEventSource_Override(a ApiEventSource, method *string, path *string, options *awsapigateway.MethodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ApiEventSource",
		[]interface{}{method, path, options},
		a,
	)
}

func (a *jsiiProxy_ApiEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{target},
	)
}

// The authentication method to use with SelfManagedKafkaEventSource.
type AuthenticationMethod string

const (
	// SASL_SCRAM_512_AUTH authentication method for your Kafka cluster.
	AuthenticationMethod_SASL_SCRAM_512_AUTH AuthenticationMethod = "SASL_SCRAM_512_AUTH"
	// SASL_SCRAM_256_AUTH authentication method for your Kafka cluster.
	AuthenticationMethod_SASL_SCRAM_256_AUTH AuthenticationMethod = "SASL_SCRAM_256_AUTH"
	// BASIC_AUTH (SASL/PLAIN) authentication method for your Kafka cluster.
	AuthenticationMethod_BASIC_AUTH AuthenticationMethod = "BASIC_AUTH"
	// CLIENT_CERTIFICATE_TLS_AUTH (mTLS) authentication method for your Kafka cluster.
	AuthenticationMethod_CLIENT_CERTIFICATE_TLS_AUTH AuthenticationMethod = "CLIENT_CERTIFICATE_TLS_AUTH"
)

// The set of properties for streaming event sources shared by Dynamo, Kinesis and Kafka.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseStreamEventSourceProps := &baseStreamEventSourceProps{
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   	maxBatchingWindow: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type BaseStreamEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
}

// Use an Amazon DynamoDB stream as an event source for AWS Lambda.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//   var fn function
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewDynamoEventSource(table, &dynamoEventSourceProps{
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	batchSize: jsii.Number(5),
//   	bisectBatchOnError: jsii.Boolean(true),
//   	onFailure: awscdk.NewSqsDlq(deadLetterQueue),
//   	retryAttempts: jsii.Number(10),
//   }))
//
type DynamoEventSource interface {
	StreamEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewDynamoEventSource_Override(d DynamoEventSource, table awsdynamodb.ITable, props *DynamoEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		[]interface{}{table, props},
		d,
	)
}

func (d *jsiiProxy_DynamoEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		d,
		"bind",
		[]interface{}{target},
	)
}

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

// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//   var fn function
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewDynamoEventSource(table, &dynamoEventSourceProps{
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	batchSize: jsii.Number(5),
//   	bisectBatchOnError: jsii.Boolean(true),
//   	onFailure: awscdk.NewSqsDlq(deadLetterQueue),
//   	retryAttempts: jsii.Number(10),
//   }))
//
type DynamoEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
}

// Properties for a Kafka event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   kafkaEventSourceProps := &kafkaEventSourceProps{
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   	maxBatchingWindow: cdk.duration.minutes(jsii.Number(30)),
//   	secret: secret,
//   }
//
type KafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
}

// Use an Amazon Kinesis stream as an event source for AWS Lambda.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//   myFunction.addEventSource(awscdk.NewKinesisEventSource(stream, &kinesisEventSourceProps{
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
type KinesisEventSource interface {
	StreamEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	Stream() awskinesis.IStream
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewKinesisEventSource_Override(k KinesisEventSource, stream awskinesis.IStream, props *KinesisEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisEventSource",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		k,
		"bind",
		[]interface{}{target},
	)
}

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

// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//   myFunction.addEventSource(awscdk.NewKinesisEventSource(stream, &kinesisEventSourceProps{
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
type KinesisEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
	// The time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp *float64 `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
}

// Use a MSK cluster as a streaming source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // The secret that allows access to your MSK cluster
//   // You still have to make sure that it is associated with your cluster as described in the documentation
//   secret := awscdk.NewSecret(this, jsii.String("Secret"), &secretProps{
//   	secretName: jsii.String("AmazonMSK_KafkaSecret"),
//   })
//   myFunction.addEventSource(awscdk.NewManagedKafkaEventSource(&managedKafkaEventSourceProps{
//   	clusterArn: jsii.String(clusterArn),
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
type ManagedKafkaEventSource interface {
	StreamEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewManagedKafkaEventSource_Override(m ManagedKafkaEventSource, props *ManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_ManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		m,
		"bind",
		[]interface{}{target},
	)
}

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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // The secret that allows access to your MSK cluster
//   // You still have to make sure that it is associated with your cluster as described in the documentation
//   secret := awscdk.NewSecret(this, jsii.String("Secret"), &secretProps{
//   	secretName: jsii.String("AmazonMSK_KafkaSecret"),
//   })
//   myFunction.addEventSource(awscdk.NewManagedKafkaEventSource(&managedKafkaEventSourceProps{
//   	clusterArn: jsii.String(clusterArn),
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
type ManagedKafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// An MSK cluster construct.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
}

// Use S3 bucket notifications as an event source for AWS Lambda.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   fn.addEventSource(eventsources.NewS3EventSource(bucket, &s3EventSourceProps{
//   	events: []eventType{
//   		s3.*eventType_OBJECT_CREATED,
//   		s3.*eventType_OBJECT_REMOVED,
//   	},
//   	filters: []notificationKeyFilter{
//   		&notificationKeyFilter{
//   			prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
type S3EventSource interface {
	awslambda.IEventSource
	Bucket() awss3.Bucket
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewS3EventSource_Override(s S3EventSource, bucket awss3.Bucket, props *S3EventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.S3EventSource",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3EventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   fn.addEventSource(eventsources.NewS3EventSource(bucket, &s3EventSourceProps{
//   	events: []eventType{
//   		s3.*eventType_OBJECT_CREATED,
//   		s3.*eventType_OBJECT_REMOVED,
//   	},
//   	filters: []notificationKeyFilter{
//   		&notificationKeyFilter{
//   			prefix: jsii.String("subdir/"),
//   		},
//   	},
//   }))
//
type S3EventSourceProps struct {
	// The s3 event types that will trigger the notification.
	Events *[]awss3.EventType `field:"required" json:"events" yaml:"events"`
	// S3 object key filter rules to determine which objects trigger this event.
	//
	// Each filter must include a `prefix` and/or `suffix` that will be matched
	// against the s3 object key. Refer to the S3 Developer Guide for details
	// about allowed filter rules.
	Filters *[]*awss3.NotificationKeyFilter `field:"optional" json:"filters" yaml:"filters"`
}

// Use a self hosted Kafka installation as a streaming source for AWS Lambda.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
//
//   // (Optional) The secret containing the root CA certificate that your Kafka brokers use for TLS encryption
//   var encryption secret
//
//   var myFunction function
//
//
//   // The list of Kafka brokers
//   bootstrapServers := []*string{
//   	"kafka-broker:9092",
//   }
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//   myFunction.addEventSource(awscdk.NewSelfManagedKafkaEventSource(&selfManagedKafkaEventSourceProps{
//   	bootstrapServers: bootstrapServers,
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	encryption: encryption,
//   }))
//
type SelfManagedKafkaEventSource interface {
	StreamEventSource
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewSelfManagedKafkaEventSource_Override(s SelfManagedKafkaEventSource, props *SelfManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SelfManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

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
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
//
//   // (Optional) The secret containing the root CA certificate that your Kafka brokers use for TLS encryption
//   var encryption secret
//
//   var myFunction function
//
//
//   // The list of Kafka brokers
//   bootstrapServers := []*string{
//   	"kafka-broker:9092",
//   }
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//   myFunction.addEventSource(awscdk.NewSelfManagedKafkaEventSource(&selfManagedKafkaEventSourceProps{
//   	bootstrapServers: bootstrapServers,
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	encryption: encryption,
//   }))
//
type SelfManagedKafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The list of host and port pairs that are the addresses of the Kafka brokers in a "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.xyz.com:xxxx`.
	BootstrapServers *[]*string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// The authentication method for your Kafka cluster.
	AuthenticationMethod AuthenticationMethod `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// The secret with the root CA certificate used by your Kafka brokers for TLS encryption This field is required if your Kafka brokers use certificates signed by a private CA.
	RootCACertificate awssecretsmanager.Secret `field:"optional" json:"rootCACertificate" yaml:"rootCACertificate"`
	// If your Kafka brokers are only reachable via VPC, provide the security group here.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// If your Kafka brokers are only reachable via VPC provide the VPC here.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// If your Kafka brokers are only reachable via VPC, provide the subnets selection here.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// An SNS dead letter queue destination configuration for a Lambda event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   snsDlq := awscdk.Aws_lambda_event_sources.NewSnsDlq(topic)
//
type SnsDlq interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SnsDlq
type jsiiProxy_SnsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

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

func NewSnsDlq_Override(s SnsDlq, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		[]interface{}{topic},
		s,
	)
}

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
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   var fn function
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewSnsEventSource(topic, &snsEventSourceProps{
//   	filterPolicy: map[string]interface{}{
//   	},
//   	deadLetterQueue: deadLetterQueue,
//   }))
//
type SnsEventSource interface {
	awslambda.IEventSource
	Topic() awssns.ITopic
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewSnsEventSource_Override(s SnsEventSource, topic awssns.ITopic, props *SnsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// Properties forwarded to the Lambda Subscription.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   var fn function
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewSnsEventSource(topic, &snsEventSourceProps{
//   	filterPolicy: map[string]interface{}{
//   	},
//   	deadLetterQueue: deadLetterQueue,
//   }))
//
type SnsEventSourceProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

// An SQS dead letter queue destination configuration for a Lambda event source.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//   var fn function
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewDynamoEventSource(table, &dynamoEventSourceProps{
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	batchSize: jsii.Number(5),
//   	bisectBatchOnError: jsii.Boolean(true),
//   	onFailure: awscdk.NewSqsDlq(deadLetterQueue),
//   	retryAttempts: jsii.Number(10),
//   }))
//
type SqsDlq interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SqsDlq
type jsiiProxy_SqsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

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

func NewSqsDlq_Override(s SqsDlq, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		[]interface{}{queue},
		s,
	)
}

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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//   eventSource := awscdk.NewSqsEventSource(queue)
//   fn.addEventSource(eventSource)
//
//   eventSourceId := eventSource.eventSourceMappingId
//
type SqsEventSource interface {
	awslambda.IEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Queue() awssqs.IQueue
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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

func NewSqsEventSource_Override(s SqsEventSource, queue awssqs.IQueue, props *SqsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"), &queueProps{
//   	visibilityTimeout: awscdk.Duration.seconds(jsii.Number(30)),
//   	 // default,
//   	receiveMessageWaitTime: awscdk.Duration.seconds(jsii.Number(20)),
//   })
//
//   fn.addEventSource(awscdk.NewSqsEventSource(queue, &sqsEventSourceProps{
//   	batchSize: jsii.Number(10),
//   	 // default
//   	maxBatchingWindow: awscdk.Duration.minutes(jsii.Number(5)),
//   	reportBatchItemFailures: jsii.Boolean(true),
//   }))
//
type SqsEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10.
	// If `maxBatchingWindow` is configured, this value can go up to 10,000.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the SQS event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Valid Range: Minimum value of 0 minutes. Maximum value of 5 minutes.
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-batchfailurereporting
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
}

// Use an stream as an event source for AWS Lambda.
type StreamEventSource interface {
	awslambda.IEventSource
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
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


func NewStreamEventSource_Override(s StreamEventSource, props *StreamEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.StreamEventSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StreamEventSource) Bind(_target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{_target},
	)
}

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

// The set of properties for streaming event sources shared by Dynamo and Kinesis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventSourceDlq iEventSourceDlq
//
//   streamEventSourceProps := &streamEventSourceProps{
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnError: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	maxBatchingWindow: cdk.duration.minutes(jsii.Number(30)),
//   	maxRecordAge: cdk.*duration.minutes(jsii.Number(30)),
//   	onFailure: eventSourceDlq,
//   	parallelizationFactor: jsii.Number(123),
//   	reportBatchItemFailures: jsii.Boolean(false),
//   	retryAttempts: jsii.Number(123),
//   	tumblingWindow: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type StreamEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `field:"optional" json:"bisectBatchOnError" yaml:"bisectBatchOnError"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days.
	MaxRecordAge awscdk.Duration `field:"optional" json:"maxRecordAge" yaml:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
	// Maximum number of retry attempts Valid Range: * Minimum value of 0 * Maximum value of 10000.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis Valid Range: 0 - 15 minutes.
	TumblingWindow awscdk.Duration `field:"optional" json:"tumblingWindow" yaml:"tumblingWindow"`
}

