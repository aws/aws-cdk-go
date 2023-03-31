package awsiotevents


// Specifies one of the following actions to receive notifications when the alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmActionProperty := &alarmActionProperty{
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
//   		assetId: jsii.String("assetId"),
//   		entryId: jsii.String("entryId"),
//   		propertyAlias: jsii.String("propertyAlias"),
//   		propertyId: jsii.String("propertyId"),
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
type CfnAlarmModel_AlarmActionProperty struct {
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
	DynamoDBv2 interface{} `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
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
	IotSiteWise interface{} `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the AWS IoT message broker.
	IotTopicPublish interface{} `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// `CfnAlarmModel.AlarmActionProperty.Sns`.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
	Sqs interface{} `field:"optional" json:"sqs" yaml:"sqs"`
}

