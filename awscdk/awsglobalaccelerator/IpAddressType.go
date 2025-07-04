package awsglobalaccelerator


// The IP address type that an accelerator supports.
//
// Example:
//   accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"), &AcceleratorProps{
//   	IpAddresses: []*string{
//   		jsii.String("1.1.1.1"),
//   		jsii.String("2.2.2.2"),
//   	},
//   	IpAddressType: globalaccelerator.IpAddressType_IPV4,
//   })
//
type IpAddressType string

const (
	// IPV4.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// DUAL_STACK.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

