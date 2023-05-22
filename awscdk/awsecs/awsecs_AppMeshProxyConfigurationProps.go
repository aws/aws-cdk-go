package awsecs


// Interface for setting the properties of proxy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationProps := &AppMeshProxyConfigurationProps{
//   	AppPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	ProxyEgressPort: jsii.Number(123),
//   	ProxyIngressPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	EgressIgnoredIPs: []*string{
//   		jsii.String("egressIgnoredIPs"),
//   	},
//   	EgressIgnoredPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	IgnoredGID: jsii.Number(123),
//   	IgnoredUID: jsii.Number(123),
//   }
//
// Experimental.
type AppMeshProxyConfigurationProps struct {
	// The list of ports that the application uses.
	//
	// Network traffic to these ports is forwarded to the ProxyIngressPort and ProxyEgressPort.
	// Experimental.
	AppPorts *[]*float64 `field:"required" json:"appPorts" yaml:"appPorts"`
	// Specifies the port that outgoing traffic from the AppPorts is directed to.
	// Experimental.
	ProxyEgressPort *float64 `field:"required" json:"proxyEgressPort" yaml:"proxyEgressPort"`
	// Specifies the port that incoming traffic to the AppPorts is directed to.
	// Experimental.
	ProxyIngressPort *float64 `field:"required" json:"proxyIngressPort" yaml:"proxyIngressPort"`
	// The egress traffic going to these specified IP addresses is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	// Experimental.
	EgressIgnoredIPs *[]*string `field:"optional" json:"egressIgnoredIPs" yaml:"egressIgnoredIPs"`
	// The egress traffic going to these specified ports is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	// Experimental.
	EgressIgnoredPorts *[]*float64 `field:"optional" json:"egressIgnoredPorts" yaml:"egressIgnoredPorts"`
	// The group ID (GID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredUID is specified, this field can be empty.
	// Experimental.
	IgnoredGID *float64 `field:"optional" json:"ignoredGID" yaml:"ignoredGID"`
	// The user ID (UID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredGID is specified, this field can be empty.
	// Experimental.
	IgnoredUID *float64 `field:"optional" json:"ignoredUID" yaml:"ignoredUID"`
}

