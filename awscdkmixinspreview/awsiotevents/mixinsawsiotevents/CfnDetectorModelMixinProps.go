package mixinsawsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDetectorModelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDetectorModelMixinProps := &CfnDetectorModelMixinProps{
//   	DetectorModelDefinition: &DetectorModelDefinitionProperty{
//   		InitialStateName: jsii.String("initialStateName"),
//   		States: []interface{}{
//   			&StateProperty{
//   				OnEnter: &OnEnterProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   				},
//   				OnExit: &OnExitProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   				},
//   				OnInput: &OnInputProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   						},
//   					},
//   					TransitionEvents: []interface{}{
//   						&TransitionEventProperty{
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyType: jsii.String("hashKeyType"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   										TableName: jsii.String("tableName"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TableName: jsii.String("tableName"),
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//   											},
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//   										},
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										TargetArn: jsii.String("targetArn"),
//   									},
//   									Sqs: &SqsProperty{
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										QueueUrl: jsii.String("queueUrl"),
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   							NextState: jsii.String("nextState"),
//   						},
//   					},
//   				},
//   				StateName: jsii.String("stateName"),
//   			},
//   		},
//   	},
//   	DetectorModelDescription: jsii.String("detectorModelDescription"),
//   	DetectorModelName: jsii.String("detectorModelName"),
//   	EvaluationMethod: jsii.String("evaluationMethod"),
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html
//
type CfnDetectorModelMixinProps struct {
	// Information that defines how a detector operates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodeldefinition
	//
	DetectorModelDefinition interface{} `field:"optional" json:"detectorModelDefinition" yaml:"detectorModelDefinition"`
	// A brief description of the detector model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodeldescription
	//
	DetectorModelDescription *string `field:"optional" json:"detectorModelDescription" yaml:"detectorModelDescription"`
	// The name of the detector model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-detectormodelname
	//
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-evaluationmethod
	//
	EvaluationMethod *string `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-detectormodel.html#cfn-iotevents-detectormodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

