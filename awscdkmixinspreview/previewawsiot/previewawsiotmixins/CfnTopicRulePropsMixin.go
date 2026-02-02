package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::TopicRule` resource to declare an AWS IoT rule.
//
// For information about working with AWS IoT rules, see [Rules for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTopicRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnTopicRulePropsMixin(&CfnTopicRuleMixinProps{
//   	RuleName: jsii.String("ruleName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicRulePayload: &TopicRulePayloadProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   					AlarmName: jsii.String("alarmName"),
//   					RoleArn: jsii.String("roleArn"),
//   					StateReason: jsii.String("stateReason"),
//   					StateValue: jsii.String("stateValue"),
//   				},
//   				CloudwatchLogs: &CloudwatchLogsActionProperty{
//   					BatchMode: jsii.Boolean(false),
//   					LogGroupName: jsii.String("logGroupName"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				CloudwatchMetric: &CloudwatchMetricActionProperty{
//   					MetricName: jsii.String("metricName"),
//   					MetricNamespace: jsii.String("metricNamespace"),
//   					MetricTimestamp: jsii.String("metricTimestamp"),
//   					MetricUnit: jsii.String("metricUnit"),
//   					MetricValue: jsii.String("metricValue"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				DynamoDb: &DynamoDBActionProperty{
//   					HashKeyField: jsii.String("hashKeyField"),
//   					HashKeyType: jsii.String("hashKeyType"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   					RoleArn: jsii.String("roleArn"),
//   					TableName: jsii.String("tableName"),
//   				},
//   				DynamoDBv2: &DynamoDBv2ActionProperty{
//   					PutItem: &PutItemInputProperty{
//   						TableName: jsii.String("tableName"),
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Elasticsearch: &ElasticsearchActionProperty{
//   					Endpoint: jsii.String("endpoint"),
//   					Id: jsii.String("id"),
//   					Index: jsii.String("index"),
//   					RoleArn: jsii.String("roleArn"),
//   					Type: jsii.String("type"),
//   				},
//   				Firehose: &FirehoseActionProperty{
//   					BatchMode: jsii.Boolean(false),
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//   					RoleArn: jsii.String("roleArn"),
//   					Separator: jsii.String("separator"),
//   				},
//   				Http: &HttpActionProperty{
//   					Auth: &HttpAuthorizationProperty{
//   						Sigv4: &SigV4AuthorizationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							ServiceName: jsii.String("serviceName"),
//   							SigningRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					BatchConfig: &BatchConfigProperty{
//   						MaxBatchOpenMs: jsii.Number(123),
//   						MaxBatchSize: jsii.Number(123),
//   						MaxBatchSizeBytes: jsii.Number(123),
//   					},
//   					ConfirmationUrl: jsii.String("confirmationUrl"),
//   					EnableBatching: jsii.Boolean(false),
//   					Headers: []interface{}{
//   						&HttpActionHeaderProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Url: jsii.String("url"),
//   				},
//   				IotAnalytics: &IotAnalyticsActionProperty{
//   					BatchMode: jsii.Boolean(false),
//   					ChannelName: jsii.String("channelName"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				IotEvents: &IotEventsActionProperty{
//   					BatchMode: jsii.Boolean(false),
//   					InputName: jsii.String("inputName"),
//   					MessageId: jsii.String("messageId"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				IotSiteWise: &IotSiteWiseActionProperty{
//   					PutAssetPropertyValueEntries: []interface{}{
//   						&PutAssetPropertyValueEntryProperty{
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   							PropertyValues: []interface{}{
//   								&AssetPropertyValueProperty{
//   									Quality: jsii.String("quality"),
//   									Timestamp: &AssetPropertyTimestampProperty{
//   										OffsetInNanos: jsii.String("offsetInNanos"),
//   										TimeInSeconds: jsii.String("timeInSeconds"),
//   									},
//   									Value: &AssetPropertyVariantProperty{
//   										BooleanValue: jsii.String("booleanValue"),
//   										DoubleValue: jsii.String("doubleValue"),
//   										IntegerValue: jsii.String("integerValue"),
//   										StringValue: jsii.String("stringValue"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Kafka: &KafkaActionProperty{
//   					ClientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					DestinationArn: jsii.String("destinationArn"),
//   					Headers: []interface{}{
//   						&KafkaActionHeaderProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Key: jsii.String("key"),
//   					Partition: jsii.String("partition"),
//   					Topic: jsii.String("topic"),
//   				},
//   				Kinesis: &KinesisActionProperty{
//   					PartitionKey: jsii.String("partitionKey"),
//   					RoleArn: jsii.String("roleArn"),
//   					StreamName: jsii.String("streamName"),
//   				},
//   				Lambda: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   				Location: &LocationActionProperty{
//   					DeviceId: jsii.String("deviceId"),
//   					Latitude: jsii.String("latitude"),
//   					Longitude: jsii.String("longitude"),
//   					RoleArn: jsii.String("roleArn"),
//   					Timestamp: &TimestampProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.String("value"),
//   					},
//   					TrackerName: jsii.String("trackerName"),
//   				},
//   				OpenSearch: &OpenSearchActionProperty{
//   					Endpoint: jsii.String("endpoint"),
//   					Id: jsii.String("id"),
//   					Index: jsii.String("index"),
//   					RoleArn: jsii.String("roleArn"),
//   					Type: jsii.String("type"),
//   				},
//   				Republish: &RepublishActionProperty{
//   					Headers: &RepublishActionHeadersProperty{
//   						ContentType: jsii.String("contentType"),
//   						CorrelationData: jsii.String("correlationData"),
//   						MessageExpiry: jsii.String("messageExpiry"),
//   						PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   						ResponseTopic: jsii.String("responseTopic"),
//   						UserProperties: []interface{}{
//   							&UserPropertyProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Qos: jsii.Number(123),
//   					RoleArn: jsii.String("roleArn"),
//   					Topic: jsii.String("topic"),
//   				},
//   				S3: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//   					CannedAcl: jsii.String("cannedAcl"),
//   					Key: jsii.String("key"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Sns: &SnsActionProperty{
//   					MessageFormat: jsii.String("messageFormat"),
//   					RoleArn: jsii.String("roleArn"),
//   					TargetArn: jsii.String("targetArn"),
//   				},
//   				Sqs: &SqsActionProperty{
//   					QueueUrl: jsii.String("queueUrl"),
//   					RoleArn: jsii.String("roleArn"),
//   					UseBase64: jsii.Boolean(false),
//   				},
//   				StepFunctions: &StepFunctionsActionProperty{
//   					ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   					RoleArn: jsii.String("roleArn"),
//   					StateMachineName: jsii.String("stateMachineName"),
//   				},
//   				Timestream: &TimestreamActionProperty{
//   					DatabaseName: jsii.String("databaseName"),
//   					Dimensions: []interface{}{
//   						&TimestreamDimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   					TableName: jsii.String("tableName"),
//   					Timestamp: &TimestreamTimestampProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		AwsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   		Description: jsii.String("description"),
//   		ErrorAction: &ActionProperty{
//   			CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   				AlarmName: jsii.String("alarmName"),
//   				RoleArn: jsii.String("roleArn"),
//   				StateReason: jsii.String("stateReason"),
//   				StateValue: jsii.String("stateValue"),
//   			},
//   			CloudwatchLogs: &CloudwatchLogsActionProperty{
//   				BatchMode: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			CloudwatchMetric: &CloudwatchMetricActionProperty{
//   				MetricName: jsii.String("metricName"),
//   				MetricNamespace: jsii.String("metricNamespace"),
//   				MetricTimestamp: jsii.String("metricTimestamp"),
//   				MetricUnit: jsii.String("metricUnit"),
//   				MetricValue: jsii.String("metricValue"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			DynamoDb: &DynamoDBActionProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyType: jsii.String("hashKeyType"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//   			},
//   			DynamoDBv2: &DynamoDBv2ActionProperty{
//   				PutItem: &PutItemInputProperty{
//   					TableName: jsii.String("tableName"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Elasticsearch: &ElasticsearchActionProperty{
//   				Endpoint: jsii.String("endpoint"),
//   				Id: jsii.String("id"),
//   				Index: jsii.String("index"),
//   				RoleArn: jsii.String("roleArn"),
//   				Type: jsii.String("type"),
//   			},
//   			Firehose: &FirehoseActionProperty{
//   				BatchMode: jsii.Boolean(false),
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//   				RoleArn: jsii.String("roleArn"),
//   				Separator: jsii.String("separator"),
//   			},
//   			Http: &HttpActionProperty{
//   				Auth: &HttpAuthorizationProperty{
//   					Sigv4: &SigV4AuthorizationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						ServiceName: jsii.String("serviceName"),
//   						SigningRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				BatchConfig: &BatchConfigProperty{
//   					MaxBatchOpenMs: jsii.Number(123),
//   					MaxBatchSize: jsii.Number(123),
//   					MaxBatchSizeBytes: jsii.Number(123),
//   				},
//   				ConfirmationUrl: jsii.String("confirmationUrl"),
//   				EnableBatching: jsii.Boolean(false),
//   				Headers: []interface{}{
//   					&HttpActionHeaderProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Url: jsii.String("url"),
//   			},
//   			IotAnalytics: &IotAnalyticsActionProperty{
//   				BatchMode: jsii.Boolean(false),
//   				ChannelName: jsii.String("channelName"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			IotEvents: &IotEventsActionProperty{
//   				BatchMode: jsii.Boolean(false),
//   				InputName: jsii.String("inputName"),
//   				MessageId: jsii.String("messageId"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			IotSiteWise: &IotSiteWiseActionProperty{
//   				PutAssetPropertyValueEntries: []interface{}{
//   					&PutAssetPropertyValueEntryProperty{
//   						AssetId: jsii.String("assetId"),
//   						EntryId: jsii.String("entryId"),
//   						PropertyAlias: jsii.String("propertyAlias"),
//   						PropertyId: jsii.String("propertyId"),
//   						PropertyValues: []interface{}{
//   							&AssetPropertyValueProperty{
//   								Quality: jsii.String("quality"),
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Kafka: &KafkaActionProperty{
//   				ClientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				DestinationArn: jsii.String("destinationArn"),
//   				Headers: []interface{}{
//   					&KafkaActionHeaderProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Key: jsii.String("key"),
//   				Partition: jsii.String("partition"),
//   				Topic: jsii.String("topic"),
//   			},
//   			Kinesis: &KinesisActionProperty{
//   				PartitionKey: jsii.String("partitionKey"),
//   				RoleArn: jsii.String("roleArn"),
//   				StreamName: jsii.String("streamName"),
//   			},
//   			Lambda: &LambdaActionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   			},
//   			Location: &LocationActionProperty{
//   				DeviceId: jsii.String("deviceId"),
//   				Latitude: jsii.String("latitude"),
//   				Longitude: jsii.String("longitude"),
//   				RoleArn: jsii.String("roleArn"),
//   				Timestamp: &TimestampProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.String("value"),
//   				},
//   				TrackerName: jsii.String("trackerName"),
//   			},
//   			OpenSearch: &OpenSearchActionProperty{
//   				Endpoint: jsii.String("endpoint"),
//   				Id: jsii.String("id"),
//   				Index: jsii.String("index"),
//   				RoleArn: jsii.String("roleArn"),
//   				Type: jsii.String("type"),
//   			},
//   			Republish: &RepublishActionProperty{
//   				Headers: &RepublishActionHeadersProperty{
//   					ContentType: jsii.String("contentType"),
//   					CorrelationData: jsii.String("correlationData"),
//   					MessageExpiry: jsii.String("messageExpiry"),
//   					PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   					ResponseTopic: jsii.String("responseTopic"),
//   					UserProperties: []interface{}{
//   						&UserPropertyProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Qos: jsii.Number(123),
//   				RoleArn: jsii.String("roleArn"),
//   				Topic: jsii.String("topic"),
//   			},
//   			S3: &S3ActionProperty{
//   				BucketName: jsii.String("bucketName"),
//   				CannedAcl: jsii.String("cannedAcl"),
//   				Key: jsii.String("key"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Sns: &SnsActionProperty{
//   				MessageFormat: jsii.String("messageFormat"),
//   				RoleArn: jsii.String("roleArn"),
//   				TargetArn: jsii.String("targetArn"),
//   			},
//   			Sqs: &SqsActionProperty{
//   				QueueUrl: jsii.String("queueUrl"),
//   				RoleArn: jsii.String("roleArn"),
//   				UseBase64: jsii.Boolean(false),
//   			},
//   			StepFunctions: &StepFunctionsActionProperty{
//   				ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   				RoleArn: jsii.String("roleArn"),
//   				StateMachineName: jsii.String("stateMachineName"),
//   			},
//   			Timestream: &TimestreamActionProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   				Dimensions: []interface{}{
//   					&TimestreamDimensionProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//   				Timestamp: &TimestreamTimestampProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		RuleDisabled: jsii.Boolean(false),
//   		Sql: jsii.String("sql"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
//
type CfnTopicRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTopicRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTopicRulePropsMixin
type jsiiProxy_CfnTopicRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTopicRulePropsMixin) Props() *CfnTopicRuleMixinProps {
	var returns *CfnTopicRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::TopicRule`.
func NewCfnTopicRulePropsMixin(props *CfnTopicRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTopicRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTopicRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnTopicRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::TopicRule`.
func NewCfnTopicRulePropsMixin_Override(c CfnTopicRulePropsMixin, props *CfnTopicRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnTopicRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTopicRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnTopicRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnTopicRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTopicRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

