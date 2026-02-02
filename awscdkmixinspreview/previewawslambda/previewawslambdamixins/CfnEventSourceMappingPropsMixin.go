package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventSourceMappingPropsMixin := awscdkmixinspreview.Mixins.NewCfnEventSourceMappingPropsMixin(&CfnEventSourceMappingMixinProps{
//   	AmazonManagedKafkaEventSourceConfig: &AmazonManagedKafkaEventSourceConfigProperty{
//   		ConsumerGroupId: jsii.String("consumerGroupId"),
//   		SchemaRegistryConfig: &SchemaRegistryConfigProperty{
//   			AccessConfigs: []interface{}{
//   				&SchemaRegistryAccessConfigProperty{
//   					Type: jsii.String("type"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			EventRecordFormat: jsii.String("eventRecordFormat"),
//   			SchemaRegistryUri: jsii.String("schemaRegistryUri"),
//   			SchemaValidationConfigs: []interface{}{
//   				&SchemaValidationConfigProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   			},
//   		},
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
//   	FunctionName: jsii.String("functionName"),
//   	FunctionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LoggingConfig: &LoggingConfigProperty{
//   		SystemLogLevel: jsii.String("systemLogLevel"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	MetricsConfig: &MetricsConfigProperty{
//   		Metrics: []*string{
//   			jsii.String("metrics"),
//   		},
//   	},
//   	ParallelizationFactor: jsii.Number(123),
//   	ProvisionedPollerConfig: &ProvisionedPollerConfigProperty{
//   		MaximumPollers: jsii.Number(123),
//   		MinimumPollers: jsii.Number(123),
//   		PollerGroupName: jsii.String("pollerGroupName"),
//   	},
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
//   		SchemaRegistryConfig: &SchemaRegistryConfigProperty{
//   			AccessConfigs: []interface{}{
//   				&SchemaRegistryAccessConfigProperty{
//   					Type: jsii.String("type"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			EventRecordFormat: jsii.String("eventRecordFormat"),
//   			SchemaRegistryUri: jsii.String("schemaRegistryUri"),
//   			SchemaValidationConfigs: []interface{}{
//   				&SchemaValidationConfigProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   			},
//   		},
//   	},
//   	SourceAccessConfigurations: []interface{}{
//   		&SourceAccessConfigurationProperty{
//   			Type: jsii.String("type"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	StartingPositionTimestamp: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Topics: []*string{
//   		jsii.String("topics"),
//   	},
//   	TumblingWindowInSeconds: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
//
type CfnEventSourceMappingPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEventSourceMappingMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventSourceMappingPropsMixin
type jsiiProxy_CfnEventSourceMappingPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEventSourceMappingPropsMixin) Props() *CfnEventSourceMappingMixinProps {
	var returns *CfnEventSourceMappingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMappingPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMappingPropsMixin(props *CfnEventSourceMappingMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEventSourceMappingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventSourceMappingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventSourceMappingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMappingPropsMixin_Override(c CfnEventSourceMappingPropsMixin, props *CfnEventSourceMappingMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEventSourceMappingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventSourceMappingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSourceMappingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnEventSourceMappingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventSourceMappingPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEventSourceMappingPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

