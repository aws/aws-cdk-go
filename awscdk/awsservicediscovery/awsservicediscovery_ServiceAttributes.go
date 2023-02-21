package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namespace iNamespace
//
//   serviceAttributes := &serviceAttributes{
//   	dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   	namespace: namespace,
//   	routingPolicy: awscdk.*Aws_servicediscovery.routingPolicy_WEIGHTED,
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceId: jsii.String("serviceId"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
// Experimental.
type ServiceAttributes struct {
	// Experimental.
	DnsRecordType DnsRecordType `field:"required" json:"dnsRecordType" yaml:"dnsRecordType"`
	// Experimental.
	Namespace INamespace `field:"required" json:"namespace" yaml:"namespace"`
	// Experimental.
	RoutingPolicy RoutingPolicy `field:"required" json:"routingPolicy" yaml:"routingPolicy"`
	// Experimental.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// Experimental.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

