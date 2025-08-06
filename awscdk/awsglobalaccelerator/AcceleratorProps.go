package awsglobalaccelerator


// Construct properties of the Accelerator.
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
type AcceleratorProps struct {
	// The name of the accelerator.
	// Default: - resource ID.
	//
	AcceleratorName *string `field:"optional" json:"acceleratorName" yaml:"acceleratorName"`
	// Indicates whether the accelerator is enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// IP addresses associated with the accelerator.
	//
	// Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose IP
	// addresses from your own pool to use for the accelerator's static IP addresses when you create an accelerator.
	// You can specify one or two addresses, separated by a comma. Do not include the /32 suffix.
	//
	// Only one IP address from each of your IP address ranges can be used for each accelerator. If you specify
	// only one IP address from your IP address range, Global Accelerator assigns a second static IP address for
	// the accelerator from the AWS IP address pool.
	//
	// Note that you can't update IP addresses for an existing accelerator. To change them, you must create a
	// new accelerator with the new addresses.
	// Default: - undefined. IP addresses will be from Amazon's pool of IP addresses.
	//
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// The IP address type that an accelerator supports.
	//
	// For a standard accelerator, the value can be IPV4 or DUAL_STACK.
	// Default: - "IPV4".
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

