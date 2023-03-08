package awsiotevents


// An action to be performed when the `condition` is TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	clearTimer: &clearTimerProperty{
//   		timerName: jsii.String("timerName"),
//   	},
//   	dynamoDb: &dynamoDBProperty{
//   		hashKeyField: jsii.String("hashKeyField"),
//   		hashKeyValue: jsii.String("hashKeyValue"),
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		hashKeyType: jsii.String("hashKeyType"),
//   		operation: jsii.String("operation"),
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		payloadField: jsii.String("payloadField"),
//   		rangeKeyField: jsii.String("rangeKeyField"),
//   		rangeKeyType: jsii.String("rangeKeyType"),
//   		rangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	dynamoDBv2: &dynamoDBv2Property{
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	firehose: &firehoseProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		separator: jsii.String("separator"),
//   	},
//   	iotEvents: &iotEventsProperty{
//   		inputName: jsii.String("inputName"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	iotSiteWise: &iotSiteWiseProperty{
//   		propertyValue: &assetPropertyValueProperty{
//   			value: &assetPropertyVariantProperty{
//   				booleanValue: jsii.String("booleanValue"),
//   				doubleValue: jsii.String("doubleValue"),
//   				integerValue: jsii.String("integerValue"),
//   				stringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			quality: jsii.String("quality"),
//   			timestamp: &assetPropertyTimestampProperty{
//   				timeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				offsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		assetId: jsii.String("assetId"),
//   		entryId: jsii.String("entryId"),
//   		propertyAlias: jsii.String("propertyAlias"),
//   		propertyId: jsii.String("propertyId"),
//   	},
//   	iotTopicPublish: &iotTopicPublishProperty{
//   		mqttTopic: jsii.String("mqttTopic"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	lambda: &lambdaProperty{
//   		functionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	resetTimer: &resetTimerProperty{
//   		timerName: jsii.String("timerName"),
//   	},
//   	setTimer: &setTimerProperty{
//   		timerName: jsii.String("timerName"),
//
//   		// the properties below are optional
//   		durationExpression: jsii.String("durationExpression"),
//   		seconds: jsii.Number(123),
//   	},
//   	setVariable: &setVariableProperty{
//   		value: jsii.String("value"),
//   		variableName: jsii.String("variableName"),
//   	},
//   	sns: &snsProperty{
//   		targetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sqs: &sqsProperty{
//   		queueUrl: jsii.String("queueUrl"),
//
//   		// the properties below are optional
//   		payload: &payloadProperty{
//   			contentExpression: jsii.String("contentExpression"),
//   			type: jsii.String("type"),
//   		},
//   		useBase64: jsii.Boolean(false),
//   	},
//   }
//
type CfnDetectorModel_ActionProperty struct {
	// Information needed to clear the timer.
	ClearTimer interface{} `field:"optional" json:"clearTimer" yaml:"clearTimer"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide* .
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide* .
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends AWS IoT Events input, which passes information about the detector model instance and the event that triggered the action.
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to an asset property in AWS IoT SiteWise .
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Publishes an MQTT message with the given topic to the AWS IoT message broker.
	IotTopicPublish interface{} `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Information needed to reset the timer.
	ResetTimer interface{} `field:"optional" json:"resetTimer" yaml:"resetTimer"`
	// Information needed to set the timer.
	SetTimer interface{} `field:"optional" json:"setTimer" yaml:"setTimer"`
	// Sets a variable to a specified value.
	SetVariable interface{} `field:"optional" json:"setVariable" yaml:"setVariable"`
	// Sends an Amazon SNS message.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Sends an Amazon SNS message.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
}

