package awsopensearchservice


// The IP address type for the domain.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
//   	IpAddressType: awscdk.IpAddressType_DUAL_STACK,
//   })
//
type IpAddressType string

const (
	// IPv4 addresses only.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// IPv4 and IPv6 addresses.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

