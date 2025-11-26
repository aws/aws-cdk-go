package previewawsioteventsmixins


// Specifies one of the following actions to receive notifications when the alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   alarmActionProperty := &AlarmActionProperty{
//   	DynamoDb: &DynamoDBProperty{
//   		HashKeyField: jsii.String("hashKeyField"),
//   		HashKeyType: jsii.String("hashKeyType"),
//   		HashKeyValue: jsii.String("hashKeyValue"),
//   		Operation: jsii.String("operation"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		PayloadField: jsii.String("payloadField"),
//   		RangeKeyField: jsii.String("rangeKeyField"),
//   		RangeKeyType: jsii.String("rangeKeyType"),
//   		RangeKeyValue: jsii.String("rangeKeyValue"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	DynamoDBv2: &DynamoDBv2Property{
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		TableName: jsii.String("tableName"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		DeliveryStreamName: jsii.String("deliveryStreamName"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		Separator: jsii.String("separator"),
//   	},
//   	IotEvents: &IotEventsProperty{
//   		InputName: jsii.String("inputName"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	IotSiteWise: &IotSiteWiseProperty{
//   		AssetId: jsii.String("assetId"),
//   		EntryId: jsii.String("entryId"),
//   		PropertyAlias: jsii.String("propertyAlias"),
//   		PropertyId: jsii.String("propertyId"),
//   		PropertyValue: &AssetPropertyValueProperty{
//   			Quality: jsii.String("quality"),
//   			Timestamp: &AssetPropertyTimestampProperty{
//   				OffsetInNanos: jsii.String("offsetInNanos"),
//   				TimeInSeconds: jsii.String("timeInSeconds"),
//   			},
//   			Value: &AssetPropertyVariantProperty{
//   				BooleanValue: jsii.String("booleanValue"),
//   				DoubleValue: jsii.String("doubleValue"),
//   				IntegerValue: jsii.String("integerValue"),
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	IotTopicPublish: &IotTopicPublishProperty{
//   		MqttTopic: jsii.String("mqttTopic"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Lambda: &LambdaProperty{
//   		FunctionArn: jsii.String("functionArn"),
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Sns: &SnsProperty{
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Sqs: &SqsProperty{
//   		Payload: &PayloadProperty{
//   			ContentExpression: jsii.String("contentExpression"),
//   			Type: jsii.String("type"),
//   		},
//   		QueueUrl: jsii.String("queueUrl"),
//   		UseBase64: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html
//
type CfnAlarmModelPropsMixin_AlarmActionProperty struct {
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
	//
	// You must use expressions for all parameters in `DynamoDBAction` . The expressions accept literals, operators, functions, references, and substitution templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `hashKeyType` parameter can be `'STRING'` .
	// - For references, you must specify either variables or input values. For example, the value for the `hashKeyField` parameter can be `$input.GreenhouseInput.name` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `hashKeyValue` parameter uses a substitution template.
	//
	// `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`
	// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `tableName` parameter uses a string concatenation.
	//
	// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	//
	// If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the `payloadField` parameter is `<payload-field>_raw` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-dynamodb
	//
	DynamoDb interface{} `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html) . A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
	//
	// You must use expressions for all parameters in `DynamoDBv2Action` . The expressions accept literals, operators, functions, references, and substitution templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `tableName` parameter can be `'GreenhouseTemperatureTable'` .
	// - For references, you must specify either variables or input values. For example, the value for the `tableName` parameter can be `$variable.ddbtableName` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `contentExpression` parameter in `Payload` uses a substitution template.
	//
	// `'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'`
	// - For a string concatenation, you must use `+` . A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `tableName` parameter uses a string concatenation.
	//
	// `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	//
	// The value for the `type` parameter in `Payload` must be `JSON` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-dynamodbv2
	//
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-iotevents
	//
	IotEvents interface{} `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise .
	//
	// You must use expressions for all parameters in `IotSiteWiseAction` . The expressions accept literals, operators, functions, references, and substitutions templates.
	//
	// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `propertyAlias` parameter can be `'/company/windfarm/3/turbine/7/temperature'` .
	// - For references, you must specify either variables or input values. For example, the value for the `assetId` parameter can be `$input.TurbineInput.assetId1` .
	// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//
	// In the following example, the value for the `propertyAlias` parameter uses a substitution template.
	//
	// `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`
	//
	// You must specify either `propertyAlias` or both `assetId` and `propertyId` to identify the target asset property in AWS IoT SiteWise .
	//
	// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-iotsitewise
	//
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the AWS IoT message broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-iottopicpublish
	//
	IotTopicPublish interface{} `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Information required to publish the Amazon SNS message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-sns
	//
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmaction.html#cfn-iotevents-alarmmodel-alarmaction-sqs
	//
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
}

