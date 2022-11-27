package awsecs


// Interface for service connect Service props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectService := &serviceConnectService{
//   	portMappingName: jsii.String("portMappingName"),
//
//   	// the properties below are optional
//   	discoveryName: jsii.String("discoveryName"),
//   	dnsName: jsii.String("dnsName"),
//   	ingressPortOverride: jsii.Number(123),
//   	port: jsii.Number(123),
//   }
//
type ServiceConnectService struct {
	// portMappingName specifies which port and protocol combination should be used for this service connect service.
	PortMappingName *string `field:"required" json:"portMappingName" yaml:"portMappingName"`
	// Optionally specifies an intermediate dns name to register in the CloudMap namespace.
	//
	// This is required if you wish to use the same port mapping name in more than one service.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// The terse DNS alias to use for this port mapping in the service connect mesh.
	//
	// Service Connect-enabled clients will be able to reach this service at
	// http://dnsName:port.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Optional.
	//
	// The port on the Service Connect agent container to use for traffic ingress to this service.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// The port for clients to use to communicate with this service via Service Connect.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

