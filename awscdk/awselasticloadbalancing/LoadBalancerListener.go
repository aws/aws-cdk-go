package awselasticloadbalancing

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Add a backend to the load balancer.
//
// Example:
//   var vpc iVpc
//
//   var myAutoScalingGroup autoScalingGroup
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   	HealthCheck: &HealthCheck{
//   		Port: jsii.Number(80),
//   	},
//   })
//   lb.AddTarget(myAutoScalingGroup)
//   lb.AddListener(&LoadBalancerListener{
//   	ExternalPort: jsii.Number(80),
//   })
//
type LoadBalancerListener struct {
	// External listening port.
	ExternalPort *float64 `field:"required" json:"externalPort" yaml:"externalPort"`
	// Allow connections to the load balancer from the given set of connection peers.
	//
	// By default, connections will be allowed from anywhere. Set this to an empty list
	// to deny connections, or supply a custom list of peers to allow connections from
	// (IP ranges or security groups).
	// Default: Anywhere.
	//
	AllowConnectionsFrom *[]awsec2.IConnectable `field:"optional" json:"allowConnectionsFrom" yaml:"allowConnectionsFrom"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the external port is either 80 or 443.
	ExternalProtocol LoadBalancingProtocol `field:"optional" json:"externalProtocol" yaml:"externalProtocol"`
	// Instance listening port.
	//
	// Same as the externalPort if not specified.
	// Default: externalPort.
	//
	InternalPort *float64 `field:"optional" json:"internalPort" yaml:"internalPort"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the internal port is either 80 or 443.
	//
	// The instance protocol is 'tcp' if the front-end protocol
	// is 'tcp' or 'ssl', the instance protocol is 'http' if the
	// front-end protocol is 'https'.
	InternalProtocol LoadBalancingProtocol `field:"optional" json:"internalProtocol" yaml:"internalProtocol"`
	// SSL policy names.
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// the ARN of the SSL certificate.
	// Default: - none.
	//
	SslCertificateArn *string `field:"optional" json:"sslCertificateArn" yaml:"sslCertificateArn"`
}

