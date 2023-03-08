package awsroute53resolver


// In a [CreateResolverEndpoint](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html) request, the IP address that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints). `IpAddressRequest` also includes the ID of the subnet that contains the IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipAddressRequestProperty := &ipAddressRequestProperty{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	ip: jsii.String("ip"),
//   	ipv6: jsii.String("ipv6"),
//   }
//
type CfnResolverEndpoint_IpAddressRequestProperty struct {
	// The ID of the subnet that contains the IP address.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The IP address that you want to use for DNS queries.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// `CfnResolverEndpoint.IpAddressRequestProperty.Ipv6`.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
}

