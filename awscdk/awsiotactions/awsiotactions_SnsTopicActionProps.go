package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration options for the SNS topic action.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type SnsTopicActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The message format of the message to publish.
	//
	// SNS uses this setting to determine if the payload should be parsed and relevant platform-specific bits of the payload should be extracted.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-and-json-formats.html
	//
	// Experimental.
	MessageFormat SnsActionMessageFormat `field:"optional" json:"messageFormat" yaml:"messageFormat"`
}

