package awsiot


// Describes a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicRulePayloadProperty := &topicRulePayloadProperty{
//   	actions: []interface{}{
//   		&actionProperty{
//   			cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   				alarmName: jsii.String("alarmName"),
//   				roleArn: jsii.String("roleArn"),
//   				stateReason: jsii.String("stateReason"),
//   				stateValue: jsii.String("stateValue"),
//   			},
//   			cloudwatchLogs: &cloudwatchLogsActionProperty{
//   				logGroupName: jsii.String("logGroupName"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			cloudwatchMetric: &cloudwatchMetricActionProperty{
//   				metricName: jsii.String("metricName"),
//   				metricNamespace: jsii.String("metricNamespace"),
//   				metricUnit: jsii.String("metricUnit"),
//   				metricValue: jsii.String("metricValue"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				metricTimestamp: jsii.String("metricTimestamp"),
//   			},
//   			dynamoDb: &dynamoDBActionProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2ActionProperty{
//   				putItem: &putItemInputProperty{
//   					tableName: jsii.String("tableName"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			elasticsearch: &elasticsearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			firehose: &firehoseActionProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				separator: jsii.String("separator"),
//   			},
//   			http: &httpActionProperty{
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				auth: &httpAuthorizationProperty{
//   					sigv4: &sigV4AuthorizationProperty{
//   						roleArn: jsii.String("roleArn"),
//   						serviceName: jsii.String("serviceName"),
//   						signingRegion: jsii.String("signingRegion"),
//   					},
//   				},
//   				confirmationUrl: jsii.String("confirmationUrl"),
//   				headers: []interface{}{
//   					&httpActionHeaderProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			iotAnalytics: &iotAnalyticsActionProperty{
//   				channelName: jsii.String("channelName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   			},
//   			iotEvents: &iotEventsActionProperty{
//   				inputName: jsii.String("inputName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				batchMode: jsii.Boolean(false),
//   				messageId: jsii.String("messageId"),
//   			},
//   			iotSiteWise: &iotSiteWiseActionProperty{
//   				putAssetPropertyValueEntries: []interface{}{
//   					&putAssetPropertyValueEntryProperty{
//   						propertyValues: []interface{}{
//   							&assetPropertyValueProperty{
//   								timestamp: &assetPropertyTimestampProperty{
//   									timeInSeconds: jsii.String("timeInSeconds"),
//
//   									// the properties below are optional
//   									offsetInNanos: jsii.String("offsetInNanos"),
//   								},
//   								value: &assetPropertyVariantProperty{
//   									booleanValue: jsii.String("booleanValue"),
//   									doubleValue: jsii.String("doubleValue"),
//   									integerValue: jsii.String("integerValue"),
//   									stringValue: jsii.String("stringValue"),
//   								},
//
//   								// the properties below are optional
//   								quality: jsii.String("quality"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						assetId: jsii.String("assetId"),
//   						entryId: jsii.String("entryId"),
//   						propertyAlias: jsii.String("propertyAlias"),
//   						propertyId: jsii.String("propertyId"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kafka: &kafkaActionProperty{
//   				clientProperties: map[string]*string{
//   					"clientPropertiesKey": jsii.String("clientProperties"),
//   				},
//   				destinationArn: jsii.String("destinationArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				partition: jsii.String("partition"),
//   			},
//   			kinesis: &kinesisActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamName: jsii.String("streamName"),
//
//   				// the properties below are optional
//   				partitionKey: jsii.String("partitionKey"),
//   			},
//   			lambda: &lambdaActionProperty{
//   				functionArn: jsii.String("functionArn"),
//   			},
//   			location: &locationActionProperty{
//   				deviceId: jsii.String("deviceId"),
//   				latitude: jsii.String("latitude"),
//   				longitude: jsii.String("longitude"),
//   				roleArn: jsii.String("roleArn"),
//   				trackerName: jsii.String("trackerName"),
//
//   				// the properties below are optional
//   				timestamp: NewDate(),
//   			},
//   			openSearch: &openSearchActionProperty{
//   				endpoint: jsii.String("endpoint"),
//   				id: jsii.String("id"),
//   				index: jsii.String("index"),
//   				roleArn: jsii.String("roleArn"),
//   				type: jsii.String("type"),
//   			},
//   			republish: &republishActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				headers: &republishActionHeadersProperty{
//   					contentType: jsii.String("contentType"),
//   					correlationData: jsii.String("correlationData"),
//   					messageExpiry: jsii.String("messageExpiry"),
//   					payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   					responseTopic: jsii.String("responseTopic"),
//   					userProperties: []interface{}{
//   						&userPropertyProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				qos: jsii.Number(123),
//   			},
//   			s3: &s3ActionProperty{
//   				bucketName: jsii.String("bucketName"),
//   				key: jsii.String("key"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				cannedAcl: jsii.String("cannedAcl"),
//   			},
//   			sns: &snsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				messageFormat: jsii.String("messageFormat"),
//   			},
//   			sqs: &sqsActionProperty{
//   				queueUrl: jsii.String("queueUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				useBase64: jsii.Boolean(false),
//   			},
//   			stepFunctions: &stepFunctionsActionProperty{
//   				roleArn: jsii.String("roleArn"),
//   				stateMachineName: jsii.String("stateMachineName"),
//
//   				// the properties below are optional
//   				executionNamePrefix: jsii.String("executionNamePrefix"),
//   			},
//   			timestream: &timestreamActionProperty{
//   				databaseName: jsii.String("databaseName"),
//   				dimensions: []interface{}{
//   					&timestreamDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				timestamp: &timestreamTimestampProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	description: jsii.String("description"),
//   	errorAction: &actionProperty{
//   		cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   			alarmName: jsii.String("alarmName"),
//   			roleArn: jsii.String("roleArn"),
//   			stateReason: jsii.String("stateReason"),
//   			stateValue: jsii.String("stateValue"),
//   		},
//   		cloudwatchLogs: &cloudwatchLogsActionProperty{
//   			logGroupName: jsii.String("logGroupName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		cloudwatchMetric: &cloudwatchMetricActionProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricUnit: jsii.String("metricUnit"),
//   			metricValue: jsii.String("metricValue"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			metricTimestamp: jsii.String("metricTimestamp"),
//   		},
//   		dynamoDb: &dynamoDBActionProperty{
//   			hashKeyField: jsii.String("hashKeyField"),
//   			hashKeyValue: jsii.String("hashKeyValue"),
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			hashKeyType: jsii.String("hashKeyType"),
//   			payloadField: jsii.String("payloadField"),
//   			rangeKeyField: jsii.String("rangeKeyField"),
//   			rangeKeyType: jsii.String("rangeKeyType"),
//   			rangeKeyValue: jsii.String("rangeKeyValue"),
//   		},
//   		dynamoDBv2: &dynamoDBv2ActionProperty{
//   			putItem: &putItemInputProperty{
//   				tableName: jsii.String("tableName"),
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		elasticsearch: &elasticsearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		firehose: &firehoseActionProperty{
//   			deliveryStreamName: jsii.String("deliveryStreamName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			separator: jsii.String("separator"),
//   		},
//   		http: &httpActionProperty{
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			auth: &httpAuthorizationProperty{
//   				sigv4: &sigV4AuthorizationProperty{
//   					roleArn: jsii.String("roleArn"),
//   					serviceName: jsii.String("serviceName"),
//   					signingRegion: jsii.String("signingRegion"),
//   				},
//   			},
//   			confirmationUrl: jsii.String("confirmationUrl"),
//   			headers: []interface{}{
//   				&httpActionHeaderProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		iotAnalytics: &iotAnalyticsActionProperty{
//   			channelName: jsii.String("channelName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   		},
//   		iotEvents: &iotEventsActionProperty{
//   			inputName: jsii.String("inputName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			messageId: jsii.String("messageId"),
//   		},
//   		iotSiteWise: &iotSiteWiseActionProperty{
//   			putAssetPropertyValueEntries: []interface{}{
//   				&putAssetPropertyValueEntryProperty{
//   					propertyValues: []interface{}{
//   						&assetPropertyValueProperty{
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kafka: &kafkaActionProperty{
//   			clientProperties: map[string]*string{
//   				"clientPropertiesKey": jsii.String("clientProperties"),
//   			},
//   			destinationArn: jsii.String("destinationArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   			partition: jsii.String("partition"),
//   		},
//   		kinesis: &kinesisActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			streamName: jsii.String("streamName"),
//
//   			// the properties below are optional
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		lambda: &lambdaActionProperty{
//   			functionArn: jsii.String("functionArn"),
//   		},
//   		location: &locationActionProperty{
//   			deviceId: jsii.String("deviceId"),
//   			latitude: jsii.String("latitude"),
//   			longitude: jsii.String("longitude"),
//   			roleArn: jsii.String("roleArn"),
//   			trackerName: jsii.String("trackerName"),
//
//   			// the properties below are optional
//   			timestamp: NewDate(),
//   		},
//   		openSearch: &openSearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		republish: &republishActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			headers: &republishActionHeadersProperty{
//   				contentType: jsii.String("contentType"),
//   				correlationData: jsii.String("correlationData"),
//   				messageExpiry: jsii.String("messageExpiry"),
//   				payloadFormatIndicator: jsii.String("payloadFormatIndicator"),
//   				responseTopic: jsii.String("responseTopic"),
//   				userProperties: []interface{}{
//   					&userPropertyProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			qos: jsii.Number(123),
//   		},
//   		s3: &s3ActionProperty{
//   			bucketName: jsii.String("bucketName"),
//   			key: jsii.String("key"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   		sns: &snsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//
//   			// the properties below are optional
//   			messageFormat: jsii.String("messageFormat"),
//   		},
//   		sqs: &sqsActionProperty{
//   			queueUrl: jsii.String("queueUrl"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			useBase64: jsii.Boolean(false),
//   		},
//   		stepFunctions: &stepFunctionsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			stateMachineName: jsii.String("stateMachineName"),
//
//   			// the properties below are optional
//   			executionNamePrefix: jsii.String("executionNamePrefix"),
//   		},
//   		timestream: &timestreamActionProperty{
//   			databaseName: jsii.String("databaseName"),
//   			dimensions: []interface{}{
//   				&timestreamDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			timestamp: &timestreamTimestampProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ruleDisabled: jsii.Boolean(false),
//   }
//
type CfnTopicRule_TopicRulePayloadProperty struct {
	// The actions associated with the rule.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The SQL statement used to query the topic.
	//
	// For more information, see [AWS IoT SQL Reference](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) in the *AWS IoT Developer Guide* .
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	//
	// The default value is 2015-10-08.
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The description of the rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The action to take when an error occurs.
	ErrorAction interface{} `field:"optional" json:"errorAction" yaml:"errorAction"`
	// Specifies whether the rule is disabled.
	RuleDisabled interface{} `field:"optional" json:"ruleDisabled" yaml:"ruleDisabled"`
}

