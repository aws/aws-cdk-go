package awselasticloadbalancing


// Specifies a listener for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenersProperty := &listenersProperty{
//   	instancePort: jsii.String("instancePort"),
//   	loadBalancerPort: jsii.String("loadBalancerPort"),
//   	protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	instanceProtocol: jsii.String("instanceProtocol"),
//   	policyNames: []*string{
//   		jsii.String("policyNames"),
//   	},
//   	sslCertificateId: jsii.String("sslCertificateId"),
//   }
//
type CfnLoadBalancer_ListenersProperty struct {
	// The port on which the instance is listening.
	InstancePort *string `field:"required" json:"instancePort" yaml:"instancePort"`
	// The port on which the load balancer is listening.
	//
	// On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.
	LoadBalancerPort *string `field:"required" json:"loadBalancerPort" yaml:"loadBalancerPort"`
	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.
	//
	// If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is secure, (HTTPS or SSL), the listener's `InstanceProtocol` must also be secure.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is HTTP or TCP, the listener's `InstanceProtocol` must be HTTP or TCP.
	InstanceProtocol *string `field:"optional" json:"instanceProtocol" yaml:"instanceProtocol"`
	// The names of the policies to associate with the listener.
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// The Amazon Resource Name (ARN) of the server certificate.
	SslCertificateId *string `field:"optional" json:"sslCertificateId" yaml:"sslCertificateId"`
}

