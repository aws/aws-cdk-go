package previewawsec2events


// Type definition for PrivateIpAddressesSet_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateIpAddressesSet2 := &PrivateIpAddressesSet2{
//   	Item: []PrivateIpAddressesSet1Item{
//   		&PrivateIpAddressesSet1Item{
//   			Primary: []*string{
//   				jsii.String("primary"),
//   			},
//   			PrivateIpAddress: []*string{
//   				jsii.String("privateIpAddress"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_PrivateIpAddressesSet2 struct {
	// item property.
	//
	// Specify an array of string values to match this event if the actual value of item is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Item *[]*InstanceEvents_AWSAPICallViaCloudTrail_PrivateIpAddressesSet1Item `field:"optional" json:"item" yaml:"item"`
}

