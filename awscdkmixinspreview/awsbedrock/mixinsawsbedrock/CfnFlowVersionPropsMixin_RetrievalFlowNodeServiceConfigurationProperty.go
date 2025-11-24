package mixinsawsbedrock


// Contains configurations for the service to use for retrieving data to return as the output from the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retrievalFlowNodeServiceConfigurationProperty := &RetrievalFlowNodeServiceConfigurationProperty{
//   	S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodeserviceconfiguration.html
//
type CfnFlowVersionPropsMixin_RetrievalFlowNodeServiceConfigurationProperty struct {
	// Contains configurations for the Amazon S3 location from which to retrieve data to return as the output from the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodeserviceconfiguration.html#cfn-bedrock-flowversion-retrievalflownodeserviceconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

