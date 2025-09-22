package awsbedrock


// Contains configurations for the service to use for storing the input into the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageFlowNodeServiceConfigurationProperty := &StorageFlowNodeServiceConfigurationProperty{
//   	S3: &StorageFlowNodeS3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-storageflownodeserviceconfiguration.html
//
type CfnFlow_StorageFlowNodeServiceConfigurationProperty struct {
	// Contains configurations for the Amazon S3 location in which to store the input into the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-storageflownodeserviceconfiguration.html#cfn-bedrock-flow-storageflownodeserviceconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

