package previewawsec2events


// Type definition for NetworkInterfaceSetItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceSetItem := &NetworkInterfaceSetItem{
//   	DeviceIndex: []*string{
//   		jsii.String("deviceIndex"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterfaceSetItem struct {
	// deviceIndex property.
	//
	// Specify an array of string values to match this event if the actual value of deviceIndex is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeviceIndex *[]*string `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// subnetId property.
	//
	// Specify an array of string values to match this event if the actual value of subnetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

