package previewawselasticloadbalancingmixins


// Specifies a listener for your Classic Load Balancer.
//
// Modifying any property replaces the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   listenersProperty := &ListenersProperty{
//   	InstancePort: jsii.String("instancePort"),
//   	InstanceProtocol: jsii.String("instanceProtocol"),
//   	LoadBalancerPort: jsii.String("loadBalancerPort"),
//   	PolicyNames: []*string{
//   		jsii.String("policyNames"),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	SslCertificateId: jsii.String("sslCertificateId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html
//
type CfnLoadBalancerPropsMixin_ListenersProperty struct {
	// The port on which the instance is listening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-instanceport
	//
	InstancePort *string `field:"optional" json:"instancePort" yaml:"instancePort"`
	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.
	//
	// If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is secure, (HTTPS or SSL), the listener's `InstanceProtocol` must also be secure.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is HTTP or TCP, the listener's `InstanceProtocol` must be HTTP or TCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-instanceprotocol
	//
	InstanceProtocol *string `field:"optional" json:"instanceProtocol" yaml:"instanceProtocol"`
	// The port on which the load balancer is listening.
	//
	// On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-loadbalancerport
	//
	LoadBalancerPort *string `field:"optional" json:"loadBalancerPort" yaml:"loadBalancerPort"`
	// The names of the policies to associate with the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-policynames
	//
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The Amazon Resource Name (ARN) of the server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-listeners.html#cfn-elasticloadbalancing-loadbalancer-listeners-sslcertificateid
	//
	SslCertificateId *string `field:"optional" json:"sslCertificateId" yaml:"sslCertificateId"`
}

