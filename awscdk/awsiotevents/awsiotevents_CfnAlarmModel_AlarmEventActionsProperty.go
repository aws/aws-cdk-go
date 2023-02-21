package awsiotevents


// Contains information about one or more alarm actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmEventActionsProperty := &alarmEventActionsProperty{
//   	alarmActions: []interface{}{
//   		&alarmActionProperty{
//   			dynamoDb: &dynamoDBProperty{
//   				hashKeyField: jsii.String("hashKeyField"),
//   				hashKeyValue: jsii.String("hashKeyValue"),
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				hashKeyType: jsii.String("hashKeyType"),
//   				operation: jsii.String("operation"),
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				payloadField: jsii.String("payloadField"),
//   				rangeKeyField: jsii.String("rangeKeyField"),
//   				rangeKeyType: jsii.String("rangeKeyType"),
//   				rangeKeyValue: jsii.String("rangeKeyValue"),
//   			},
//   			dynamoDBv2: &dynamoDBv2Property{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			firehose: &firehoseProperty{
//   				deliveryStreamName: jsii.String("deliveryStreamName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				separator: jsii.String("separator"),
//   			},
//   			iotEvents: &iotEventsProperty{
//   				inputName: jsii.String("inputName"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			iotSiteWise: &iotSiteWiseProperty{
//   				assetId: jsii.String("assetId"),
//   				entryId: jsii.String("entryId"),
//   				propertyAlias: jsii.String("propertyAlias"),
//   				propertyId: jsii.String("propertyId"),
//   				propertyValue: &assetPropertyValueProperty{
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   				},
//   			},
//   			iotTopicPublish: &iotTopicPublishProperty{
//   				mqttTopic: jsii.String("mqttTopic"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			lambda: &lambdaProperty{
//   				functionArn: jsii.String("functionArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sns: &snsProperty{
//   				targetArn: jsii.String("targetArn"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			sqs: &sqsProperty{
//   				queueUrl: jsii.String("queueUrl"),
//
//   				// the properties below are optional
//   				payload: &payloadProperty{
//   					contentExpression: jsii.String("contentExpression"),
//   					type: jsii.String("type"),
//   				},
//   				useBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnAlarmModel_AlarmEventActionsProperty struct {
	// Specifies one or more supported actions to receive notifications when the alarm state changes.
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
}

