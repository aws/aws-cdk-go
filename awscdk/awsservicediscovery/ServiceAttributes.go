package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namespace iNamespace
//
//   serviceAttributes := &ServiceAttributes{
//   	DnsRecordType: awscdk.Aws_servicediscovery.DnsRecordType_A,
//   	Namespace: namespace,
//   	RoutingPolicy: awscdk.*Aws_servicediscovery.RoutingPolicy_WEIGHTED,
//   	ServiceArn: jsii.String("serviceArn"),
//   	ServiceId: jsii.String("serviceId"),
//   	ServiceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	DiscoveryType: awscdk.*Aws_servicediscovery.DiscoveryType_API,
//   }
//
type ServiceAttributes struct {
	DnsRecordType DnsRecordType `field:"required" json:"dnsRecordType" yaml:"dnsRecordType"`
	Namespace INamespace `field:"required" json:"namespace" yaml:"namespace"`
	RoutingPolicy RoutingPolicy `field:"required" json:"routingPolicy" yaml:"routingPolicy"`
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	DiscoveryType DiscoveryType `field:"optional" json:"discoveryType" yaml:"discoveryType"`
}

