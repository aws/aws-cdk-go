package awselasticloadbalancing


// Specifies a policy for duration-based session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lBCookieStickinessPolicyProperty := &LBCookieStickinessPolicyProperty{
//   	CookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   	PolicyName: jsii.String("policyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy.html
//
type CfnLoadBalancer_LBCookieStickinessPolicyProperty struct {
	// The time period, in seconds, after which the cookie should be considered stale.
	//
	// If this parameter is not specified, the stickiness session lasts for the duration of the browser session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy.html#cfn-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy-cookieexpirationperiod
	//
	CookieExpirationPeriod *string `field:"optional" json:"cookieExpirationPeriod" yaml:"cookieExpirationPeriod"`
	// The name of the policy.
	//
	// This name must be unique within the set of policies for this load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy.html#cfn-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

