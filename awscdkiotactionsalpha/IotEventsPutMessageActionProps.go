package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for the IoT Events.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var role iRole
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &InputProps{
//   	AttributeJsonPaths: []*string{
//   		jsii.String("payload.temperature"),
//   		jsii.String("payload.transactionId"),
//   	},
//   })
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewIotEventsPutMessageAction(input, &IotEventsPutMessageActionProps{
//   			BatchMode: jsii.Boolean(true),
//   			 // optional property, default is 'false'
//   			MessageId: jsii.String("${payload.transactionId}"),
//   			 // optional property, default is a new UUID
//   			Role: role,
//   		}),
//   	},
//   })
//
// Experimental.
type IotEventsPutMessageActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to process the event actions as a batch.
	//
	// When batchMode is true, you can't specify a messageId.
	//
	// When batchMode is true and the rule SQL statement evaluates to an Array,
	// each Array element is treated as a separate message when Events by calling BatchPutMessage.
	// The resulting array can't have more than 10 messages.
	// Experimental.
	BatchMode *bool `field:"optional" json:"batchMode" yaml:"batchMode"`
	// The ID of the message.
	//
	// When batchMode is true, you can't specify a messageId--a new UUID value will be assigned.
	// Assign a value to this property to ensure that only one input (message) with a given messageId will be processed by an AWS IoT Events detector.
	// Experimental.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}

