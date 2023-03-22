package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsEntryProperty := &dnsEntryProperty{
//   	domainName: jsii.String("domainName"),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type CfnService_DnsEntryProperty struct {
	// `CfnService.DnsEntryProperty.DomainName`.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// `CfnService.DnsEntryProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}

