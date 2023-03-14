package awsiot


// Describes an action that writes data to an Amazon OpenSearch Service domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openSearchActionProperty := &openSearchActionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	id: jsii.String("id"),
//   	index: jsii.String("index"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnTopicRule_OpenSearchActionProperty struct {
	// The endpoint of your OpenSearch domain.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for the document you are storing.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The OpenSearch index where you want to store your data.
	Index *string `field:"required" json:"index" yaml:"index"`
	// The IAM role ARN that has access to OpenSearch.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of document you are storing.
	Type *string `field:"required" json:"type" yaml:"type"`
}

