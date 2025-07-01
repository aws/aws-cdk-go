package awselasticloadbalancing


// Specifies policies for your Classic Load Balancer.
//
// To associate policies with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   policiesProperty := &PoliciesProperty{
//   	Attributes: []interface{}{
//   		attributes,
//   	},
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	InstancePorts: []*string{
//   		jsii.String("instancePorts"),
//   	},
//   	LoadBalancerPorts: []*string{
//   		jsii.String("loadBalancerPorts"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html
//
type CfnLoadBalancer_PoliciesProperty struct {
	// The policy attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-attributes
	//
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the policy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-policytype
	//
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The instance ports for the policy.
	//
	// Required only for some policy types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-instanceports
	//
	InstancePorts *[]*string `field:"optional" json:"instancePorts" yaml:"instancePorts"`
	// The load balancer ports for the policy.
	//
	// Required only for some policy types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-loadbalancerports
	//
	LoadBalancerPorts *[]*string `field:"optional" json:"loadBalancerPorts" yaml:"loadBalancerPorts"`
}

