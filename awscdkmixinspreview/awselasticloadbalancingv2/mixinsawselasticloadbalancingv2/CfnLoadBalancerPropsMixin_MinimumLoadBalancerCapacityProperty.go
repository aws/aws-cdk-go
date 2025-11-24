package mixinsawselasticloadbalancingv2


// The minimum capacity for a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   minimumLoadBalancerCapacityProperty := &MinimumLoadBalancerCapacityProperty{
//   	CapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-minimumloadbalancercapacity.html
//
type CfnLoadBalancerPropsMixin_MinimumLoadBalancerCapacityProperty struct {
	// The number of capacity units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-minimumloadbalancercapacity.html#cfn-elasticloadbalancingv2-loadbalancer-minimumloadbalancercapacity-capacityunits
	//
	CapacityUnits *float64 `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
}

