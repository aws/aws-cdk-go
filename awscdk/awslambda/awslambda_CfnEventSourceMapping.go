package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSourceMapping := awscdk.Aws_lambda.NewCfnEventSourceMapping(this, jsii.String("MyCfnEventSourceMapping"), &cfnEventSourceMappingProps{
//   	functionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	amazonManagedKafkaEventSourceConfig: &amazonManagedKafkaEventSourceConfigProperty{
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnFunctionError: jsii.Boolean(false),
//   	destinationConfig: &destinationConfigProperty{
//   		onFailure: &onFailureProperty{
//   			destination: jsii.String("destination"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	eventSourceArn: jsii.String("eventSourceArn"),
//   	filterCriteria: &filterCriteriaProperty{
//   		filters: []interface{}{
//   			&filterProperty{
//   				pattern: jsii.String("pattern"),
//   			},
//   		},
//   	},
//   	functionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	parallelizationFactor: jsii.Number(123),
//   	queues: []*string{
//   		jsii.String("queues"),
//   	},
//   	selfManagedEventSource: &selfManagedEventSourceProperty{
//   		endpoints: &endpointsProperty{
//   			kafkaBootstrapServers: []*string{
//   				jsii.String("kafkaBootstrapServers"),
//   			},
//   		},
//   	},
//   	selfManagedKafkaEventSourceConfig: &selfManagedKafkaEventSourceConfigProperty{
//   		consumerGroupId: jsii.String("consumerGroupId"),
//   	},
//   	sourceAccessConfigurations: []interface{}{
//   		&sourceAccessConfigurationProperty{
//   			type: jsii.String("type"),
//   			uri: jsii.String("uri"),
//   		},
//   	},
//   	startingPosition: jsii.String("startingPosition"),
//   	startingPositionTimestamp: jsii.Number(123),
//   	topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	tumblingWindowInSeconds: jsii.Number(123),
//   })
//
type CfnEventSourceMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Lambda::EventSourceMapping.AmazonManagedKafkaEventSourceConfig`.
	AmazonManagedKafkaEventSourceConfig() interface{}
	SetAmazonManagedKafkaEventSourceConfig(val interface{})
	// The event source mapping's ID.
	AttrId() *string
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//
	// - *Amazon Kinesis* - Default 100. Max 10,000.
	// - *Amazon DynamoDB Streams* - Default 100. Max 10,000.
	// - *Amazon Simple Queue Service* - Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	// - *Amazon Managed Streaming for Apache Kafka* - Default 100. Max 10,000.
	// - *Self-Managed Apache Kafka* - Default 100. Max 10,000.
	// - *Amazon MQ (ActiveMQ and RabbitMQ)* - Default 100. Max 10,000.
	BatchSize() *float64
	SetBatchSize(val *float64)
	// (Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	BisectBatchOnFunctionError() interface{}
	SetBisectBatchOnFunctionError(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// (Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	DestinationConfig() interface{}
	SetDestinationConfig(val interface{})
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//
	// Default: True.
	Enabled() interface{}
	SetEnabled(val interface{})
	// The Amazon Resource Name (ARN) of the event source.
	//
	// - *Amazon Kinesis* - The ARN of the data stream or a stream consumer.
	// - *Amazon DynamoDB Streams* - The ARN of the stream.
	// - *Amazon Simple Queue Service* - The ARN of the queue.
	// - *Amazon Managed Streaming for Apache Kafka* - The ARN of the cluster.
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	// (Streams and Amazon SQS) An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) .
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Version or Alias ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
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
	LogicalId() *string
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// *Default ( Kinesis , DynamoDB , Amazon SQS event sources)* : 0
	//
	// *Default ( Amazon MSK , Kafka, Amazon MQ event sources)* : 500 ms.
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	// (Streams only) Discard records older than the specified age.
	//
	// The default value is -1,
	// which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	// (Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1,
	// which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	// The tree node.
	Node() constructs.Node
	// (Streams only) The number of batches to process concurrently from each shard.
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
	Ref() *string
	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource() interface{}
	SetSelfManagedEventSource(val interface{})
	// `AWS::Lambda::EventSourceMapping.SelfManagedKafkaEventSourceConfig`.
	SelfManagedKafkaEventSourceConfig() interface{}
	SetSelfManagedKafkaEventSourceConfig(val interface{})
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	SourceAccessConfigurations() interface{}
	SetSourceAccessConfigurations(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
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
	// (Streams only) The duration in seconds of a processing window.
	//
	// The range is between 1 second up to 900 seconds.
	TumblingWindowInSeconds() *float64
	SetTumblingWindowInSeconds(val *float64)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnEventSourceMapping) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping(scope constructs.Construct, id *string, props *CfnEventSourceMappingProps) CfnEventSourceMapping {
	_init_.Initialize()

	if err := validateNewCfnEventSourceMappingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventSourceMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping_Override(c CfnEventSourceMapping, scope constructs.Construct, id *string, props *CfnEventSourceMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
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
func CfnEventSourceMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventSourceMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
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
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
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

