package mixinsawseks


// Indicates the current configuration of the load balancing capability on your EKS Auto Mode cluster.
//
// For example, if the capability is enabled or disabled. For more information, see EKS Auto Mode load balancing capability in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   elasticLoadBalancingProperty := &ElasticLoadBalancingProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-elasticloadbalancing.html
//
type CfnClusterPropsMixin_ElasticLoadBalancingProperty struct {
	// Indicates if the load balancing capability is enabled on your EKS Auto Mode cluster.
	//
	// If the load balancing capability is enabled, EKS Auto Mode will create and delete load balancers in your AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-elasticloadbalancing.html#cfn-eks-cluster-elasticloadbalancing-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

