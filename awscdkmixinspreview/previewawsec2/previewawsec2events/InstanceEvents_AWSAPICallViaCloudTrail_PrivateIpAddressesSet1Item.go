package previewawsec2events


// Type definition for PrivateIpAddressesSet_1Item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateIpAddressesSet1Item := &PrivateIpAddressesSet1Item{
//   	Primary: []*string{
//   		jsii.String("primary"),
//   	},
//   	PrivateIpAddress: []*string{
//   		jsii.String("privateIpAddress"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_PrivateIpAddressesSet1Item struct {
	// primary property.
	//
	// Specify an array of string values to match this event if the actual value of primary is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Primary *[]*string `field:"optional" json:"primary" yaml:"primary"`
	// privateIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddress *[]*string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

