package awseks


// Todo: add description.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConfigProperty := &StorageConfigProperty{
//   	BlockStorage: &BlockStorageProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-storageconfig.html
//
type CfnCluster_StorageConfigProperty struct {
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-storageconfig.html#cfn-eks-cluster-storageconfig-blockstorage
	//
	BlockStorage interface{} `field:"optional" json:"blockStorage" yaml:"blockStorage"`
}

