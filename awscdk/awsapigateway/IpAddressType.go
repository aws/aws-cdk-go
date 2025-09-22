package awsapigateway


// Supported IP Address Types.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
//   	EndpointConfiguration: &EndpointConfiguration{
//   		Types: []endpointType{
//   			apigateway.*endpointType_REGIONAL,
//   		},
//   		IpAddressType: apigateway.IpAddressType_DUAL_STACK,
//   	},
//   })
//
type IpAddressType string

const (
	// IPv4 address type.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// IPv4 and IPv6 address type.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

