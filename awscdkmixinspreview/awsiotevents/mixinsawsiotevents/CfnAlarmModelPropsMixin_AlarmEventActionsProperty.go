package mixinsawsiotevents


// Contains information about one or more alarm actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   alarmEventActionsProperty := &AlarmEventActionsProperty{
//   	AlarmActions: []interface{}{
//   		&AlarmActionProperty{
//   			DynamoDb: &DynamoDBProperty{
//   				HashKeyField: jsii.String("hashKeyField"),
//   				HashKeyType: jsii.String("hashKeyType"),
//   				HashKeyValue: jsii.String("hashKeyValue"),
//   				Operation: jsii.String("operation"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				PayloadField: jsii.String("payloadField"),
//   				RangeKeyField: jsii.String("rangeKeyField"),
//   				RangeKeyType: jsii.String("rangeKeyType"),
//   				RangeKeyValue: jsii.String("rangeKeyValue"),
//   				TableName: jsii.String("tableName"),
//   			},
//   			DynamoDBv2: &DynamoDBv2Property{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				TableName: jsii.String("tableName"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				DeliveryStreamName: jsii.String("deliveryStreamName"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				Separator: jsii.String("separator"),
//   			},
//   			IotEvents: &IotEventsProperty{
//   				InputName: jsii.String("inputName"),
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
//   					Quality: jsii.String("quality"),
//   					Timestamp: &AssetPropertyTimestampProperty{
//   						OffsetInNanos: jsii.String("offsetInNanos"),
//   						TimeInSeconds: jsii.String("timeInSeconds"),
//   					},
//   					Value: &AssetPropertyVariantProperty{
//   						BooleanValue: jsii.String("booleanValue"),
//   						DoubleValue: jsii.String("doubleValue"),
//   						IntegerValue: jsii.String("integerValue"),
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			IotTopicPublish: &IotTopicPublishProperty{
//   				MqttTopic: jsii.String("mqttTopic"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Lambda: &LambdaProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Sns: &SnsProperty{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				TargetArn: jsii.String("targetArn"),
//   			},
//   			Sqs: &SqsProperty{
//   				Payload: &PayloadProperty{
//   					ContentExpression: jsii.String("contentExpression"),
//   					Type: jsii.String("type"),
//   				},
//   				QueueUrl: jsii.String("queueUrl"),
//   				UseBase64: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmeventactions.html
//
type CfnAlarmModelPropsMixin_AlarmEventActionsProperty struct {
	// Specifies one or more supported actions to receive notifications when the alarm state changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmeventactions.html#cfn-iotevents-alarmmodel-alarmeventactions-alarmactions
	//
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
}

