package mixinsawsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAlarmModelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlarmModelMixinProps := &CfnAlarmModelMixinProps{
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
//   					HashKeyType: jsii.String("hashKeyType"),
//   					HashKeyValue: jsii.String("hashKeyValue"),
//   					Operation: jsii.String("operation"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					PayloadField: jsii.String("payloadField"),
//   					RangeKeyField: jsii.String("rangeKeyField"),
//   					RangeKeyType: jsii.String("rangeKeyType"),
//   					RangeKeyValue: jsii.String("rangeKeyValue"),
//   					TableName: jsii.String("tableName"),
//   				},
//   				DynamoDBv2: &DynamoDBv2Property{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					TableName: jsii.String("tableName"),
//   				},
//   				Firehose: &FirehoseProperty{
//   					DeliveryStreamName: jsii.String("deliveryStreamName"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					Separator: jsii.String("separator"),
//   				},
//   				IotEvents: &IotEventsProperty{
//   					InputName: jsii.String("inputName"),
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
//   						Quality: jsii.String("quality"),
//   						Timestamp: &AssetPropertyTimestampProperty{
//   							OffsetInNanos: jsii.String("offsetInNanos"),
//   							TimeInSeconds: jsii.String("timeInSeconds"),
//   						},
//   						Value: &AssetPropertyVariantProperty{
//   							BooleanValue: jsii.String("booleanValue"),
//   							DoubleValue: jsii.String("doubleValue"),
//   							IntegerValue: jsii.String("integerValue"),
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   				IotTopicPublish: &IotTopicPublishProperty{
//   					MqttTopic: jsii.String("mqttTopic"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Lambda: &LambdaProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Sns: &SnsProperty{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					TargetArn: jsii.String("targetArn"),
//   				},
//   				Sqs: &SqsProperty{
//   					Payload: &PayloadProperty{
//   						ContentExpression: jsii.String("contentExpression"),
//   						Type: jsii.String("type"),
//   					},
//   					QueueUrl: jsii.String("queueUrl"),
//   					UseBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	AlarmModelDescription: jsii.String("alarmModelDescription"),
//   	AlarmModelName: jsii.String("alarmModelName"),
//   	AlarmRule: &AlarmRuleProperty{
//   		SimpleRule: &SimpleRuleProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			InputProperty: jsii.String("inputProperty"),
//   			Threshold: jsii.String("threshold"),
//   		},
//   	},
//   	Key: jsii.String("key"),
//   	RoleArn: jsii.String("roleArn"),
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
type CfnAlarmModelMixinProps struct {
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
	// Defines when your alarm is invoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-alarmrule
	//
	AlarmRule interface{} `field:"optional" json:"alarmRule" yaml:"alarmRule"`
	// An input attribute used as a key to create an alarm.
	//
	// AWS IoT Events routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-alarmmodel.html#cfn-iotevents-alarmmodel-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
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

