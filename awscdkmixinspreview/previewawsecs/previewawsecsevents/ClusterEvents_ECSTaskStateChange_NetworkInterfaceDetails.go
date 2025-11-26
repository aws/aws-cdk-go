package previewawsecsevents


// Type definition for NetworkInterfaceDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceDetails := &NetworkInterfaceDetails{
//   	AttachmentId: []*string{
//   		jsii.String("attachmentId"),
//   	},
//   	Ipv6Address: []*string{
//   		jsii.String("ipv6Address"),
//   	},
//   	PrivateIpv4Address: []*string{
//   		jsii.String("privateIpv4Address"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_NetworkInterfaceDetails struct {
	// attachmentId property.
	//
	// Specify an array of string values to match this event if the actual value of attachmentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AttachmentId *[]*string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// ipv6Address property.
	//
	// Specify an array of string values to match this event if the actual value of ipv6Address is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ipv6Address *[]*string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// privateIpv4Address property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpv4Address is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpv4Address *[]*string `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
}

