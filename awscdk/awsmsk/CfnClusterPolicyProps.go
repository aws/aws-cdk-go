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
type CfnClusterPolicyProps struct {
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Resource policy for the cluster.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

