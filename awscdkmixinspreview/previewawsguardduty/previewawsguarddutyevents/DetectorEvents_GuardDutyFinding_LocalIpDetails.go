package previewawsguarddutyevents


// Type definition for LocalIpDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localIpDetails := &LocalIpDetails{
//   	IpAddressV4: []*string{
//   		jsii.String("ipAddressV4"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_LocalIpDetails struct {
	// ipAddressV4 property.
	//
	// Specify an array of string values to match this event if the actual value of ipAddressV4 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IpAddressV4 *[]*string `field:"optional" json:"ipAddressV4" yaml:"ipAddressV4"`
}

