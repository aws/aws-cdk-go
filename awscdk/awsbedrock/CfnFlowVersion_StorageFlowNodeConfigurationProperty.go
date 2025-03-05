package awsbedrock


// Contains configurations for a Storage node in a flow.
//
// This node stores the input in an Amazon S3 location that you specify.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageFlowNodeConfigurationProperty := &StorageFlowNodeConfigurationProperty{
//   	ServiceConfiguration: &StorageFlowNodeServiceConfigurationProperty{
//   		S3: &StorageFlowNodeS3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodeconfiguration.html
//
type CfnFlowVersion_StorageFlowNodeConfigurationProperty struct {
	// Contains configurations for the service to use for storing the input into the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-storageflownodeconfiguration.html#cfn-bedrock-flowversion-storageflownodeconfiguration-serviceconfiguration
	//
	ServiceConfiguration interface{} `field:"required" json:"serviceConfiguration" yaml:"serviceConfiguration"`
}

