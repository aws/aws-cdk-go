package awsec2


// Describes the DNS options for an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsOptionsSpecificationProperty := &DnsOptionsSpecificationProperty{
//   	DnsRecordIpType: jsii.String("dnsRecordIpType"),
//   	PrivateDnsOnlyForInboundResolverEndpoint: jsii.String("privateDnsOnlyForInboundResolverEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html
//
type CfnVPCEndpoint_DnsOptionsSpecificationProperty struct {
	// The DNS records created for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-dnsrecordiptype
	//
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Indicates whether to enable private DNS only for inbound endpoints.
	//
	// This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-privatednsonlyforinboundresolverendpoint
	//
	PrivateDnsOnlyForInboundResolverEndpoint *string `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
}

