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
//   cfnDetectorModelProps := &cfnDetectorModelProps{
//   	detectorModelDefinition: &detectorModelDefinitionProperty{
//   		initialStateName: jsii.String("initialStateName"),
//   		states: []interface{}{
//   			&stateProperty{
//   				stateName: jsii.String("stateName"),
//
//   				// the properties below are optional
//   				onEnter: &onEnterProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onExit: &onExitProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   				},
//   				onInput: &onInputProperty{
//   					events: []interface{}{
//   						&eventProperty{
//   							eventName: jsii.String("eventName"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							condition: jsii.String("condition"),
//   						},
//   					},
//   					transitionEvents: []interface{}{
//   						&transitionEventProperty{
//   							condition: jsii.String("condition"),
//   							eventName: jsii.String("eventName"),
//   							nextState: jsii.String("nextState"),
//
//   							// the properties below are optional
//   							actions: []interface{}{
//   								&actionProperty{
//   									clearTimer: &clearTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									dynamoDb: &dynamoDBProperty{
//   										hashKeyField: jsii.String("hashKeyField"),
//   										hashKeyValue: jsii.String("hashKeyValue"),
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										hashKeyType: jsii.String("hashKeyType"),
//   										operation: jsii.String("operation"),
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										payloadField: jsii.String("payloadField"),
//   										rangeKeyField: jsii.String("rangeKeyField"),
//   										rangeKeyType: jsii.String("rangeKeyType"),
//   										rangeKeyValue: jsii.String("rangeKeyValue"),
//   									},
//   									dynamoDBv2: &dynamoDBv2Property{
//   										tableName: jsii.String("tableName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									firehose: &firehoseProperty{
//   										deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										separator: jsii.String("separator"),
//   									},
//   									iotEvents: &iotEventsProperty{
//   										inputName: jsii.String("inputName"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									iotSiteWise: &iotSiteWiseProperty{
//   										propertyValue: &assetPropertyValueProperty{
//   											value: &assetPropertyVariantProperty{
//   												booleanValue: jsii.String("booleanValue"),
//   												doubleValue: jsii.String("doubleValue"),
//   												integerValue: jsii.String("integerValue"),
//   												stringValue: jsii.String("stringValue"),
//   											},
//
//   											// the properties below are optional
//   											quality: jsii.String("quality"),
//   											timestamp: &assetPropertyTimestampProperty{
//   												timeInSeconds: jsii.String("timeInSeconds"),
//
//   												// the properties below are optional
//   												offsetInNanos: jsii.String("offsetInNanos"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										assetId: jsii.String("assetId"),
//   										entryId: jsii.String("entryId"),
//   										propertyAlias: jsii.String("propertyAlias"),
//   										propertyId: jsii.String("propertyId"),
//   									},
//   									iotTopicPublish: &iotTopicPublishProperty{
//   										mqttTopic: jsii.String("mqttTopic"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									lambda: &lambdaProperty{
//   										functionArn: jsii.String("functionArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									resetTimer: &resetTimerProperty{
//   										timerName: jsii.String("timerName"),
//   									},
//   									setTimer: &setTimerProperty{
//   										timerName: jsii.String("timerName"),
//
//   										// the properties below are optional
//   										durationExpression: jsii.String("durationExpression"),
//   										seconds: jsii.Number(123),
//   									},
//   									setVariable: &setVariableProperty{
//   										value: jsii.String("value"),
//   										variableName: jsii.String("variableName"),
//   									},
//   									sns: &snsProperty{
//   										targetArn: jsii.String("targetArn"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   									},
//   									sqs: &sqsProperty{
//   										queueUrl: jsii.String("queueUrl"),
//
//   										// the properties below are optional
//   										payload: &payloadProperty{
//   											contentExpression: jsii.String("contentExpression"),
//   											type: jsii.String("type"),
//   										},
//   										useBase64: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	detectorModelDescription: jsii.String("detectorModelDescription"),
//   	detectorModelName: jsii.String("detectorModelName"),
//   	evaluationMethod: jsii.String("evaluationMethod"),
//   	key: jsii.String("key"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

