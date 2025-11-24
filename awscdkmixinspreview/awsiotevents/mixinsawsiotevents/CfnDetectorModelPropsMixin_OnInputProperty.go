package mixinsawsiotevents


// Specifies the actions performed when the `condition` evaluates to TRUE.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   onInputProperty := &OnInputProperty{
//   	Events: []interface{}{
//   		&EventProperty{
//   			Actions: []interface{}{
//   				&ActionProperty{
//   					ClearTimer: &ClearTimerProperty{
//   						TimerName: jsii.String("timerName"),
//   					},
//   					DynamoDb: &DynamoDBProperty{
//   						HashKeyField: jsii.String("hashKeyField"),
//   						HashKeyType: jsii.String("hashKeyType"),
//   						HashKeyValue: jsii.String("hashKeyValue"),
//   						Operation: jsii.String("operation"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						PayloadField: jsii.String("payloadField"),
//   						RangeKeyField: jsii.String("rangeKeyField"),
//   						RangeKeyType: jsii.String("rangeKeyType"),
//   						RangeKeyValue: jsii.String("rangeKeyValue"),
//   						TableName: jsii.String("tableName"),
//   					},
//   					DynamoDBv2: &DynamoDBv2Property{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						TableName: jsii.String("tableName"),
//   					},
//   					Firehose: &FirehoseProperty{
//   						DeliveryStreamName: jsii.String("deliveryStreamName"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						Separator: jsii.String("separator"),
//   					},
//   					IotEvents: &IotEventsProperty{
//   						InputName: jsii.String("inputName"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					IotSiteWise: &IotSiteWiseProperty{
//   						AssetId: jsii.String("assetId"),
//   						EntryId: jsii.String("entryId"),
//   						PropertyAlias: jsii.String("propertyAlias"),
//   						PropertyId: jsii.String("propertyId"),
//   						PropertyValue: &AssetPropertyValueProperty{
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
//   					IotTopicPublish: &IotTopicPublishProperty{
//   						MqttTopic: jsii.String("mqttTopic"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Lambda: &LambdaProperty{
//   						FunctionArn: jsii.String("functionArn"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					ResetTimer: &ResetTimerProperty{
//   						TimerName: jsii.String("timerName"),
//   					},
//   					SetTimer: &SetTimerProperty{
//   						DurationExpression: jsii.String("durationExpression"),
//   						Seconds: jsii.Number(123),
//   						TimerName: jsii.String("timerName"),
//   					},
//   					SetVariable: &SetVariableProperty{
//   						Value: jsii.String("value"),
//   						VariableName: jsii.String("variableName"),
//   					},
//   					Sns: &SnsProperty{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						TargetArn: jsii.String("targetArn"),
//   					},
//   					Sqs: &SqsProperty{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						QueueUrl: jsii.String("queueUrl"),
//   						UseBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Condition: jsii.String("condition"),
//   			EventName: jsii.String("eventName"),
//   		},
//   	},
//   	TransitionEvents: []interface{}{
//   		&TransitionEventProperty{
//   			Actions: []interface{}{
//   				&ActionProperty{
//   					ClearTimer: &ClearTimerProperty{
//   						TimerName: jsii.String("timerName"),
//   					},
//   					DynamoDb: &DynamoDBProperty{
//   						HashKeyField: jsii.String("hashKeyField"),
//   						HashKeyType: jsii.String("hashKeyType"),
//   						HashKeyValue: jsii.String("hashKeyValue"),
//   						Operation: jsii.String("operation"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						PayloadField: jsii.String("payloadField"),
//   						RangeKeyField: jsii.String("rangeKeyField"),
//   						RangeKeyType: jsii.String("rangeKeyType"),
//   						RangeKeyValue: jsii.String("rangeKeyValue"),
//   						TableName: jsii.String("tableName"),
//   					},
//   					DynamoDBv2: &DynamoDBv2Property{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						TableName: jsii.String("tableName"),
//   					},
//   					Firehose: &FirehoseProperty{
//   						DeliveryStreamName: jsii.String("deliveryStreamName"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						Separator: jsii.String("separator"),
//   					},
//   					IotEvents: &IotEventsProperty{
//   						InputName: jsii.String("inputName"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					IotSiteWise: &IotSiteWiseProperty{
//   						AssetId: jsii.String("assetId"),
//   						EntryId: jsii.String("entryId"),
//   						PropertyAlias: jsii.String("propertyAlias"),
//   						PropertyId: jsii.String("propertyId"),
//   						PropertyValue: &AssetPropertyValueProperty{
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
//   					IotTopicPublish: &IotTopicPublishProperty{
//   						MqttTopic: jsii.String("mqttTopic"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Lambda: &LambdaProperty{
//   						FunctionArn: jsii.String("functionArn"),
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					ResetTimer: &ResetTimerProperty{
//   						TimerName: jsii.String("timerName"),
//   					},
//   					SetTimer: &SetTimerProperty{
//   						DurationExpression: jsii.String("durationExpression"),
//   						Seconds: jsii.Number(123),
//   						TimerName: jsii.String("timerName"),
//   					},
//   					SetVariable: &SetVariableProperty{
//   						Value: jsii.String("value"),
//   						VariableName: jsii.String("variableName"),
//   					},
//   					Sns: &SnsProperty{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						TargetArn: jsii.String("targetArn"),
//   					},
//   					Sqs: &SqsProperty{
//   						Payload: &PayloadProperty{
//   							ContentExpression: jsii.String("contentExpression"),
//   							Type: jsii.String("type"),
//   						},
//   						QueueUrl: jsii.String("queueUrl"),
//   						UseBase64: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Condition: jsii.String("condition"),
//   			EventName: jsii.String("eventName"),
//   			NextState: jsii.String("nextState"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-oninput.html
//
type CfnDetectorModelPropsMixin_OnInputProperty struct {
	// Specifies the actions performed when the `condition` evaluates to TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-oninput.html#cfn-iotevents-detectormodel-oninput-events
	//
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// Specifies the actions performed, and the next state entered, when a `condition` evaluates to TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-oninput.html#cfn-iotevents-detectormodel-oninput-transitionevents
	//
	TransitionEvents interface{} `field:"optional" json:"transitionEvents" yaml:"transitionEvents"`
}

