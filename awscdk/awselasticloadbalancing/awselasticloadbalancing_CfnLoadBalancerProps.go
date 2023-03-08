package awselasticloadbalancing

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnLoadBalancerProps := &cfnLoadBalancerProps{
//   	listeners: []interface{}{
//   		&listenersProperty{
//   			instancePort: jsii.String("instancePort"),
//   			loadBalancerPort: jsii.String("loadBalancerPort"),
//   			protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			instanceProtocol: jsii.String("instanceProtocol"),
//   			policyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			sslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	accessLoggingPolicy: &accessLoggingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//   		s3BucketName: jsii.String("s3BucketName"),
//
//   		// the properties below are optional
//   		emitInterval: jsii.Number(123),
//   		s3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	appCookieStickinessPolicy: []interface{}{
//   		&appCookieStickinessPolicyProperty{
//   			cookieName: jsii.String("cookieName"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	connectionDrainingPolicy: &connectionDrainingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		timeout: jsii.Number(123),
//   	},
//   	connectionSettings: &connectionSettingsProperty{
//   		idleTimeout: jsii.Number(123),
//   	},
//   	crossZone: jsii.Boolean(false),
//   	healthCheck: &healthCheckProperty{
//   		healthyThreshold: jsii.String("healthyThreshold"),
//   		interval: jsii.String("interval"),
//   		target: jsii.String("target"),
//   		timeout: jsii.String("timeout"),
//   		unhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	lbCookieStickinessPolicy: []interface{}{
//   		&lBCookieStickinessPolicyProperty{
//   			cookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	policies: []interface{}{
//   		&policiesProperty{
//   			attributes: []interface{}{
//   				attributes,
//   			},
//   			policyName: jsii.String("policyName"),
//   			policyType: jsii.String("policyType"),
//
//   			// the properties below are optional
//   			instancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			loadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   		},
//   	},
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLoadBalancerProps struct {
	// The listeners for the load balancer. You can specify at most one listener per port.
	//
	// If you update the properties for a listener, AWS CloudFormation deletes the existing listener and creates a new one with the specified properties. While the new listener is being created, clients cannot connect to the load balancer.
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// Information about where and how access logs are stored for the load balancer.
	AccessLoggingPolicy interface{} `field:"optional" json:"accessLoggingPolicy" yaml:"accessLoggingPolicy"`
	// Information about a policy for application-controlled session stickiness.
	AppCookieStickinessPolicy interface{} `field:"optional" json:"appCookieStickinessPolicy" yaml:"appCookieStickinessPolicy"`
	// The Availability Zones for the load balancer. For load balancers in a VPC, specify `Subnets` instead.
	//
	// Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones. Otherwise, update requires no interruption.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see [Configure Connection Draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html) in the *Classic Load Balancers Guide* .
	ConnectionDrainingPolicy interface{} `field:"optional" json:"connectionDrainingPolicy" yaml:"connectionDrainingPolicy"`
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer. For more information, see [Configure Idle Connection Timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html) in the *Classic Load Balancers Guide* .
	ConnectionSettings interface{} `field:"optional" json:"connectionSettings" yaml:"connectionSettings"`
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	//
	// For more information, see [Configure Cross-Zone Load Balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html) in the *Classic Load Balancers Guide* .
	CrossZone interface{} `field:"optional" json:"crossZone" yaml:"crossZone"`
	// The health check settings to use when evaluating the health of your EC2 instances.
	//
	// Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings. Otherwise, update requires no interruption.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The IDs of the instances for the load balancer.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Information about a policy for duration-based session stickiness.
	LbCookieStickinessPolicy interface{} `field:"optional" json:"lbCookieStickinessPolicy" yaml:"lbCookieStickinessPolicy"`
	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) . If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The policies defined for your Classic Load Balancer.
	//
	// Specify only back-end server policies.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If `Scheme` is `internet-facing` , the load balancer has a public DNS name that resolves to a public IP address.
	//
	// If `Scheme` is `internal` , the load balancer has a public DNS name that resolves to a private IP address.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The security groups for the load balancer.
	//
	// Valid only for load balancers in a VPC.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.
	//
	// Update requires replacement if you did not previously specify a subnet or if you are removing all subnets. Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone, you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original Availability Zone.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags associated with a load balancer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

