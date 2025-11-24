package mixinsawsbedrock


// Contains configurations for a Retrieval node in a flow.
//
// This node retrieves data from the Amazon S3 location that you specify and returns it as the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retrievalFlowNodeConfigurationProperty := &RetrievalFlowNodeConfigurationProperty{
//   	ServiceConfiguration: &RetrievalFlowNodeServiceConfigurationProperty{
//   		S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodeconfiguration.html
//
type CfnFlowVersionPropsMixin_RetrievalFlowNodeConfigurationProperty struct {
	// Contains configurations for the service to use for retrieving data to return as the output from the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-retrievalflownodeconfiguration.html#cfn-bedrock-flowversion-retrievalflownodeconfiguration-serviceconfiguration
	//
	ServiceConfiguration interface{} `field:"optional" json:"serviceConfiguration" yaml:"serviceConfiguration"`
}

