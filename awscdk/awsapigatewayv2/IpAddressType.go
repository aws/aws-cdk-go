package awsapigatewayv2


// Supported IP Address Types.
//
// Example:
//   apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
//   	IpAddressType: apigwv2.IpAddressType_DUAL_STACK,
//   })
//
type IpAddressType string

const (
	// IPv4 address type.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// IPv4 and IPv6 address type.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

