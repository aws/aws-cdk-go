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
//   policiesProperty := &policiesProperty{
//   	attributes: []interface{}{
//   		attributes,
//   	},
//   	policyName: jsii.String("policyName"),
//   	policyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	instancePorts: []*string{
//   		jsii.String("instancePorts"),
//   	},
//   	loadBalancerPorts: []*string{
//   		jsii.String("loadBalancerPorts"),
//   	},
//   }
//
type CfnLoadBalancer_PoliciesProperty struct {
	// The policy attributes.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the policy type.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The instance ports for the policy.
	//
	// Required only for some policy types.
	InstancePorts *[]*string `field:"optional" json:"instancePorts" yaml:"instancePorts"`
	// The load balancer ports for the policy.
	//
	// Required only for some policy types.
	LoadBalancerPorts *[]*string `field:"optional" json:"loadBalancerPorts" yaml:"loadBalancerPorts"`
}

