package awseks


// Indicates the current configuration of the block storage capability on your EKS Auto Mode cluster.
//
// For example, if the capability is enabled or disabled. If the block storage capability is enabled, EKS Auto Mode will create and delete EBS volumes in your AWS account. For more information, see EKS Auto Mode block storage capability in the EKS User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockStorageProperty := &BlockStorageProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-blockstorage.html
//
type CfnCluster_BlockStorageProperty struct {
	// Indicates if the block storage capability is enabled on your EKS Auto Mode cluster.
	//
	// If the block storage capability is enabled, EKS Auto Mode will create and delete EBS volumes in your AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-blockstorage.html#cfn-eks-cluster-blockstorage-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

