package awsiotevents


// Specifies the actions performed and the next state entered when a `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitionEventProperty := &TransitionEventProperty{
//   	Condition: jsii.String("condition"),
//   	EventName: jsii.String("eventName"),
//   	NextState: jsii.String("nextState"),
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ClearTimer: &ClearTimerProperty{
//   				TimerName: jsii.String("timerName"),
//   			},
//   			DynamoDb: &DynamoDBProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				HashKeyType: jsii.String("hashKeyType"),
//   				Operation: jsii.String("operation"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			DynamoDBv2: &DynamoDBv2Property{
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Firehose: &FirehoseProperty{
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				Separator: jsii.String("separator"),
//   			},
//   			IotEvents: &IotEventsProperty{
//   				InputName: jsii.String("inputName"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			IotSiteWise: &IotSiteWiseProperty{
//   				PropertyValue: &AssetPropertyValueProperty{
//   					Value: &AssetPropertyVariantProperty{
//   						BooleanValue: jsii.String("booleanValue"),
//   						DoubleValue: jsii.String("doubleValue"),
//   						IntegerValue: jsii.String("integerValue"),
//   						StringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					Quality: jsii.String("quality"),
//   					Timestamp: &AssetPropertyTimestampProperty{
//   						TimeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						OffsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				AssetId: jsii.String("assetId"),
//   				EntryId: jsii.String("entryId"),
//   				PropertyAlias: jsii.String("propertyAlias"),
//   				PropertyId: jsii.String("propertyId"),
//   			},
//   			IotTopicPublish: &IotTopicPublishProperty{
//   				MqttTopic: jsii.String("mqttTopic"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Lambda: &LambdaProperty{
//   				FunctionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			ResetTimer: &ResetTimerProperty{
//   				TimerName: jsii.String("timerName"),
//   			},
//   			SetTimer: &SetTimerProperty{
//   				TimerName: jsii.String("timerName"),
//
//   				// the properties below are optional
//   				DurationExpression: jsii.String("durationExpression"),
//   				Seconds: jsii.Number(123),
//   			},
//   			SetVariable: &SetVariableProperty{
//   				Value: jsii.String("value"),
//   				VariableName: jsii.String("variableName"),
//   			},
//   			Sns: &SnsProperty{
//   				TargetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Sqs: &SqsProperty{
//   				QueueUrl: jsii.String("queueUrl"),
//
//   				// the properties below are optional
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				UseBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html
//
type CfnDetectorModel_TransitionEventProperty struct {
	// Required.
	//
	// A Boolean expression that when TRUE causes the actions to be performed and the `nextState` to be entered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-condition
	//
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// The name of the transition event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-eventname
	//
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The next state to enter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-nextstate
	//
	NextState *string `field:"required" json:"nextState" yaml:"nextState"`
	// The actions to be performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
}

