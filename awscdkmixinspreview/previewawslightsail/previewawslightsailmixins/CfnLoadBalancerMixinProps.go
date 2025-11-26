package previewawslightsailmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLoadBalancerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerMixinProps := &CfnLoadBalancerMixinProps{
//   	AttachedInstances: []*string{
//   		jsii.String("attachedInstances"),
//   	},
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	InstancePort: jsii.Number(123),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	SessionStickinessEnabled: jsii.Boolean(false),
//   	SessionStickinessLbCookieDurationSeconds: jsii.String("sessionStickinessLbCookieDurationSeconds"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsPolicyName: jsii.String("tlsPolicyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html
//
type CfnLoadBalancerMixinProps struct {
	// The Lightsail instances to attach to the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-attachedinstances
	//
	AttachedInstances *[]*string `field:"optional" json:"attachedInstances" yaml:"attachedInstances"`
	// The path on the attached instance where the health check will be performed.
	//
	// If no path is specified, the load balancer tries to make a request to the default (root) page ( `/index.html` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-healthcheckpath
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port that the load balancer uses to direct traffic to your Lightsail instances.
	//
	// For HTTP traffic, specify port `80` . For HTTPS traffic, specify port `443` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-instanceport
	//
	InstancePort *float64 `field:"optional" json:"instancePort" yaml:"instancePort"`
	// The IP address type of the load balancer.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for both IPv4 and IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The name of the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-loadbalancername
	//
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// A Boolean value indicating whether session stickiness is enabled.
	//
	// Enable session stickiness (also known as *session affinity* ) to bind a user's session to a specific instance. This ensures that all requests from the user during the session are sent to the same instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-sessionstickinessenabled
	//
	SessionStickinessEnabled interface{} `field:"optional" json:"sessionStickinessEnabled" yaml:"sessionStickinessEnabled"`
	// The time period, in seconds, after which the load balancer session stickiness cookie should be considered stale.
	//
	// If you do not specify this parameter, the default value is 0, which indicates that the sticky session should last for the duration of the browser session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-sessionstickinesslbcookiedurationseconds
	//
	SessionStickinessLbCookieDurationSeconds *string `field:"optional" json:"sessionStickinessLbCookieDurationSeconds" yaml:"sessionStickinessLbCookieDurationSeconds"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the TLS security policy for the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancer.html#cfn-lightsail-loadbalancer-tlspolicyname
	//
	TlsPolicyName *string `field:"optional" json:"tlsPolicyName" yaml:"tlsPolicyName"`
}

