package awsmsk


// Details of an Amazon MSK Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonMskClusterProperty := &AmazonMskClusterProperty{
//   	MskClusterArn: jsii.String("mskClusterArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-amazonmskcluster.html
//
type CfnReplicator_AmazonMskClusterProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-amazonmskcluster.html#cfn-msk-replicator-amazonmskcluster-mskclusterarn
	//
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
}

