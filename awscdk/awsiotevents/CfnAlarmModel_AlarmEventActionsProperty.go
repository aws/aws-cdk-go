package awsiotevents


// Contains information about one or more alarm actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmEventActionsProperty := &AlarmEventActionsProperty{
//   	AlarmActions: []interface{}{
//   		&AlarmActionProperty{
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
//   				AssetId: jsii.String("assetId"),
//   				EntryId: jsii.String("entryId"),
//   				PropertyAlias: jsii.String("propertyAlias"),
//   				PropertyId: jsii.String("propertyId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmeventactions.html
//
type CfnAlarmModel_AlarmEventActionsProperty struct {
	// Specifies one or more supported actions to receive notifications when the alarm state changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmeventactions.html#cfn-iotevents-alarmmodel-alarmeventactions-alarmactions
	//
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
}

