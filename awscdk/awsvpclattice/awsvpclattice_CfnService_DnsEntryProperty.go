package awsvpclattice


// Describes the DNS information of a service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsEntryProperty := &DnsEntryProperty{
//   	DomainName: jsii.String("domainName"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type CfnService_DnsEntryProperty struct {
	// The domain name of the service.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The ID of the hosted zone.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}

