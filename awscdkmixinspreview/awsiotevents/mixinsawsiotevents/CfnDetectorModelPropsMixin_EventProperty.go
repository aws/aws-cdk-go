package mixinsawsiotevents


// Specifies the `actions` to be performed when the `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventProperty := &EventProperty{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ClearTimer: &ClearTimerProperty{
//   				TimerName: jsii.String("timerName"),
//   			},
//   			DynamoDb: &DynamoDBProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyType: jsii.String("hashKeyType"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				Operation: jsii.String("operation"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
//   				TableName: jsii.String("tableName"),
//   			},
//   			DynamoDBv2: &DynamoDBv2Property{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				TableName: jsii.String("tableName"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				Separator: jsii.String("separator"),
//   			},
//   			IotEvents: &IotEventsProperty{
//   				InputName: jsii.String("inputName"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			IotSiteWise: &IotSiteWiseProperty{
//   				AssetId: jsii.String("assetId"),
//   				EntryId: jsii.String("entryId"),
//   				PropertyAlias: jsii.String("propertyAlias"),
//   				PropertyId: jsii.String("propertyId"),
//   				PropertyValue: &AssetPropertyValueProperty{
//   					Quality: jsii.String("quality"),
//   					Timestamp: &AssetPropertyTimestampProperty{
//   						OffsetInNanos: jsii.String("offsetInNanos"),
//   						TimeInSeconds: jsii.String("timeInSeconds"),
//   					},
//   					Value: &AssetPropertyVariantProperty{
//   						BooleanValue: jsii.String("booleanValue"),
//   						DoubleValue: jsii.String("doubleValue"),
//   						IntegerValue: jsii.String("integerValue"),
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			IotTopicPublish: &IotTopicPublishProperty{
//   				MqttTopic: jsii.String("mqttTopic"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Lambda: &LambdaProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			ResetTimer: &ResetTimerProperty{
//   				TimerName: jsii.String("timerName"),
//   			},
//   			SetTimer: &SetTimerProperty{
//   				DurationExpression: jsii.String("durationExpression"),
//   				Seconds: jsii.Number(123),
//   				TimerName: jsii.String("timerName"),
//   			},
//   			SetVariable: &SetVariableProperty{
//   				Value: jsii.String("value"),
//   				VariableName: jsii.String("variableName"),
//   			},
//   			Sns: &SnsProperty{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				TargetArn: jsii.String("targetArn"),
//   			},
//   			Sqs: &SqsProperty{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				QueueUrl: jsii.String("queueUrl"),
//   				UseBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Condition: jsii.String("condition"),
//   	EventName: jsii.String("eventName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html
//
type CfnDetectorModelPropsMixin_EventProperty struct {
	// The actions to be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Optional.
	//
	// The Boolean expression that, when TRUE, causes the `actions` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The name of the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-event.html#cfn-iotevents-detectormodel-event-eventname
	//
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
}

