package awsiot


// Describes the actions associated with a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   		alarmName: jsii.String("alarmName"),
//   		roleArn: jsii.String("roleArn"),
//   		stateReason: jsii.String("stateReason"),
//   		stateValue: jsii.String("stateValue"),
//   	},
//   	cloudwatchLogs: &cloudwatchLogsActionProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	cloudwatchMetric: &cloudwatchMetricActionProperty{
//   		metricName: jsii.String("metricName"),
//   		metricNamespace: jsii.String("metricNamespace"),
//   		metricUnit: jsii.String("metricUnit"),
//   		metricValue: jsii.String("metricValue"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		metricTimestamp: jsii.String("metricTimestamp"),
//   	},
//   	dynamoDb: &dynamoDBActionProperty{
//   		hashKeyField: jsii.String("hashKeyField"),
//   		hashKeyValue: jsii.String("hashKeyValue"),
//   		roleArn: jsii.String("roleArn"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		hashKeyType: jsii.String("hashKeyType"),
//   		payloadField: jsii.String("payloadField"),
//   		rangeKeyField: jsii.String("rangeKeyField"),
//   		rangeKeyType: jsii.String("rangeKeyType"),
//   		rangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	dynamoDBv2: &dynamoDBv2ActionProperty{
//   		putItem: &putItemInputProperty{
//   			tableName: jsii.String("tableName"),
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	elasticsearch: &elasticsearchActionProperty{
//   		endpoint: jsii.String("endpoint"),
//   		id: jsii.String("id"),
//   		index: jsii.String("index"),
//   		roleArn: jsii.String("roleArn"),
//   		type: jsii.String("type"),
//   	},
//   	firehose: &firehoseActionProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   		separator: jsii.String("separator"),
//   	},
//   	http: &httpActionProperty{
//   		url: jsii.String("url"),
//
//   		// the properties below are optional
//   		auth: &httpAuthorizationProperty{
//   			sigv4: &sigV4AuthorizationProperty{
//   				roleArn: jsii.String("roleArn"),
//   				serviceName: jsii.String("serviceName"),
//   				signingRegion: jsii.String("signingRegion"),
//   			},
//   		},
//   		confirmationUrl: jsii.String("confirmationUrl"),
//   		headers: []interface{}{
//   			&httpActionHeaderProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	iotAnalytics: &iotAnalyticsActionProperty{
//   		channelName: jsii.String("channelName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   	},
//   	iotEvents: &iotEventsActionProperty{
//   		inputName: jsii.String("inputName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		batchMode: jsii.Boolean(false),
//   		messageId: jsii.String("messageId"),
//   	},
//   	iotSiteWise: &iotSiteWiseActionProperty{
//   		putAssetPropertyValueEntries: []interface{}{
//   			&putAssetPropertyValueEntryProperty{
//   				propertyValues: []interface{}{
//   					&assetPropertyValueProperty{
//   						timestamp: &assetPropertyTimestampProperty{
//   							timeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							offsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   						value: &assetPropertyVariantProperty{
//   							booleanValue: jsii.String("booleanValue"),
//   							doubleValue: jsii.String("doubleValue"),
//   							integerValue: jsii.String("integerValue"),
//   							stringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						quality: jsii.String("quality"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	kafka: &kafkaActionProperty{
//   		clientProperties: map[string]*string{
//   			"clientPropertiesKey": jsii.String("clientProperties"),
//   		},
//   		destinationArn: jsii.String("destinationArn"),
//   		topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   		partition: jsii.String("partition"),
//   	},
//   	kinesis: &kinesisActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		streamName: jsii.String("streamName"),
//
//   		// the properties below are optional
//   		partitionKey: jsii.String("partitionKey"),
//   	},
//   	lambda: &lambdaActionProperty{
//   		functionArn: jsii.String("functionArn"),
//   	},
//   	location: &locationActionProperty{
//   		deviceId: jsii.String("deviceId"),
//   		latitude: jsii.String("latitude"),
//   		longitude: jsii.String("longitude"),
//   		roleArn: jsii.String("roleArn"),
//   		trackerName: jsii.String("trackerName"),
//
//   		// the properties below are optional
//   		timestamp: NewDate(),
//   	},
//   	openSearch: &openSearchActionProperty{
//   		endpoint: jsii.String("endpoint"),
//   		id: jsii.String("id"),
//   		index: jsii.String("index"),
//   		roleArn: jsii.String("roleArn"),
//   		type: jsii.String("type"),
//   	},
//   	republish: &republishActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		topic: jsii.String("topic"),
//
//   		// the properties below are optional
//   		headers: &republishActionHeadersProperty{
//   			contentType: jsii.String("contentType"),
//   			correlationData: jsii.String("correlationData"),
//   			messageExpiry: jsii.String("messageExpiry"),
//   			payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   			responseTopic: jsii.String("responseTopic"),
//   			userProperties: []interface{}{
//   				&userPropertyProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		qos: jsii.Number(123),
//   	},
//   	s3: &s3ActionProperty{
//   		bucketName: jsii.String("bucketName"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   	sns: &snsActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		targetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		messageFormat: jsii.String("messageFormat"),
//   	},
//   	sqs: &sqsActionProperty{
//   		queueUrl: jsii.String("queueUrl"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		useBase64: jsii.Boolean(false),
//   	},
//   	stepFunctions: &stepFunctionsActionProperty{
//   		roleArn: jsii.String("roleArn"),
//   		stateMachineName: jsii.String("stateMachineName"),
//
//   		// the properties below are optional
//   		executionNamePrefix: jsii.String("executionNamePrefix"),
//   	},
//   	timestream: &timestreamActionProperty{
//   		databaseName: jsii.String("databaseName"),
//   		dimensions: []interface{}{
//   			&timestreamDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		timestamp: &timestreamTimestampProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTopicRule_ActionProperty struct {
	// Change the state of a CloudWatch alarm.
	CloudwatchAlarm interface{} `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// Sends data to CloudWatch.
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Capture a CloudWatch metric.
	CloudwatchMetric interface{} `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Write to a DynamoDB table.
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Write to a DynamoDB table.
	//
	// This is a new version of the DynamoDB action. It allows you to write each attribute in an MQTT message payload into a separate DynamoDB column.
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Write data to an Amazon OpenSearch Service domain.
	//
	// > The `Elasticsearch` action can only be used by existing rule actions. To create a new rule action or to update an existing rule action, use the `OpenSearch` rule action instead. For more information, see [OpenSearchAction](https://docs.aws.amazon.com//iot/latest/apireference/API_OpenSearchAction.html) .
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Write to an Amazon Kinesis Firehose stream.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Send data to an HTTPS endpoint.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// Sends message data to an AWS IoT Analytics channel.
	IotAnalytics interface{} `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// Sends an input to an AWS IoT Events detector.
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends data from the MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// Write data to an Amazon Kinesis stream.
	Kinesis interface{} `field:"optional" json:"kinesis" yaml:"kinesis"`
	// Invoke a Lambda function.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// `CfnTopicRule.ActionProperty.Location`.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// Write data to an Amazon OpenSearch Service domain.
	OpenSearch interface{} `field:"optional" json:"openSearch" yaml:"openSearch"`
	// Publish to another MQTT topic.
	Republish interface{} `field:"optional" json:"republish" yaml:"republish"`
	// Write to an Amazon S3 bucket.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// Publish to an Amazon SNS topic.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Publish to an Amazon SQS queue.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
	// Starts execution of a Step Functions state machine.
	StepFunctions interface{} `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Writes attributes from an MQTT message.
	Timestream interface{} `field:"optional" json:"timestream" yaml:"timestream"`
}

