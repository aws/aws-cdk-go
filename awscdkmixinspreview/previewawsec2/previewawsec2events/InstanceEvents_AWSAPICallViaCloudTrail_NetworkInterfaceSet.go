package previewawsec2events


// Type definition for NetworkInterfaceSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceSet := &NetworkInterfaceSet{
//   	Items: []NetworkInterfaceSetItem{
//   		&NetworkInterfaceSetItem{
//   			DeviceIndex: []*string{
//   				jsii.String("deviceIndex"),
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSet struct {
	// items property.
	//
	// Specify an array of string values to match this event if the actual value of items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Items *[]*InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSetItem `field:"optional" json:"items" yaml:"items"`
}

