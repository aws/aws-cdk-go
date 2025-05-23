package awsiot


// Describes the actions associated with a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	CloudwatchAlarm: &CloudwatchAlarmActionProperty{
//   		AlarmName: jsii.String("alarmName"),
//   		RoleArn: jsii.String("roleArn"),
//   		StateReason: jsii.String("stateReason"),
//   		StateValue: jsii.String("stateValue"),
//   	},
//   	CloudwatchLogs: &CloudwatchLogsActionProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BatchMode: jsii.Boolean(false),
//   	},
//   	CloudwatchMetric: &CloudwatchMetricActionProperty{
//   		MetricName: jsii.String("metricName"),
//   		MetricNamespace: jsii.String("metricNamespace"),
//   		MetricUnit: jsii.String("metricUnit"),
//   		MetricValue: jsii.String("metricValue"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		MetricTimestamp: jsii.String("metricTimestamp"),
//   	},
//   	DynamoDb: &DynamoDBActionProperty{
//   		HashKeyField: jsii.String("hashKeyField"),
//   		HashKeyValue: jsii.String("hashKeyValue"),
//   		RoleArn: jsii.String("roleArn"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		HashKeyType: jsii.String("hashKeyType"),
//   		PayloadField: jsii.String("payloadField"),
//   		RangeKeyField: jsii.String("rangeKeyField"),
//   		RangeKeyType: jsii.String("rangeKeyType"),
//   		RangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	DynamoDBv2: &DynamoDBv2ActionProperty{
//   		PutItem: &PutItemInputProperty{
//   			TableName: jsii.String("tableName"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Elasticsearch: &ElasticsearchActionProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Id: jsii.String("id"),
//   		Index: jsii.String("index"),
//   		RoleArn: jsii.String("roleArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Firehose: &FirehoseActionProperty{
//   		DeliveryStreamName: jsii.String("deliveryStreamName"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BatchMode: jsii.Boolean(false),
//   		Separator: jsii.String("separator"),
//   	},
//   	Http: &HttpActionProperty{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		Auth: &HttpAuthorizationProperty{
//   			Sigv4: &SigV4AuthorizationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				ServiceName: jsii.String("serviceName"),
//   				SigningRegion: jsii.String("signingRegion"),
//   			},
//   		},
//   		ConfirmationUrl: jsii.String("confirmationUrl"),
//   		Headers: []interface{}{
//   			&HttpActionHeaderProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	IotAnalytics: &IotAnalyticsActionProperty{
//   		ChannelName: jsii.String("channelName"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BatchMode: jsii.Boolean(false),
//   	},
//   	IotEvents: &IotEventsActionProperty{
//   		InputName: jsii.String("inputName"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BatchMode: jsii.Boolean(false),
//   		MessageId: jsii.String("messageId"),
//   	},
//   	IotSiteWise: &IotSiteWiseActionProperty{
//   		PutAssetPropertyValueEntries: []interface{}{
//   			&PutAssetPropertyValueEntryProperty{
//   				PropertyValues: []interface{}{
//   					&AssetPropertyValueProperty{
//   						Timestamp: &AssetPropertyTimestampProperty{
//   							TimeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							OffsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   						Value: &AssetPropertyVariantProperty{
//   							BooleanValue: jsii.String("booleanValue"),
//   							DoubleValue: jsii.String("doubleValue"),
//   							IntegerValue: jsii.String("integerValue"),
//   							StringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						Quality: jsii.String("quality"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				AssetId: jsii.String("assetId"),
//   				EntryId: jsii.String("entryId"),
//   				PropertyAlias: jsii.String("propertyAlias"),
//   				PropertyId: jsii.String("propertyId"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Kafka: &KafkaActionProperty{
//   		ClientProperties: map[string]*string{
//   			"clientPropertiesKey": jsii.String("clientProperties"),
//   		},
//   		DestinationArn: jsii.String("destinationArn"),
//   		Topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		Headers: []interface{}{
//   			&KafkaActionHeaderProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Key: jsii.String("key"),
//   		Partition: jsii.String("partition"),
//   	},
//   	Kinesis: &KinesisActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		StreamName: jsii.String("streamName"),
//
//   		// the properties below are optional
//   		PartitionKey: jsii.String("partitionKey"),
//   	},
//   	Lambda: &LambdaActionProperty{
//   		FunctionArn: jsii.String("functionArn"),
//   	},
//   	Location: &LocationActionProperty{
//   		DeviceId: jsii.String("deviceId"),
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   		RoleArn: jsii.String("roleArn"),
//   		TrackerName: jsii.String("trackerName"),
//
//   		// the properties below are optional
//   		Timestamp: &TimestampProperty{
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	OpenSearch: &OpenSearchActionProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Id: jsii.String("id"),
//   		Index: jsii.String("index"),
//   		RoleArn: jsii.String("roleArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Republish: &RepublishActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		Topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		Headers: &RepublishActionHeadersProperty{
//   			ContentType: jsii.String("contentType"),
//   			CorrelationData: jsii.String("correlationData"),
//   			MessageExpiry: jsii.String("messageExpiry"),
//   			PayloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   			ResponseTopic: jsii.String("responseTopic"),
//   			UserProperties: []interface{}{
//   				&UserPropertyProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Qos: jsii.Number(123),
//   	},
//   	S3: &S3ActionProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Key: jsii.String("key"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		CannedAcl: jsii.String("cannedAcl"),
//   	},
//   	Sns: &SnsActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		TargetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		MessageFormat: jsii.String("messageFormat"),
//   	},
//   	Sqs: &SqsActionProperty{
//   		QueueUrl: jsii.String("queueUrl"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		UseBase64: jsii.Boolean(false),
//   	},
//   	StepFunctions: &StepFunctionsActionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		StateMachineName: jsii.String("stateMachineName"),
//
//   		// the properties below are optional
//   		ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   	},
//   	Timestream: &TimestreamActionProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Dimensions: []interface{}{
//   			&TimestreamDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		Timestamp: &TimestreamTimestampProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html
//
type CfnTopicRule_ActionProperty struct {
	// Change the state of a CloudWatch alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchalarm
	//
	CloudwatchAlarm interface{} `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// Sends data to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchlogs
	//
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Capture a CloudWatch metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchmetric
	//
	CloudwatchMetric interface{} `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Write to a DynamoDB table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodb
	//
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Write to a DynamoDB table.
	//
	// This is a new version of the DynamoDB action. It allows you to write each attribute in an MQTT message payload into a separate DynamoDB column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodbv2
	//
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Write data to an Amazon OpenSearch Service domain.
	//
	// > The `Elasticsearch` action can only be used by existing rule actions. To create a new rule action or to update an existing rule action, use the `OpenSearch` rule action instead. For more information, see [OpenSearchAction](https://docs.aws.amazon.com//iot/latest/apireference/API_OpenSearchAction.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-elasticsearch
	//
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Write to an Amazon Kinesis Firehose stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Send data to an HTTPS endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-http
	//
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// Sends message data to an AWS IoT Analytics channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotanalytics
	//
	IotAnalytics interface{} `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// Sends an input to an AWS IoT Events detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotevents
	//
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends data from the MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotsitewise
	//
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-kafka
	//
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// Write data to an Amazon Kinesis stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-kinesis
	//
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
	// Invoke a Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Sends device location data to [Amazon Location Service](https://docs.aws.amazon.com//location/latest/developerguide/welcome.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// Write data to an Amazon OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-opensearch
	//
	OpenSearch interface{} `field:"optional" json:"openSearch" yaml:"openSearch"`
	// Publish to another MQTT topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-republish
	//
	Republish interface{} `field:"optional" json:"republish" yaml:"republish"`
	// Write to an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// Publish to an Amazon SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sns
	//
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Publish to an Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sqs
	//
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
	// Starts execution of a Step Functions state machine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-stepfunctions
	//
	StepFunctions interface{} `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Writes attributes from an MQTT message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-timestream
	//
	Timestream interface{} `field:"optional" json:"timestream" yaml:"timestream"`
}

