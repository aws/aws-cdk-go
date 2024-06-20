package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTopicRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopicRuleProps := &CfnTopicRuleProps{
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
//   					LogGroupName: jsii.String("logGroupName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   				},
//   				CloudwatchMetric: &CloudwatchMetricActionProperty{
//   					MetricName: jsii.String("metricName"),
//   					MetricNamespace: jsii.String("metricNamespace"),
//   					MetricUnit: jsii.String("metricUnit"),
//   					MetricValue: jsii.String("metricValue"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					MetricTimestamp: jsii.String("metricTimestamp"),
//   				},
//   				DynamoDb: &DynamoDBActionProperty{
//   					HashKeyField: jsii.String("hashKeyField"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					RoleArn: jsii.String("roleArn"),
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					HashKeyType: jsii.String("hashKeyType"),
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
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
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   					Separator: jsii.String("separator"),
//   				},
//   				Http: &HttpActionProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					Auth: &HttpAuthorizationProperty{
//   						Sigv4: &SigV4AuthorizationProperty{
//   							RoleArn: jsii.String("roleArn"),
//   							ServiceName: jsii.String("serviceName"),
//   							SigningRegion: jsii.String("signingRegion"),
//   						},
//   					},
//   					ConfirmationUrl: jsii.String("confirmationUrl"),
//   					Headers: []interface{}{
//   						&HttpActionHeaderProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				IotAnalytics: &IotAnalyticsActionProperty{
//   					ChannelName: jsii.String("channelName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   				},
//   				IotEvents: &IotEventsActionProperty{
//   					InputName: jsii.String("inputName"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					BatchMode: jsii.Boolean(false),
//   					MessageId: jsii.String("messageId"),
//   				},
//   				IotSiteWise: &IotSiteWiseActionProperty{
//   					PutAssetPropertyValueEntries: []interface{}{
//   						&PutAssetPropertyValueEntryProperty{
//   							PropertyValues: []interface{}{
//   								&AssetPropertyValueProperty{
//   									Timestamp: &AssetPropertyTimestampProperty{
//   										TimeInSeconds: jsii.String("timeInSeconds"),
//
//   										// the properties below are optional
//   										OffsetInNanos: jsii.String("offsetInNanos"),
//   									},
//   									Value: &AssetPropertyVariantProperty{
//   										BooleanValue: jsii.String("booleanValue"),
//   										DoubleValue: jsii.String("doubleValue"),
//   										IntegerValue: jsii.String("integerValue"),
//   										StringValue: jsii.String("stringValue"),
//   									},
//
//   									// the properties below are optional
//   									Quality: jsii.String("quality"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   						},
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				Kafka: &KafkaActionProperty{
//   					ClientProperties: map[string]*string{
//   						"clientPropertiesKey": jsii.String("clientProperties"),
//   					},
//   					DestinationArn: jsii.String("destinationArn"),
//   					Topic: jsii.String("topic"),
//
//   					// the properties below are optional
//   					Headers: []interface{}{
//   						&KafkaActionHeaderProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Key: jsii.String("key"),
//   					Partition: jsii.String("partition"),
//   				},
//   				Kinesis: &KinesisActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					StreamName: jsii.String("streamName"),
//
//   					// the properties below are optional
//   					PartitionKey: jsii.String("partitionKey"),
//   				},
//   				Lambda: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   				Location: &LocationActionProperty{
//   					DeviceId: jsii.String("deviceId"),
//   					Latitude: jsii.String("latitude"),
//   					Longitude: jsii.String("longitude"),
//   					RoleArn: jsii.String("roleArn"),
//   					TrackerName: jsii.String("trackerName"),
//
//   					// the properties below are optional
//   					Timestamp: &TimestampProperty{
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   				},
//   				OpenSearch: &OpenSearchActionProperty{
//   					Endpoint: jsii.String("endpoint"),
//   					Id: jsii.String("id"),
//   					Index: jsii.String("index"),
//   					RoleArn: jsii.String("roleArn"),
//   					Type: jsii.String("type"),
//   				},
//   				Republish: &RepublishActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					Topic: jsii.String("topic"),
//
//   					// the properties below are optional
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
//   				},
//   				S3: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//   					Key: jsii.String("key"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					CannedAcl: jsii.String("cannedAcl"),
//   				},
//   				Sns: &SnsActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					TargetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					MessageFormat: jsii.String("messageFormat"),
//   				},
//   				Sqs: &SqsActionProperty{
//   					QueueUrl: jsii.String("queueUrl"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					UseBase64: jsii.Boolean(false),
//   				},
//   				StepFunctions: &StepFunctionsActionProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					StateMachineName: jsii.String("stateMachineName"),
//
//   					// the properties below are optional
//   					ExecutionNamePrefix: jsii.String("executionNamePrefix"),
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
//
//   					// the properties below are optional
//   					Timestamp: &TimestreamTimestampProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Sql: jsii.String("sql"),
//
//   		// the properties below are optional
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
//   				LogGroupName: jsii.String("logGroupName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   			},
//   			CloudwatchMetric: &CloudwatchMetricActionProperty{
//   				MetricName: jsii.String("metricName"),
//   				MetricNamespace: jsii.String("metricNamespace"),
//   				MetricUnit: jsii.String("metricUnit"),
//   				MetricValue: jsii.String("metricValue"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				MetricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			DynamoDb: &DynamoDBActionProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				HashKeyType: jsii.String("hashKeyType"),
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
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
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   				Separator: jsii.String("separator"),
//   			},
//   			Http: &HttpActionProperty{
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				Auth: &HttpAuthorizationProperty{
//   					Sigv4: &SigV4AuthorizationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						ServiceName: jsii.String("serviceName"),
//   						SigningRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				ConfirmationUrl: jsii.String("confirmationUrl"),
//   				Headers: []interface{}{
//   					&HttpActionHeaderProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			IotAnalytics: &IotAnalyticsActionProperty{
//   				ChannelName: jsii.String("channelName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   			},
//   			IotEvents: &IotEventsActionProperty{
//   				InputName: jsii.String("inputName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				BatchMode: jsii.Boolean(false),
//   				MessageId: jsii.String("messageId"),
//   			},
//   			IotSiteWise: &IotSiteWiseActionProperty{
//   				PutAssetPropertyValueEntries: []interface{}{
//   					&PutAssetPropertyValueEntryProperty{
//   						PropertyValues: []interface{}{
//   							&AssetPropertyValueProperty{
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								Quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						AssetId: jsii.String("assetId"),
//   						EntryId: jsii.String("entryId"),
//   						PropertyAlias: jsii.String("propertyAlias"),
//   						PropertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Kafka: &KafkaActionProperty{
//   				ClientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				DestinationArn: jsii.String("destinationArn"),
//   				Topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				Headers: []interface{}{
//   					&KafkaActionHeaderProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Key: jsii.String("key"),
//   				Partition: jsii.String("partition"),
//   			},
//   			Kinesis: &KinesisActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				StreamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				PartitionKey: jsii.String("partitionKey"),
//   			},
//   			Lambda: &LambdaActionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   			},
//   			Location: &LocationActionProperty{
//   				DeviceId: jsii.String("deviceId"),
//   				Latitude: jsii.String("latitude"),
//   				Longitude: jsii.String("longitude"),
//   				RoleArn: jsii.String("roleArn"),
//   				TrackerName: jsii.String("trackerName"),
//
//   				// the properties below are optional
//   				Timestamp: &TimestampProperty{
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   			OpenSearch: &OpenSearchActionProperty{
//   				Endpoint: jsii.String("endpoint"),
//   				Id: jsii.String("id"),
//   				Index: jsii.String("index"),
//   				RoleArn: jsii.String("roleArn"),
//   				Type: jsii.String("type"),
//   			},
//   			Republish: &RepublishActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				Topic: jsii.String("topic"),
//
//   				// the properties below are optional
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
//   			},
//   			S3: &S3ActionProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				CannedAcl: jsii.String("cannedAcl"),
//   			},
//   			Sns: &SnsActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				TargetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				MessageFormat: jsii.String("messageFormat"),
//   			},
//   			Sqs: &SqsActionProperty{
//   				QueueUrl: jsii.String("queueUrl"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				UseBase64: jsii.Boolean(false),
//   			},
//   			StepFunctions: &StepFunctionsActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				StateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				ExecutionNamePrefix: jsii.String("executionNamePrefix"),
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
//
//   				// the properties below are optional
//   				Timestamp: &TimestreamTimestampProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		RuleDisabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	RuleName: jsii.String("ruleName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
//
type CfnTopicRuleProps struct {
	// The rule payload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-topicrulepayload
	//
	TopicRulePayload interface{} `field:"required" json:"topicRulePayload" yaml:"topicRulePayload"`
	// The name of the rule.
	//
	// *Pattern* : `^[a-zA-Z0-9_]+$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

