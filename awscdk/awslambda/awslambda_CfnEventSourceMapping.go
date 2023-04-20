package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Lambda::EventSourceMapping`.
//
// The `AWS::Lambda::EventSourceMapping` resource creates a mapping between an event source and an AWS Lambda function. Lambda reads items from the event source and triggers the function.
//
// For details about each event source type, see the following topics. In particular, each of the topics describes the required and optional parameters for the specific event source.
//
// - [Configuring a Dynamo DB stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-dynamodb-eventsourcemapping)
// - [Configuring a Kinesis stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-eventsourcemapping)
// - [Configuring an SQS queue as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-eventsource)
// - [Configuring an MQ broker as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping)
// - [Configuring MSK as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)
// - [Configuring Self-Managed Apache Kafka as an event source](https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html)
// - [Configuring Amazon DocumentDB as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSourceMapping := awscdk.Aws_lambda.NewCfnEventSourceMapping(this, jsii.String("MyCfnEventSourceMapping"), &CfnEventSourceMappingProps{
//   	FunctionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	AmazonManagedKafkaEventSourceConfig: &AmazonManagedKafkaEventSourceConfigProperty{
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnFunctionError: jsii.Boolean(false),
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &OnFailureProperty{
//   			Destination: jsii.String("destination"),
//   		},
//   	},
//   	DocumentDbEventSourceConfig: &DocumentDBEventSourceConfigProperty{
//   		CollectionName: jsii.String("collectionName"),
//   		DatabaseName: jsii.String("databaseName"),
//   		FullDocument: jsii.String("fullDocument"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	EventSourceArn: jsii.String("eventSourceArn"),
//   	FilterCriteria: &FilterCriteriaProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	FunctionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	ParallelizationFactor: jsii.Number(123),
//   	Queues: []*string{
//   		jsii.String("queues"),
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		MaximumConcurrency: jsii.Number(123),
//   	},
//   	SelfManagedEventSource: &SelfManagedEventSourceProperty{
//   		Endpoints: &EndpointsProperty{
//   			KafkaBootstrapServers: []*string{
//   				jsii.String("kafkaBootstrapServers"),
//   			},
//   		},
//   	},
//   	SelfManagedKafkaEventSourceConfig: &SelfManagedKafkaEventSourceConfigProperty{
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	SourceAccessConfigurations: []interface{}{
//   		&SourceAccessConfigurationProperty{
//   			Type: jsii.String("type"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	StartingPositionTimestamp: jsii.Number(123),
//   	Topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	TumblingWindowInSeconds: jsii.Number(123),
//   })
//
type CfnEventSourceMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig() interface{}
	SetAmazonManagedKafkaEventSourceConfig(val interface{})
	// The event source mapping's ID.
	AttrId() *string
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//
	// - *Amazon Kinesis* – Default 100. Max 10,000.
	// - *Amazon DynamoDB Streams* – Default 100. Max 10,000.
	// - *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	// - *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.
	// - *Self-managed Apache Kafka* – Default 100. Max 10,000.
	// - *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.
	// - *DocumentDB* – Default 100. Max 10,000.
	BatchSize() *float64
	SetBatchSize(val *float64)
	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	BisectBatchOnFunctionError() interface{}
	SetBisectBatchOnFunctionError(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// (Kinesis and DynamoDB Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	DestinationConfig() interface{}
	SetDestinationConfig(val interface{})
	// Specific configuration settings for a DocumentDB event source.
	DocumentDbEventSourceConfig() interface{}
	SetDocumentDbEventSourceConfig(val interface{})
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//
	// Default: True.
	Enabled() interface{}
	SetEnabled(val interface{})
	// The Amazon Resource Name (ARN) of the event source.
	//
	// - *Amazon Kinesis* – The ARN of the data stream or a stream consumer.
	// - *Amazon DynamoDB Streams* – The ARN of the stream.
	// - *Amazon Simple Queue Service* – The ARN of the queue.
	// - *Amazon Managed Streaming for Apache Kafka* – The ARN of the cluster.
	// - *Amazon MQ* – The ARN of the broker.
	// - *Amazon DocumentDB* – The ARN of the DocumentDB change stream.
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	// An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) .
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* – `MyFunction` .
	// - *Function ARN* – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Version or Alias ARN* – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` .
	// - *Partial ARN* – `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	FunctionName() *string
	SetFunctionName(val *string)
	// (Streams and SQS) A list of current response type enums applied to the event source mapping.
	//
	// Valid Values: `ReportBatchItemFailures`.
	FunctionResponseTypes() *[]*string
	SetFunctionResponseTypes(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// *Default ( Kinesis , DynamoDB , Amazon SQS event sources)* : 0
	//
	// *Default ( Amazon MSK , Kafka, Amazon MQ , Amazon DocumentDB event sources)* : 500 ms
	//
	// *Related setting:* For Amazon SQS event sources, when you set `BatchSize` to a value greater than 10, you must set `MaximumBatchingWindowInSeconds` to at least 1.
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	// (Kinesis and DynamoDB Streams only) Discard records older than the specified age.
	//
	// The default value is -1,
	// which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	//
	// > The minimum value that can be set is 60 seconds.
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	// (Kinesis and DynamoDB Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1,
	// which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	ParallelizationFactor() *float64
	SetParallelizationFactor(val *float64)
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues() *[]*string
	SetQueues(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// ( Amazon Simple Queue Service only) The scaling configuration for the event source.
	//
	// For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency) .
	ScalingConfig() interface{}
	SetScalingConfig(val interface{})
	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource() interface{}
	SetSelfManagedEventSource(val interface{})
	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig() interface{}
	SetSelfManagedKafkaEventSourceConfig(val interface{})
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	SourceAccessConfigurations() interface{}
	SetSourceAccessConfigurations(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB.
	//
	// - *LATEST* - Read only new records.
	// - *TRIM_HORIZON* - Process all available records.
	// - *AT_TIMESTAMP* - Specify a time from which to start reading records.
	StartingPosition() *string
	SetStartingPosition(val *string)
	// With `StartingPosition` set to `AT_TIMESTAMP` , the time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp() *float64
	SetStartingPositionTimestamp(val *float64)
	// The name of the Kafka topic.
	Topics() *[]*string
	SetTopics(val *[]*string)
	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.
	//
	// A value of 0 seconds indicates no tumbling window.
	TumblingWindowInSeconds() *float64
	SetTumblingWindowInSeconds(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEventSourceMapping
type jsiiProxy_CfnEventSourceMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventSourceMapping) AmazonManagedKafkaEventSourceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) DestinationConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) DocumentDbEventSourceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentDbEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FilterCriteria() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) ScalingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) SelfManagedEventSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) SelfManagedKafkaEventSourceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) SourceAccessConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) StartingPositionTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping(scope awscdk.Construct, id *string, props *CfnEventSourceMappingProps) CfnEventSourceMapping {
	_init_.Initialize()

	if err := validateNewCfnEventSourceMappingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventSourceMapping{}

	_jsii_.Create(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping_Override(c CfnEventSourceMapping, scope awscdk.Construct, id *string, props *CfnEventSourceMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetAmazonManagedKafkaEventSourceConfig(val interface{}) {
	if err := j.validateSetAmazonManagedKafkaEventSourceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonManagedKafkaEventSourceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetBisectBatchOnFunctionError(val interface{}) {
	if err := j.validateSetBisectBatchOnFunctionErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetDestinationConfig(val interface{}) {
	if err := j.validateSetDestinationConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetDocumentDbEventSourceConfig(val interface{}) {
	if err := j.validateSetDocumentDbEventSourceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentDbEventSourceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetEventSourceArn(val *string) {
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetFilterCriteria(val interface{}) {
	if err := j.validateSetFilterCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCriteria",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetFunctionResponseTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetMaximumBatchingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetMaximumRecordAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetMaximumRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetParallelizationFactor(val *float64) {
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetQueues(val *[]*string) {
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetScalingConfig(val interface{}) {
	if err := j.validateSetScalingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetSelfManagedEventSource(val interface{}) {
	if err := j.validateSetSelfManagedEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedEventSource",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetSelfManagedKafkaEventSourceConfig(val interface{}) {
	if err := j.validateSetSelfManagedKafkaEventSourceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedKafkaEventSourceConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetSourceAccessConfigurations(val interface{}) {
	if err := j.validateSetSourceAccessConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAccessConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetStartingPosition(val *string) {
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetStartingPositionTimestamp(val *float64) {
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetTopics(val *[]*string) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping)SetTumblingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnEventSourceMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEventSourceMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSourceMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lambda.CfnEventSourceMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

