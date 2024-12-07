package awseks


// Todo: add description.
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
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-blockstorage.html#cfn-eks-cluster-blockstorage-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

