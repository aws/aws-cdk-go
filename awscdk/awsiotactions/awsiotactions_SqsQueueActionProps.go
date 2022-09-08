package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration properties of an action for SQS.
//
// Example:
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSqsQueueAction(queue, &sqsQueueActionProps{
//   			useBase64: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type SqsQueueActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specifies whether to use Base64 encoding.
	// Experimental.
	UseBase64 *bool `field:"optional" json:"useBase64" yaml:"useBase64"`
}

