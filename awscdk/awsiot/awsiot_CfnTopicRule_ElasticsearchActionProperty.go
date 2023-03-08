package awsiot


// Describes an action that writes data to an Amazon OpenSearch Service domain.
//
// > The `Elasticsearch` action can only be used by existing rule actions. To create a new rule action or to update an existing rule action, use the `OpenSearch` rule action instead. For more information, see [OpenSearchAction](https://docs.aws.amazon.com//iot/latest/apireference/API_OpenSearchAction.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchActionProperty := &elasticsearchActionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	id: jsii.String("id"),
//   	index: jsii.String("index"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnTopicRule_ElasticsearchActionProperty struct {
	// The endpoint of your OpenSearch domain.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for the document you are storing.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The index where you want to store your data.
	Index *string `field:"required" json:"index" yaml:"index"`
	// The IAM role ARN that has access to OpenSearch.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of document you are storing.
	Type *string `field:"required" json:"type" yaml:"type"`
}

