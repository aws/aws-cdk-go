package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDetectorModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorModelProps := &CfnDetectorModelProps{
//   	DetectorModelDefinition: &DetectorModelDefinitionProperty{
//   		InitialStateName: jsii.String("initialStateName"),
//   		States: []interface{}{
//   			&StateProperty{
//   				StateName: jsii.String("stateName"),
//
//   				// the properties below are optional
//   				OnEnter: &OnEnterProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				OnExit: &OnExitProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				OnInput: &OnInputProperty{
//   					Events: []interface{}{
//   						&EventProperty{
//   							EventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							Condition: jsii.String("condition"),
//   						},
//   					},
//   					TransitionEvents: []interface{}{
//   						&TransitionEventProperty{
//   							Condition: jsii.String("condition"),
//   							EventName: jsii.String("eventName"),
//   							NextState: jsii.String("nextState"),
//
//   							// the properties below are optional
//   							Actions: []interface{}{
//   								&ActionProperty{
//   									ClearTimer: &ClearTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									DynamoDb: &DynamoDBProperty{
//   										HashKeyField: jsii.String("hashKeyField"),
//   										HashKeyValue: jsii.String("hashKeyValue"),
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										HashKeyType: jsii.String("hashKeyType"),
//   										Operation: jsii.String("operation"),
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										PayloadField: jsii.String("payloadField"),
//   										RangeKeyField: jsii.String("rangeKeyField"),
//   										RangeKeyType: jsii.String("rangeKeyType"),
//   										RangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									DynamoDBv2: &DynamoDBv2Property{
//   										TableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Firehose: &FirehoseProperty{
//   										DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										Separator: jsii.String("separator"),
//   									},
//   									IotEvents: &IotEventsProperty{
//   										InputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									IotSiteWise: &IotSiteWiseProperty{
//   										PropertyValue: &AssetPropertyValueProperty{
//   											Value: &AssetPropertyVariantProperty{
//   												BooleanValue: jsii.String("booleanValue"),
//   												DoubleValue: jsii.String("doubleValue"),
//   												IntegerValue: jsii.String("integerValue"),
//   												StringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											Quality: jsii.String("quality"),
//   											Timestamp: &AssetPropertyTimestampProperty{
//   												TimeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												OffsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										AssetId: jsii.String("assetId"),
//   										EntryId: jsii.String("entryId"),
//   										PropertyAlias: jsii.String("propertyAlias"),
//   										PropertyId: jsii.String("propertyId"),
//   									},
//   									IotTopicPublish: &IotTopicPublishProperty{
//   										MqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Lambda: &LambdaProperty{
//   										FunctionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									ResetTimer: &ResetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//   									},
//   									SetTimer: &SetTimerProperty{
//   										TimerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										DurationExpression: jsii.String("durationExpression"),
//   										Seconds: jsii.Number(123),
//   									},
//   									SetVariable: &SetVariableProperty{
//   										Value: jsii.String("value"),
//   										VariableName: jsii.String("variableName"),
//   									},
//   									Sns: &SnsProperty{
//   										TargetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   									},
//   									Sqs: &SqsProperty{
//   										QueueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										Payload: &PayloadProperty{
//   											ContentExpression: jsii.String("contentExpression"),
//   											Type: jsii.String("type"),
//   										},
//   										UseBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	DetectorModelDescription: jsii.String("detectorModelDescription"),
//   	DetectorModelName: jsii.String("detectorModelName"),
//   	EvaluationMethod: jsii.String("evaluationMethod"),
//   	Key: jsii.String("key"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetectorModelProps struct {
	// Information that defines how a detector operates.
	DetectorModelDefinition interface{} `field:"required" json:"detectorModelDefinition" yaml:"detectorModelDefinition"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief description of the detector model.
	DetectorModelDescription *string `field:"optional" json:"detectorModelDescription" yaml:"detectorModelDescription"`
	// The name of the detector model.
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	EvaluationMethod *string `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

