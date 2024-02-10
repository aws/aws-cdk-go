package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Interface for service connect Service props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectService := &ServiceConnectService{
//   	PortMappingName: jsii.String("portMappingName"),
//
//   	// the properties below are optional
//   	DiscoveryName: jsii.String("discoveryName"),
//   	DnsName: jsii.String("dnsName"),
//   	IdleTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	IngressPortOverride: jsii.Number(123),
//   	PerRequestTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	Port: jsii.Number(123),
//   }
//
type ServiceConnectService struct {
	// portMappingName specifies which port and protocol combination should be used for this service connect service.
	PortMappingName *string `field:"required" json:"portMappingName" yaml:"portMappingName"`
	// Optionally specifies an intermediate dns name to register in the CloudMap namespace.
	//
	// This is required if you wish to use the same port mapping name in more than one service.
	// Default: - port mapping name.
	//
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// The terse DNS alias to use for this port mapping in the service connect mesh.
	//
	// Service Connect-enabled clients will be able to reach this service at
	// http://dnsName:port.
	// Default: - No alias is created. The service is reachable at `portMappingName.namespace:port`.
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The amount of time in seconds a connection for Service Connect will stay active while idle.
	//
	// A value of 0 can be set to disable `idleTimeout`.
	//
	// If `idleTimeout` is set to a time that is less than `perRequestTimeout`, the connection will close
	// when the `idleTimeout` is reached and not the `perRequestTimeout`.
	// Default: - Duration.minutes(5) for HTTP/HTTP2/GRPC, Duration.hours(1) for TCP.
	//
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Optional.
	//
	// The port on the Service Connect agent container to use for traffic ingress to this service.
	// Default: - none.
	//
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// The amount of time waiting for the upstream to respond with a complete response per request for Service Connect.
	//
	// A value of 0 can be set to disable `perRequestTimeout`.
	// Can only be set when the `appProtocol` for the application container is HTTP/HTTP2/GRPC.
	//
	// If `idleTimeout` is set to a time that is less than `perRequestTimeout`, the connection will close
	// when the `idleTimeout` is reached and not the `perRequestTimeout`.
	// Default: - Duration.seconds(15)
	//
	PerRequestTimeout awscdk.Duration `field:"optional" json:"perRequestTimeout" yaml:"perRequestTimeout"`
	// The port for clients to use to communicate with this service via Service Connect.
	// Default: the container port specified by the port mapping in portMappingName.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

