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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-service-dnsentry.html
//
type CfnService_DnsEntryProperty struct {
	// The domain name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-service-dnsentry.html#cfn-vpclattice-service-dnsentry-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The ID of the hosted zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-service-dnsentry.html#cfn-vpclattice-service-dnsentry-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}

