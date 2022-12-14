package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectServiceProperty := &serviceConnectServiceProperty{
//   	portName: jsii.String("portName"),
//
//   	// the properties below are optional
//   	clientAliases: []interface{}{
//   		&serviceConnectClientAliasProperty{
//   			port: jsii.Number(123),
//
//   			// the properties below are optional
//   			dnsName: jsii.String("dnsName"),
//   		},
//   	},
//   	discoveryName: jsii.String("discoveryName"),
//   	ingressPortOverride: jsii.Number(123),
//   }
//
type CfnService_ServiceConnectServiceProperty struct {
	// `CfnService.ServiceConnectServiceProperty.PortName`.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// `CfnService.ServiceConnectServiceProperty.ClientAliases`.
	ClientAliases interface{} `field:"optional" json:"clientAliases" yaml:"clientAliases"`
	// `CfnService.ServiceConnectServiceProperty.DiscoveryName`.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// `CfnService.ServiceConnectServiceProperty.IngressPortOverride`.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
}

