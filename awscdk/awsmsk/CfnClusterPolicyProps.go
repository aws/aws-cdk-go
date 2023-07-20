package awsmsk


// Properties for defining a `CfnClusterPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnClusterPolicyProps := &CfnClusterPolicyProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html
//
type CfnClusterPolicyProps struct {
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html#cfn-msk-clusterpolicy-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Resource policy for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html#cfn-msk-clusterpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

