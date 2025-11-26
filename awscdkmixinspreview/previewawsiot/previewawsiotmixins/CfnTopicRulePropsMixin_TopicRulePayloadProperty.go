package previewawsiotmixins


// Describes a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicRulePayloadProperty := &TopicRulePayloadProperty{
//   	Actions: []interface{}{
//   		&ActionProperty{
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
//   				ConfirmationUrl: jsii.String("confirmationUrl"),
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
//   	},
//   	AwsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	Description: jsii.String("description"),
//   	ErrorAction: &ActionProperty{
//   		CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   			AlarmName: jsii.String("alarmName"),
//   			RoleArn: jsii.String("roleArn"),
//   			StateReason: jsii.String("stateReason"),
//   			StateValue: jsii.String("stateValue"),
//   		},
//   		CloudwatchLogs: &CloudwatchLogsActionProperty{
//   			BatchMode: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		CloudwatchMetric: &CloudwatchMetricActionProperty{
//   			MetricName: jsii.String("metricName"),
//   			MetricNamespace: jsii.String("metricNamespace"),
//   			MetricTimestamp: jsii.String("metricTimestamp"),
//   			MetricUnit: jsii.String("metricUnit"),
//   			MetricValue: jsii.String("metricValue"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		DynamoDb: &DynamoDBActionProperty{
//   			HashKeyField: jsii.String("hashKeyField"),
//   			HashKeyType: jsii.String("hashKeyType"),
//   			HashKeyValue: jsii.String("hashKeyValue"),
//   			PayloadField: jsii.String("payloadField"),
//   			RangeKeyField: jsii.String("rangeKeyField"),
//   			RangeKeyType: jsii.String("rangeKeyType"),
//   			RangeKeyValue: jsii.String("rangeKeyValue"),
//   			RoleArn: jsii.String("roleArn"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		DynamoDBv2: &DynamoDBv2ActionProperty{
//   			PutItem: &PutItemInputProperty{
//   				TableName: jsii.String("tableName"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Elasticsearch: &ElasticsearchActionProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Id: jsii.String("id"),
//   			Index: jsii.String("index"),
//   			RoleArn: jsii.String("roleArn"),
//   			Type: jsii.String("type"),
//   		},
//   		Firehose: &FirehoseActionProperty{
//   			BatchMode: jsii.Boolean(false),
//   			DeliveryStreamName: jsii.String("deliveryStreamName"),
//   			RoleArn: jsii.String("roleArn"),
//   			Separator: jsii.String("separator"),
//   		},
//   		Http: &HttpActionProperty{
//   			Auth: &HttpAuthorizationProperty{
//   				Sigv4: &SigV4AuthorizationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					ServiceName: jsii.String("serviceName"),
//   					SigningRegion: jsii.String("signingRegion"),
//   				},
//   			},
//   			ConfirmationUrl: jsii.String("confirmationUrl"),
//   			Headers: []interface{}{
//   				&HttpActionHeaderProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Url: jsii.String("url"),
//   		},
//   		IotAnalytics: &IotAnalyticsActionProperty{
//   			BatchMode: jsii.Boolean(false),
//   			ChannelName: jsii.String("channelName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		IotEvents: &IotEventsActionProperty{
//   			BatchMode: jsii.Boolean(false),
//   			InputName: jsii.String("inputName"),
//   			MessageId: jsii.String("messageId"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		IotSiteWise: &IotSiteWiseActionProperty{
//   			PutAssetPropertyValueEntries: []interface{}{
//   				&PutAssetPropertyValueEntryProperty{
//   					AssetId: jsii.String("assetId"),
//   					EntryId: jsii.String("entryId"),
//   					PropertyAlias: jsii.String("propertyAlias"),
//   					PropertyId: jsii.String("propertyId"),
//   					PropertyValues: []interface{}{
//   						&AssetPropertyValueProperty{
//   							Quality: jsii.String("quality"),
//   							Timestamp: &AssetPropertyTimestampProperty{
//   								OffsetInNanos: jsii.String("offsetInNanos"),
//   								TimeInSeconds: jsii.String("timeInSeconds"),
//   							},
//   							Value: &AssetPropertyVariantProperty{
//   								BooleanValue: jsii.String("booleanValue"),
//   								DoubleValue: jsii.String("doubleValue"),
//   								IntegerValue: jsii.String("integerValue"),
//   								StringValue: jsii.String("stringValue"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Kafka: &KafkaActionProperty{
//   			ClientProperties: map[string]*string{
//   				"clientPropertiesKey": jsii.String("clientProperties"),
//   			},
//   			DestinationArn: jsii.String("destinationArn"),
//   			Headers: []interface{}{
//   				&KafkaActionHeaderProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Key: jsii.String("key"),
//   			Partition: jsii.String("partition"),
//   			Topic: jsii.String("topic"),
//   		},
//   		Kinesis: &KinesisActionProperty{
//   			PartitionKey: jsii.String("partitionKey"),
//   			RoleArn: jsii.String("roleArn"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   		Lambda: &LambdaActionProperty{
//   			FunctionArn: jsii.String("functionArn"),
//   		},
//   		Location: &LocationActionProperty{
//   			DeviceId: jsii.String("deviceId"),
//   			Latitude: jsii.String("latitude"),
//   			Longitude: jsii.String("longitude"),
//   			RoleArn: jsii.String("roleArn"),
//   			Timestamp: &TimestampProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.String("value"),
//   			},
//   			TrackerName: jsii.String("trackerName"),
//   		},
//   		OpenSearch: &OpenSearchActionProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Id: jsii.String("id"),
//   			Index: jsii.String("index"),
//   			RoleArn: jsii.String("roleArn"),
//   			Type: jsii.String("type"),
//   		},
//   		Republish: &RepublishActionProperty{
//   			Headers: &RepublishActionHeadersProperty{
//   				ContentType: jsii.String("contentType"),
//   				CorrelationData: jsii.String("correlationData"),
//   				MessageExpiry: jsii.String("messageExpiry"),
//   				PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   				ResponseTopic: jsii.String("responseTopic"),
//   				UserProperties: []interface{}{
//   					&UserPropertyProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Qos: jsii.Number(123),
//   			RoleArn: jsii.String("roleArn"),
//   			Topic: jsii.String("topic"),
//   		},
//   		S3: &S3ActionProperty{
//   			BucketName: jsii.String("bucketName"),
//   			CannedAcl: jsii.String("cannedAcl"),
//   			Key: jsii.String("key"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Sns: &SnsActionProperty{
//   			MessageFormat: jsii.String("messageFormat"),
//   			RoleArn: jsii.String("roleArn"),
//   			TargetArn: jsii.String("targetArn"),
//   		},
//   		Sqs: &SqsActionProperty{
//   			QueueUrl: jsii.String("queueUrl"),
//   			RoleArn: jsii.String("roleArn"),
//   			UseBase64: jsii.Boolean(false),
//   		},
//   		StepFunctions: &StepFunctionsActionProperty{
//   			ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   			RoleArn: jsii.String("roleArn"),
//   			StateMachineName: jsii.String("stateMachineName"),
//   		},
//   		Timestream: &TimestreamActionProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Dimensions: []interface{}{
//   				&TimestreamDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			TableName: jsii.String("tableName"),
//   			Timestamp: &TimestreamTimestampProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RuleDisabled: jsii.Boolean(false),
//   	Sql: jsii.String("sql"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
//
type CfnTopicRulePropsMixin_TopicRulePayloadProperty struct {
	// The actions associated with the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The version of the SQL rules engine to use when evaluating the rule.
	//
	// The default value is 2015-10-08.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
	//
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The description of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The action to take when an error occurs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-erroraction
	//
	ErrorAction interface{} `field:"optional" json:"errorAction" yaml:"errorAction"`
	// Specifies whether the rule is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
	//
	RuleDisabled interface{} `field:"optional" json:"ruleDisabled" yaml:"ruleDisabled"`
	// The SQL statement used to query the topic.
	//
	// For more information, see [AWS IoT SQL Reference](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) in the *AWS IoT Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
	//
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
}

