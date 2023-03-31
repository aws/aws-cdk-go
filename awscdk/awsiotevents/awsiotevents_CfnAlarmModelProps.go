package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAlarmModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmModelProps := &cfnAlarmModelProps{
//   	alarmRule: &alarmRuleProperty{
//   		simpleRule: &simpleRuleProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			inputProperty: jsii.String("inputProperty"),
//   			threshold: jsii.String("threshold"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	alarmCapabilities: &alarmCapabilitiesProperty{
//   		acknowledgeFlow: &acknowledgeFlowProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		initializationConfiguration: &initializationConfigurationProperty{
//   			disabledOnInitialization: jsii.Boolean(false),
//   		},
//   	},
//   	alarmEventActions: &alarmEventActionsProperty{
//   		alarmActions: []interface{}{
//   			&alarmActionProperty{
//   				dynamoDb: &dynamoDBProperty{
//   					hashKeyField: jsii.String("hashKeyField"),
//   					hashKeyValue: jsii.String("hashKeyValue"),
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					hashKeyType: jsii.String("hashKeyType"),
//   					operation: jsii.String("operation"),
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					payloadField: jsii.String("payloadField"),
//   					rangeKeyField: jsii.String("rangeKeyField"),
//   					rangeKeyType: jsii.String("rangeKeyType"),
//   					rangeKeyValue: jsii.String("rangeKeyValue"),
//   				},
//   				dynamoDBv2: &dynamoDBv2Property{
//   					tableName: jsii.String("tableName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				firehose: &firehoseProperty{
//   					deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					separator: jsii.String("separator"),
//   				},
//   				iotEvents: &iotEventsProperty{
//   					inputName: jsii.String("inputName"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				iotSiteWise: &iotSiteWiseProperty{
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   					propertyValue: &assetPropertyValueProperty{
//   						value: &assetPropertyVariantProperty{
//   							booleanValue: jsii.String("booleanValue"),
//   							doubleValue: jsii.String("doubleValue"),
//   							integerValue: jsii.String("integerValue"),
//   							stringValue: jsii.String("stringValue"),
//   						},
//
//   						// the properties below are optional
//   						quality: jsii.String("quality"),
//   						timestamp: &assetPropertyTimestampProperty{
//   							timeInSeconds: jsii.String("timeInSeconds"),
//
//   							// the properties below are optional
//   							offsetInNanos: jsii.String("offsetInNanos"),
//   						},
//   					},
//   				},
//   				iotTopicPublish: &iotTopicPublishProperty{
//   					mqttTopic: jsii.String("mqttTopic"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				lambda: &lambdaProperty{
//   					functionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sns: &snsProperty{
//   					targetArn: jsii.String("targetArn"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				sqs: &sqsProperty{
//   					queueUrl: jsii.String("queueUrl"),
//
//   					// the properties below are optional
//   					payload: &payloadProperty{
//   						contentExpression: jsii.String("contentExpression"),
//   						type: jsii.String("type"),
//   					},
//   					useBase64: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	alarmModelDescription: jsii.String("alarmModelDescription"),
//   	alarmModelName: jsii.String("alarmModelName"),
//   	key: jsii.String("key"),
//   	severity: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAlarmModelProps struct {
	// Defines when your alarm is invoked.
	AlarmRule interface{} `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources.
	//
	// For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of alarm state changes.
	AlarmCapabilities interface{} `field:"optional" json:"alarmCapabilities" yaml:"alarmCapabilities"`
	// Contains information about one or more alarm actions.
	AlarmEventActions interface{} `field:"optional" json:"alarmEventActions" yaml:"alarmEventActions"`
	// The description of the alarm model.
	AlarmModelDescription *string `field:"optional" json:"alarmModelDescription" yaml:"alarmModelDescription"`
	// The name of the alarm model.
	AlarmModelName *string `field:"optional" json:"alarmModelName" yaml:"alarmModelName"`
	// An input attribute used as a key to create an alarm.
	//
	// AWS IoT Events routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A non-negative integer that reflects the severity level of the alarm.
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// A list of key-value pairs that contain metadata for the alarm model.
	//
	// The tags help you manage the alarm model. For more information, see [Tagging your AWS IoT Events resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *AWS IoT Events Developer Guide* .
	//
	// You can create up to 50 tags for one alarm model.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

