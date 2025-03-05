package awsbedrock


// Contains configurations for the Amazon S3 location from which to retrieve data to return as the output from the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retrievalFlowNodeS3ConfigurationProperty := &RetrievalFlowNodeS3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-retrievalflownodes3configuration.html
//
type CfnFlow_RetrievalFlowNodeS3ConfigurationProperty struct {
	// The name of the Amazon S3 bucket from which to retrieve data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-retrievalflownodes3configuration.html#cfn-bedrock-flow-retrievalflownodes3configuration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

