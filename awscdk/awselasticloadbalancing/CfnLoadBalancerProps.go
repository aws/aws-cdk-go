package awselasticloadbalancing

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLoadBalancer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   cfnLoadBalancerProps := &CfnLoadBalancerProps{
//   	Listeners: []interface{}{
//   		&ListenersProperty{
//   			InstancePort: jsii.String("instancePort"),
//   			LoadBalancerPort: jsii.String("loadBalancerPort"),
//   			Protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			InstanceProtocol: jsii.String("instanceProtocol"),
//   			PolicyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			SslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AccessLoggingPolicy: &AccessLoggingPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//   		S3BucketName: jsii.String("s3BucketName"),
//
//   		// the properties below are optional
//   		EmitInterval: jsii.Number(123),
//   		S3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	AppCookieStickinessPolicy: []interface{}{
//   		&AppCookieStickinessPolicyProperty{
//   			CookieName: jsii.String("cookieName"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ConnectionDrainingPolicy: &ConnectionDrainingPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Timeout: jsii.Number(123),
//   	},
//   	ConnectionSettings: &ConnectionSettingsProperty{
//   		IdleTimeout: jsii.Number(123),
//   	},
//   	CrossZone: jsii.Boolean(false),
//   	HealthCheck: &HealthCheckProperty{
//   		HealthyThreshold: jsii.String("healthyThreshold"),
//   		Interval: jsii.String("interval"),
//   		Target: jsii.String("target"),
//   		Timeout: jsii.String("timeout"),
//   		UnhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	Instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	LbCookieStickinessPolicy: []interface{}{
//   		&LBCookieStickinessPolicyProperty{
//   			CookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	Policies: []interface{}{
//   		&PoliciesProperty{
//   			Attributes: []interface{}{
//   				attributes,
//   			},
//   			PolicyName: jsii.String("policyName"),
//   			PolicyType: jsii.String("policyType"),
//
//   			// the properties below are optional
//   			InstancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			LoadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   		},
//   	},
//   	Scheme: jsii.String("scheme"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html
//
type CfnLoadBalancerProps struct {
	// The listeners for the load balancer. You can specify at most one listener per port.
	//
	// If you update the properties for a listener, AWS CloudFormation deletes the existing listener and creates a new one with the specified properties. While the new listener is being created, clients cannot connect to the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-listeners
	//
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// Information about where and how access logs are stored for the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-accessloggingpolicy
	//
	AccessLoggingPolicy interface{} `field:"optional" json:"accessLoggingPolicy" yaml:"accessLoggingPolicy"`
	// Information about a policy for application-controlled session stickiness.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-appcookiestickinesspolicy
	//
	AppCookieStickinessPolicy interface{} `field:"optional" json:"appCookieStickinessPolicy" yaml:"appCookieStickinessPolicy"`
	// The Availability Zones for a load balancer in a default VPC.
	//
	// For a load balancer in a nondefault VPC, specify `Subnets` instead.
	//
	// Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones. Otherwise, update requires no interruption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see [Configure connection draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html) in the *User Guide for Classic Load Balancers* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-connectiondrainingpolicy
	//
	ConnectionDrainingPolicy interface{} `field:"optional" json:"connectionDrainingPolicy" yaml:"connectionDrainingPolicy"`
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer. For more information, see [Configure idle connection timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html) in the *User Guide for Classic Load Balancers* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-connectionsettings
	//
	ConnectionSettings interface{} `field:"optional" json:"connectionSettings" yaml:"connectionSettings"`
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	//
	// For more information, see [Configure cross-zone load balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html) in the *User Guide for Classic Load Balancers* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-crosszone
	//
	CrossZone interface{} `field:"optional" json:"crossZone" yaml:"crossZone"`
	// The health check settings to use when evaluating the health of your EC2 instances.
	//
	// Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings. Otherwise, update requires no interruption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The IDs of the instances for the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-instances
	//
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Information about a policy for duration-based session stickiness.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-lbcookiestickinesspolicy
	//
	LbCookieStickinessPolicy interface{} `field:"optional" json:"lbCookieStickinessPolicy" yaml:"lbCookieStickinessPolicy"`
	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) . If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-loadbalancername
	//
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The policies defined for your Classic Load Balancer.
	//
	// Specify only back-end server policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If `Scheme` is `internet-facing` , the load balancer has a public DNS name that resolves to a public IP address.
	//
	// If `Scheme` is `internal` , the load balancer has a public DNS name that resolves to a private IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-scheme
	//
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The security groups for the load balancer.
	//
	// Valid only for load balancers in a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.
	//
	// Update requires replacement if you did not previously specify a subnet or if you are removing all subnets. Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone, you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags associated with a load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html#cfn-elasticloadbalancing-loadbalancer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

