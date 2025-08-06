package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for the Kinesis Data stream.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewKinesisPutRecordAction(stream, &KinesisPutRecordActionProps{
//   			PartitionKey: jsii.String("${newuuid()}"),
//   		}),
//   	},
//   })
//
// Experimental.
type KinesisPutRecordActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The partition key used to determine to which shard the data is written.
	//
	// The partition key is usually composed of an expression (for example, ${topic()} or ${timestamp()}).
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html#API_PutRecord_RequestParameters
	//
	// Experimental.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

