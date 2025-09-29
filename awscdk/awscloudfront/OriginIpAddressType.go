package awscloudfront


// The IP address type for the origin.
//
// Determines whether CloudFront uses IPv4, IPv6, or both when connecting to the origin.
//
// Example:
//   origin := origins.NewHttpOrigin(jsii.String("www.example.com"), &HttpOriginProps{
//   	IpAddressType: cloudfront.OriginIpAddressType_IPV6,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: *Origin,
//   	},
//   })
//
type OriginIpAddressType string

const (
	// Use only IPv4 addresses.
	OriginIpAddressType_IPV4 OriginIpAddressType = "IPV4"
	// Use only IPv6 addresses.
	OriginIpAddressType_IPV6 OriginIpAddressType = "IPV6"
	// Use both IPv4 and IPv6 addresses.
	OriginIpAddressType_DUALSTACK OriginIpAddressType = "DUALSTACK"
)

