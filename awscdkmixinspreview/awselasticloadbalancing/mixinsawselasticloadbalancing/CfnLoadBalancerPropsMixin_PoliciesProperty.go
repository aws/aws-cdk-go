package mixinsawselasticloadbalancing


// Specifies policies for your Classic Load Balancer.
//
// To associate policies with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var attributes interface{}
//
//   policiesProperty := &PoliciesProperty{
//   	Attributes: []interface{}{
//   		attributes,
//   	},
//   	InstancePorts: []*string{
//   		jsii.String("instancePorts"),
//   	},
//   	LoadBalancerPorts: []*string{
//   		jsii.String("loadBalancerPorts"),
//   	},
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html
//
type CfnLoadBalancerPropsMixin_PoliciesProperty struct {
	// The policy attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
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
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The name of the policy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-policies.html#cfn-elasticloadbalancing-loadbalancer-policies-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
}

