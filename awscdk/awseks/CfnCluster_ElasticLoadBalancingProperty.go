package awseks


// Todo: add description.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticLoadBalancingProperty := &ElasticLoadBalancingProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-elasticloadbalancing.html
//
type CfnCluster_ElasticLoadBalancingProperty struct {
	// Todo: add description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-elasticloadbalancing.html#cfn-eks-cluster-elasticloadbalancing-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

