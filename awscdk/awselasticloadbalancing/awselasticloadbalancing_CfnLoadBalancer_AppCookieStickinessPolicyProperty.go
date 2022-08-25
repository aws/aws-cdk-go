package awselasticloadbalancing


// Specifies a policy for application-controlled session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appCookieStickinessPolicyProperty := &appCookieStickinessPolicyProperty{
//   	cookieName: jsii.String("cookieName"),
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnLoadBalancer_AppCookieStickinessPolicyProperty struct {
	// The name of the application cookie used for stickiness.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
	// The mnemonic name for the policy being created.
	//
	// The name must be unique within a set of policies for this load balancer.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

