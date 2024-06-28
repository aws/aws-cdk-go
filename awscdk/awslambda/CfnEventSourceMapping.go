package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::EventSourceMapping` resource creates a mapping between an event source and an AWS Lambda function.
//
// Lambda reads items from the event source and triggers the function.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
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
	BatchSize() *float64
	SetBatchSize(val *float64)
	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.
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
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event sources only) A configuration object that specifies the destination of an event after Lambda processes it.
	DestinationConfig() interface{}
	SetDestinationConfig(val interface{})
	// Specific configuration settings for a DocumentDB event source.
	DocumentDbEventSourceConfig() interface{}
	SetDocumentDbEventSourceConfig(val interface{})
	// When true, the event source mapping is active.
	//
	// When false, Lambda pauses polling and invocation.
	Enabled() interface{}
	SetEnabled(val interface{})
	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	// An object that defines the filter criteria that determine whether Lambda should process an event.
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
	// The name or ARN of the Lambda function.
	FunctionName() *string
	SetFunctionName(val *string)
	// (Streams and SQS) A list of current response type enums applied to the event source mapping.
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
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	// (Kinesis and DynamoDB Streams only) Discard records older than the specified age.
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	// (Kinesis and DynamoDB Streams only) Discard records after the specified number of retries.
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	// The tree node.
	Node() constructs.Node
	// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard.
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
	// (Amazon SQS only) The scaling configuration for the event source.
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
	Stack() awscdk.Stack
	// The position in a stream from which to start reading.
	//
	// Required for Amazon Kinesis and Amazon DynamoDB.
	StartingPosition() *string
	SetStartingPosition(val *string)
	// With `StartingPosition` set to `AT_TIMESTAMP` , the time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp() *float64
	SetStartingPositionTimestamp(val *float64)
	// The name of the Kafka topic.
	Topics() *[]*string
	SetTopics(val *[]*string)
	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.
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
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

func (j *jsiiProxy_CfnEventSourceMapping) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


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

// Check whether the given object is a CfnResource.
func CfnEventSourceMapping_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMapping_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"isCfnResource",
		[]interface{}{x},
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

func (c *jsiiProxy_CfnEventSourceMapping) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnEventSourceMapping) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnEventSourceMapping) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnEventSourceMapping) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnEventSourceMapping) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

