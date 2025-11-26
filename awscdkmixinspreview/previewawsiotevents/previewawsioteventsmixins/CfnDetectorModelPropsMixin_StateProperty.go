package previewawsioteventsmixins


// Information that defines a state of a detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stateProperty := &StateProperty{
//   	OnEnter: &OnEnterProperty{
//   		Events: []interface{}{
//   			&EventProperty{
//   				Actions: []interface{}{
//   					&ActionProperty{
//   						ClearTimer: &ClearTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						DynamoDb: &DynamoDBProperty{
//   							HashKeyField: jsii.String("hashKeyField"),
//   							HashKeyType: jsii.String("hashKeyType"),
//   							HashKeyValue: jsii.String("hashKeyValue"),
//   							Operation: jsii.String("operation"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							PayloadField: jsii.String("payloadField"),
//   							RangeKeyField: jsii.String("rangeKeyField"),
//   							RangeKeyType: jsii.String("rangeKeyType"),
//   							RangeKeyValue: jsii.String("rangeKeyValue"),
//   							TableName: jsii.String("tableName"),
//   						},
//   						DynamoDBv2: &DynamoDBv2Property{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TableName: jsii.String("tableName"),
//   						},
//   						Firehose: &FirehoseProperty{
//   							DeliveryStreamName: jsii.String("deliveryStreamName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							Separator: jsii.String("separator"),
//   						},
//   						IotEvents: &IotEventsProperty{
//   							InputName: jsii.String("inputName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						IotSiteWise: &IotSiteWiseProperty{
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   							PropertyValue: &AssetPropertyValueProperty{
//   								Quality: jsii.String("quality"),
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//   							},
//   						},
//   						IotTopicPublish: &IotTopicPublishProperty{
//   							MqttTopic: jsii.String("mqttTopic"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						Lambda: &LambdaProperty{
//   							FunctionArn: jsii.String("functionArn"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ResetTimer: &ResetTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetTimer: &SetTimerProperty{
//   							DurationExpression: jsii.String("durationExpression"),
//   							Seconds: jsii.Number(123),
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetVariable: &SetVariableProperty{
//   							Value: jsii.String("value"),
//   							VariableName: jsii.String("variableName"),
//   						},
//   						Sns: &SnsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TargetArn: jsii.String("targetArn"),
//   						},
//   						Sqs: &SqsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							QueueUrl: jsii.String("queueUrl"),
//   							UseBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				Condition: jsii.String("condition"),
//   				EventName: jsii.String("eventName"),
//   			},
//   		},
//   	},
//   	OnExit: &OnExitProperty{
//   		Events: []interface{}{
//   			&EventProperty{
//   				Actions: []interface{}{
//   					&ActionProperty{
//   						ClearTimer: &ClearTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						DynamoDb: &DynamoDBProperty{
//   							HashKeyField: jsii.String("hashKeyField"),
//   							HashKeyType: jsii.String("hashKeyType"),
//   							HashKeyValue: jsii.String("hashKeyValue"),
//   							Operation: jsii.String("operation"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							PayloadField: jsii.String("payloadField"),
//   							RangeKeyField: jsii.String("rangeKeyField"),
//   							RangeKeyType: jsii.String("rangeKeyType"),
//   							RangeKeyValue: jsii.String("rangeKeyValue"),
//   							TableName: jsii.String("tableName"),
//   						},
//   						DynamoDBv2: &DynamoDBv2Property{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TableName: jsii.String("tableName"),
//   						},
//   						Firehose: &FirehoseProperty{
//   							DeliveryStreamName: jsii.String("deliveryStreamName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							Separator: jsii.String("separator"),
//   						},
//   						IotEvents: &IotEventsProperty{
//   							InputName: jsii.String("inputName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						IotSiteWise: &IotSiteWiseProperty{
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   							PropertyValue: &AssetPropertyValueProperty{
//   								Quality: jsii.String("quality"),
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//   							},
//   						},
//   						IotTopicPublish: &IotTopicPublishProperty{
//   							MqttTopic: jsii.String("mqttTopic"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						Lambda: &LambdaProperty{
//   							FunctionArn: jsii.String("functionArn"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ResetTimer: &ResetTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetTimer: &SetTimerProperty{
//   							DurationExpression: jsii.String("durationExpression"),
//   							Seconds: jsii.Number(123),
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetVariable: &SetVariableProperty{
//   							Value: jsii.String("value"),
//   							VariableName: jsii.String("variableName"),
//   						},
//   						Sns: &SnsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TargetArn: jsii.String("targetArn"),
//   						},
//   						Sqs: &SqsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							QueueUrl: jsii.String("queueUrl"),
//   							UseBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				Condition: jsii.String("condition"),
//   				EventName: jsii.String("eventName"),
//   			},
//   		},
//   	},
//   	OnInput: &OnInputProperty{
//   		Events: []interface{}{
//   			&EventProperty{
//   				Actions: []interface{}{
//   					&ActionProperty{
//   						ClearTimer: &ClearTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						DynamoDb: &DynamoDBProperty{
//   							HashKeyField: jsii.String("hashKeyField"),
//   							HashKeyType: jsii.String("hashKeyType"),
//   							HashKeyValue: jsii.String("hashKeyValue"),
//   							Operation: jsii.String("operation"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							PayloadField: jsii.String("payloadField"),
//   							RangeKeyField: jsii.String("rangeKeyField"),
//   							RangeKeyType: jsii.String("rangeKeyType"),
//   							RangeKeyValue: jsii.String("rangeKeyValue"),
//   							TableName: jsii.String("tableName"),
//   						},
//   						DynamoDBv2: &DynamoDBv2Property{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TableName: jsii.String("tableName"),
//   						},
//   						Firehose: &FirehoseProperty{
//   							DeliveryStreamName: jsii.String("deliveryStreamName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							Separator: jsii.String("separator"),
//   						},
//   						IotEvents: &IotEventsProperty{
//   							InputName: jsii.String("inputName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						IotSiteWise: &IotSiteWiseProperty{
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   							PropertyValue: &AssetPropertyValueProperty{
//   								Quality: jsii.String("quality"),
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//   							},
//   						},
//   						IotTopicPublish: &IotTopicPublishProperty{
//   							MqttTopic: jsii.String("mqttTopic"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						Lambda: &LambdaProperty{
//   							FunctionArn: jsii.String("functionArn"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ResetTimer: &ResetTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetTimer: &SetTimerProperty{
//   							DurationExpression: jsii.String("durationExpression"),
//   							Seconds: jsii.Number(123),
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetVariable: &SetVariableProperty{
//   							Value: jsii.String("value"),
//   							VariableName: jsii.String("variableName"),
//   						},
//   						Sns: &SnsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TargetArn: jsii.String("targetArn"),
//   						},
//   						Sqs: &SqsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							QueueUrl: jsii.String("queueUrl"),
//   							UseBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				Condition: jsii.String("condition"),
//   				EventName: jsii.String("eventName"),
//   			},
//   		},
//   		TransitionEvents: []interface{}{
//   			&TransitionEventProperty{
//   				Actions: []interface{}{
//   					&ActionProperty{
//   						ClearTimer: &ClearTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						DynamoDb: &DynamoDBProperty{
//   							HashKeyField: jsii.String("hashKeyField"),
//   							HashKeyType: jsii.String("hashKeyType"),
//   							HashKeyValue: jsii.String("hashKeyValue"),
//   							Operation: jsii.String("operation"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							PayloadField: jsii.String("payloadField"),
//   							RangeKeyField: jsii.String("rangeKeyField"),
//   							RangeKeyType: jsii.String("rangeKeyType"),
//   							RangeKeyValue: jsii.String("rangeKeyValue"),
//   							TableName: jsii.String("tableName"),
//   						},
//   						DynamoDBv2: &DynamoDBv2Property{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TableName: jsii.String("tableName"),
//   						},
//   						Firehose: &FirehoseProperty{
//   							DeliveryStreamName: jsii.String("deliveryStreamName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							Separator: jsii.String("separator"),
//   						},
//   						IotEvents: &IotEventsProperty{
//   							InputName: jsii.String("inputName"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						IotSiteWise: &IotSiteWiseProperty{
//   							AssetId: jsii.String("assetId"),
//   							EntryId: jsii.String("entryId"),
//   							PropertyAlias: jsii.String("propertyAlias"),
//   							PropertyId: jsii.String("propertyId"),
//   							PropertyValue: &AssetPropertyValueProperty{
//   								Quality: jsii.String("quality"),
//   								Timestamp: &AssetPropertyTimestampProperty{
//   									OffsetInNanos: jsii.String("offsetInNanos"),
//   									TimeInSeconds: jsii.String("timeInSeconds"),
//   								},
//   								Value: &AssetPropertyVariantProperty{
//   									BooleanValue: jsii.String("booleanValue"),
//   									DoubleValue: jsii.String("doubleValue"),
//   									IntegerValue: jsii.String("integerValue"),
//   									StringValue: jsii.String("stringValue"),
//   								},
//   							},
//   						},
//   						IotTopicPublish: &IotTopicPublishProperty{
//   							MqttTopic: jsii.String("mqttTopic"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						Lambda: &LambdaProperty{
//   							FunctionArn: jsii.String("functionArn"),
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   						ResetTimer: &ResetTimerProperty{
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetTimer: &SetTimerProperty{
//   							DurationExpression: jsii.String("durationExpression"),
//   							Seconds: jsii.Number(123),
//   							TimerName: jsii.String("timerName"),
//   						},
//   						SetVariable: &SetVariableProperty{
//   							Value: jsii.String("value"),
//   							VariableName: jsii.String("variableName"),
//   						},
//   						Sns: &SnsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							TargetArn: jsii.String("targetArn"),
//   						},
//   						Sqs: &SqsProperty{
//   							Payload: &PayloadProperty{
//   								ContentExpression: jsii.String("contentExpression"),
//   								Type: jsii.String("type"),
//   							},
//   							QueueUrl: jsii.String("queueUrl"),
//   							UseBase64: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				Condition: jsii.String("condition"),
//   				EventName: jsii.String("eventName"),
//   				NextState: jsii.String("nextState"),
//   			},
//   		},
//   	},
//   	StateName: jsii.String("stateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html
//
type CfnDetectorModelPropsMixin_StateProperty struct {
	// When entering this state, perform these `actions` if the `condition` is TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onenter
	//
	OnEnter interface{} `field:"optional" json:"onEnter" yaml:"onEnter"`
	// When exiting this state, perform these `actions` if the specified `condition` is `TRUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onexit
	//
	OnExit interface{} `field:"optional" json:"onExit" yaml:"onExit"`
	// When an input is received and the `condition` is TRUE, perform the specified `actions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-oninput
	//
	OnInput interface{} `field:"optional" json:"onInput" yaml:"onInput"`
	// The name of the state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-statename
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
}

