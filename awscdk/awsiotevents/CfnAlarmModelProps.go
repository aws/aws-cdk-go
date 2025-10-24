package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAlarmModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmModelProps := &CfnAlarmModelProps{
//   	AlarmRule: &AlarmRuleProperty{
//   		SimpleRule: &SimpleRuleProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			InputProperty: jsii.String("inputProperty"),
//   			Threshold: jsii.String("threshold"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AlarmCapabilities: &AlarmCapabilitiesProperty{
//   		AcknowledgeFlow: &AcknowledgeFlowProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		InitializationConfiguration: &InitializationConfigurationProperty{
//   			DisabledOnInitialization: jsii.Boolean(false),
//   		},
//   	},
//   	AlarmEventActions: &AlarmEventActionsProperty{
//   		AlarmActions: []interface{}{
//   			&AlarmActionProperty{
//   				DynamoDb: &DynamoDBProperty{
//   					HashKeyField: jsii.String("hashKeyField"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					HashKeyType: jsii.String("hashKeyType"),
//   					Operation: jsii.String("operation"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				DynamoDBv2: &DynamoDBv2Property{
//   					TableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Firehose: &FirehoseProperty{
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					Separator: jsii.String("separator"),
//   				},
//   				IotEvents: &IotEventsProperty{
//   					InputName: jsii.String("inputName"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				IotSiteWise: &IotSiteWiseProperty{
//   					AssetId: jsii.String("assetId"),
//   					EntryId: jsii.String("entryId"),
//   					PropertyAlias: jsii.String("propertyAlias"),
//   					PropertyId: jsii.String("propertyId"),
//   					PropertyValue: &AssetPropertyValueProperty{
//   						Value: &AssetPropertyVariantProperty{
//   							BooleanValue: jsii.String("booleanValue"),
//   							DoubleValue: jsii.String("doubleValue"),
//   							IntegerValue: jsii.String("integerValue"),
//   							StringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						Quality: jsii.String("quality"),
//   						Timestamp: &AssetPropertyTimestampProperty{
//   							TimeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							OffsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   					},
//   				},
//   				IotTopicPublish: &IotTopicPublishProperty{
//   					MqttTopic: jsii.String("mqttTopic"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Lambda: &LambdaProperty{
//   					FunctionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sns: &SnsProperty{
//   					TargetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sqs: &SqsProperty{
//   					QueueUrl: jsii.String("queueUrl"),
//
//   					// the properties below are optional
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					UseBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	AlarmModelDescription: jsii.String("alarmModelDescription"),
//   	AlarmModelName: jsii.String("alarmModelName"),
//   	Key: jsii.String("key"),
//   	Severity: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html
//
type CfnAlarmModelProps struct {
	// Defines when your alarm is invoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmrule
	//
	AlarmRule interface{} `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of alarm state changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmcapabilities
	//
	AlarmCapabilities interface{} `field:"optional" json:"alarmCapabilities" yaml:"alarmCapabilities"`
	// Contains information about one or more alarm actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmeventactions
	//
	AlarmEventActions interface{} `field:"optional" json:"alarmEventActions" yaml:"alarmEventActions"`
	// The description of the alarm model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmmodeldescription
	//
	AlarmModelDescription *string `field:"optional" json:"alarmModelDescription" yaml:"alarmModelDescription"`
	// The name of the alarm model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmmodelname
	//
	AlarmModelName *string `field:"optional" json:"alarmModelName" yaml:"alarmModelName"`
	// An input attribute used as a key to create an alarm.
	//
	// AWS IoT Events routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A non-negative integer that reflects the severity level of the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-severity
	//
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// A list of key-value pairs that contain metadata for the alarm model.
	//
	// The tags help you manage the alarm model. For more information, see [Tagging your AWS IoT Events resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *AWS IoT Events Developer Guide* .
	//
	// You can create up to 50 tags for one alarm model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

