package awselasticloadbalancingv2


// Properties for defining a standalone ApplicationListener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancer applicationLoadBalancer
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCertificate listenerCertificate
//
//   applicationListenerProps := &applicationListenerProps{
//   	loadBalancer: applicationLoadBalancer,
//
//   	// the properties below are optional
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   	defaultAction: listenerAction,
//   	defaultTargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   	open: jsii.Boolean(false),
//   	port: jsii.Number(123),
//   	protocol: awscdk.Aws_elasticloadbalancingv2.applicationProtocol_HTTP,
//   	sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   }
//
type ApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction ListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]IApplicationTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer IApplicationLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
}

