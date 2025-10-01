package awsbedrock


// Contains configurations for the Amazon S3 location in which to store the input into the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageFlowNodeS3ConfigurationProperty := &StorageFlowNodeS3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodes3configuration.html
//
type CfnFlowVersion_StorageFlowNodeS3ConfigurationProperty struct {
	// The name of the Amazon S3 bucket in which to store the input into the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodes3configuration.html#cfn-bedrock-flowversion-storageflownodes3configuration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

