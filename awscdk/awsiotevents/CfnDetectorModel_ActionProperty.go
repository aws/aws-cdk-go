package awsiotevents


// An action to be performed when the `condition` is TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	ClearTimer: &ClearTimerProperty{
//   		TimerName: jsii.String("timerName"),
//   	},
//   	DynamoDb: &DynamoDBProperty{
//   		HashKeyField: jsii.String("hashKeyField"),
//   		HashKeyValue: jsii.String("hashKeyValue"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		HashKeyType: jsii.String("hashKeyType"),
//   		Operation: jsii.String("operation"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		PayloadField: jsii.String("payloadField"),
//   		RangeKeyField: jsii.String("rangeKeyField"),
//   		RangeKeyType: jsii.String("rangeKeyType"),
//   		RangeKeyValue: jsii.String("rangeKeyValue"),
//   	},
//   	DynamoDBv2: &DynamoDBv2Property{
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Firehose: &FirehoseProperty{
//   		DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		Separator: jsii.String("separator"),
//   	},
//   	IotEvents: &IotEventsProperty{
//   		InputName: jsii.String("inputName"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	IotSiteWise: &IotSiteWiseProperty{
//   		PropertyValue: &AssetPropertyValueProperty{
//   			Value: &AssetPropertyVariantProperty{
//   				BooleanValue: jsii.String("booleanValue"),
//   				DoubleValue: jsii.String("doubleValue"),
//   				IntegerValue: jsii.String("integerValue"),
//   				StringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			Quality: jsii.String("quality"),
//   			Timestamp: &AssetPropertyTimestampProperty{
//   				TimeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				OffsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		AssetId: jsii.String("assetId"),
//   		EntryId: jsii.String("entryId"),
//   		PropertyAlias: jsii.String("propertyAlias"),
//   		PropertyId: jsii.String("propertyId"),
//   	},
//   	IotTopicPublish: &IotTopicPublishProperty{
//   		MqttTopic: jsii.String("mqttTopic"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Lambda: &LambdaProperty{
//   		FunctionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ResetTimer: &ResetTimerProperty{
//   		TimerName: jsii.String("timerName"),
//   	},
//   	SetTimer: &SetTimerProperty{
//   		TimerName: jsii.String("timerName"),
//
//   		// the properties below are optional
//   		DurationExpression: jsii.String("durationExpression"),
//   		Seconds: jsii.Number(123),
//   	},
//   	SetVariable: &SetVariableProperty{
//   		Value: jsii.String("value"),
//   		VariableName: jsii.String("variableName"),
//   	},
//   	Sns: &SnsProperty{
//   		TargetArn: jsii.String("targetArn"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Sqs: &SqsProperty{
//   		QueueUrl: jsii.String("queueUrl"),
//
//   		// the properties below are optional
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		UseBase64: jsii.Boolean(false),
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

