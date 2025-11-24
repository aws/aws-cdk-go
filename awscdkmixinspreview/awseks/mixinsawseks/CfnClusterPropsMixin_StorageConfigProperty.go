package mixinsawseks


// Request to update the configuration of the storage capability of your EKS Auto Mode cluster.
//
// For example, enable the capability. For more information, see EKS Auto Mode block storage capability in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageConfigProperty := &StorageConfigProperty{
//   	BlockStorage: &BlockStorageProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-storageconfig.html
//
type CfnClusterPropsMixin_StorageConfigProperty struct {
	// Request to configure EBS Block Storage settings for your EKS Auto Mode cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-storageconfig.html#cfn-eks-cluster-storageconfig-blockstorage
	//
	BlockStorage interface{} `field:"optional" json:"blockStorage" yaml:"blockStorage"`
}

