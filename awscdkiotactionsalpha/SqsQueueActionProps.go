package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for SQS.
//
// Example:
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewSqsQueueAction(queue, &SqsQueueActionProps{
//   			UseBase64: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type SqsQueueActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specifies whether to use Base64 encoding.
	// Default: false.
	//
	// Experimental.
	UseBase64 *bool `field:"optional" json:"useBase64" yaml:"useBase64"`
}

