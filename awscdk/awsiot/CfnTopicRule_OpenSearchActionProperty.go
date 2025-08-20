package awsiot


// Describes an action that writes data to an Amazon OpenSearch Service domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openSearchActionProperty := &OpenSearchActionProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	Id: jsii.String("id"),
//   	Index: jsii.String("index"),
//   	RoleArn: jsii.String("roleArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html
//
type CfnTopicRule_OpenSearchActionProperty struct {
	// The endpoint of your OpenSearch domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html#cfn-iot-topicrule-opensearchaction-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for the document you are storing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html#cfn-iot-topicrule-opensearchaction-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The OpenSearch index where you want to store your data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html#cfn-iot-topicrule-opensearchaction-index
	//
	Index *string `field:"required" json:"index" yaml:"index"`
	// The IAM role ARN that has access to OpenSearch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html#cfn-iot-topicrule-opensearchaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of document you are storing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-opensearchaction.html#cfn-iot-topicrule-opensearchaction-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

