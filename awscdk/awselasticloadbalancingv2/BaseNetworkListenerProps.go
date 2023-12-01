package awselasticloadbalancingv2


// Basic properties for a Network Listener.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//   import elb2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var clb loadBalancer
//   var alb applicationLoadBalancer
//   var nlb networkLoadBalancer
//
//
//   albListener := alb.AddListener(jsii.String("ALBListener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   albTargetGroup := albListener.AddTargets(jsii.String("ALBFleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   nlbListener := nlb.AddListener(jsii.String("NLBListener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(80),
//   })
//   nlbTargetGroup := nlbListener.AddTargets(jsii.String("NLBFleet"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
//   	LoadBalancers: []loadBalancer{
//   		codedeploy.*loadBalancer_Classic(clb),
//   		codedeploy.*loadBalancer_Application(albTargetGroup),
//   		codedeploy.*loadBalancer_Network(nlbTargetGroup),
//   	},
//   })
//
type BaseNetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	// Default: - None.
	//
	AlpnPolicy AlpnPolicy `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	// Default: - No certificates.
	//
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	// Default: - None.
	//
	DefaultAction NetworkListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	// Default: - None.
	//
	DefaultTargetGroups *[]INetworkTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	// Default: - TLS if certificates are provided. TCP otherwise.
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SSL Policy.
	// Default: - Current predefined security policy.
	//
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

