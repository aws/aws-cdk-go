package awselasticloadbalancingv2


// Properties for a Network Listener attached to a Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCertificate listenerCertificate
//   var networkListenerAction networkListenerAction
//   var networkLoadBalancer networkLoadBalancer
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerProps := &NetworkListenerProps{
//   	LoadBalancer: networkLoadBalancer,
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	AlpnPolicy: awscdk.Aws_elasticloadbalancingv2.AlpnPolicy_HTTP1_ONLY,
//   	Certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   	DefaultAction: networkListenerAction,
//   	DefaultTargetGroups: []iNetworkTargetGroup{
//   		networkTargetGroup,
//   	},
//   	Protocol: awscdk.*Aws_elasticloadbalancingv2.Protocol_HTTP,
//   	SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED,
//   }
//
// Experimental.
type NetworkListenerProps struct {
	// The port on which the listener listens for requests.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	// Experimental.
	AlpnPolicy AlpnPolicy `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	// Experimental.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	// Experimental.
	DefaultAction NetworkListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	// Experimental.
	DefaultTargetGroups *[]INetworkTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	// Experimental.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SSL Policy.
	// Experimental.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	// Experimental.
	LoadBalancer INetworkLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
}

