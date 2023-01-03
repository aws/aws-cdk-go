package awsecs


// Interface for setting the properties of proxy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationProps := &appMeshProxyConfigurationProps{
//   	appPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	proxyEgressPort: jsii.Number(123),
//   	proxyIngressPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	egressIgnoredIPs: []*string{
//   		jsii.String("egressIgnoredIPs"),
//   	},
//   	egressIgnoredPorts: []*f64{
//   		jsii.Number(123),
//   	},
//   	ignoredGID: jsii.Number(123),
//   	ignoredUID: jsii.Number(123),
//   }
//
type AppMeshProxyConfigurationProps struct {
	// The list of ports that the application uses.
	//
	// Network traffic to these ports is forwarded to the ProxyIngressPort and ProxyEgressPort.
	AppPorts *[]*float64 `field:"required" json:"appPorts" yaml:"appPorts"`
	// Specifies the port that outgoing traffic from the AppPorts is directed to.
	ProxyEgressPort *float64 `field:"required" json:"proxyEgressPort" yaml:"proxyEgressPort"`
	// Specifies the port that incoming traffic to the AppPorts is directed to.
	ProxyIngressPort *float64 `field:"required" json:"proxyIngressPort" yaml:"proxyIngressPort"`
	// The egress traffic going to these specified IP addresses is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredIPs *[]*string `field:"optional" json:"egressIgnoredIPs" yaml:"egressIgnoredIPs"`
	// The egress traffic going to these specified ports is ignored and not redirected to the ProxyEgressPort.
	//
	// It can be an empty list.
	EgressIgnoredPorts *[]*float64 `field:"optional" json:"egressIgnoredPorts" yaml:"egressIgnoredPorts"`
	// The group ID (GID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredUID is specified, this field can be empty.
	IgnoredGID *float64 `field:"optional" json:"ignoredGID" yaml:"ignoredGID"`
	// The user ID (UID) of the proxy container as defined by the user parameter in a container definition.
	//
	// This is used to ensure the proxy ignores its own traffic. If IgnoredGID is specified, this field can be empty.
	IgnoredUID *float64 `field:"optional" json:"ignoredUID" yaml:"ignoredUID"`
}

