package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for Open Search.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//   var domain domain
//
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   })
//
//   topicRule.AddAction(actions.NewOpenSearchAction(domain, &OpenSearchActionProps{
//   	Id: jsii.String("my-id"),
//   	Index: jsii.String("my-index"),
//   	Type: jsii.String("my-type"),
//   }))
//
// Experimental.
type OpenSearchActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The unique identifier for the document you are storing.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The OpenSearch index where you want to store your data.
	// Experimental.
	Index *string `field:"required" json:"index" yaml:"index"`
	// The type of document you are storing.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

