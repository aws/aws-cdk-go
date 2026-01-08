package awselasticloadbalancingv2


// Properties for defining a standalone ApplicationListener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancer ApplicationLoadBalancer
//   var applicationTargetGroup ApplicationTargetGroup
//   var listenerAction ListenerAction
//   var listenerCertificate ListenerCertificate
//   var trustStoreRef ITrustStoreRef
//
//   applicationListenerProps := &ApplicationListenerProps{
//   	LoadBalancer: applicationLoadBalancer,
//
//   	// the properties below are optional
//   	Certificates: []IListenerCertificate{
//   		listenerCertificate,
//   	},
//   	DefaultAction: listenerAction,
//   	DefaultTargetGroups: []IApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   	MutualAuthentication: &MutualAuthentication{
//   		AdvertiseTrustStoreCaNames: jsii.Boolean(false),
//   		IgnoreClientCertificateExpiry: jsii.Boolean(false),
//   		MutualAuthenticationMode: awscdk.Aws_elasticloadbalancingv2.MutualAuthenticationMode_OFF,
//   		TrustStore: trustStoreRef,
//   	},
//   	Open: jsii.Boolean(false),
//   	Port: jsii.Number(123),
//   	Protocol: awscdk.*Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   	SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED_TLS,
//   }
//
type ApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	// Default: - No certificates.
	//
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	// Default: - None.
	//
	DefaultAction ListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	// Default: - None.
	//
	DefaultTargetGroups *[]IApplicationTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// The mutual authentication configuration information.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/mutual-authentication.html
	//
	// Default: - No mutual authentication configuration.
	//
	MutualAuthentication *MutualAuthentication `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	// Default: true.
	//
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	// Default: - Determined from protocol if known.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	// Default: - Determined from port if known.
	//
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	// Default: - The current predefined security policy.
	//
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer IApplicationLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
}

