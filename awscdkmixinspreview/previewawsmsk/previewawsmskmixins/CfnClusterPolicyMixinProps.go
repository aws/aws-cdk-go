package previewawsmskmixins


// Properties for CfnClusterPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnClusterPolicyMixinProps := &CfnClusterPolicyMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html
//
type CfnClusterPolicyMixinProps struct {
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html#cfn-msk-clusterpolicy-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Resource policy for the cluster.
	//
	// The maximum size supported for a resource-based policy document is 20 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-clusterpolicy.html#cfn-msk-clusterpolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

