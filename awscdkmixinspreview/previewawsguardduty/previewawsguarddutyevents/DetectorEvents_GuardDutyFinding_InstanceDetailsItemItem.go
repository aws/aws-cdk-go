package previewawsguarddutyevents


// Type definition for InstanceDetailsItemItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceDetailsItemItem := &InstanceDetailsItemItem{
//   	PrivateDnsName: []*string{
//   		jsii.String("privateDnsName"),
//   	},
//   	PrivateIpAddress: []*string{
//   		jsii.String("privateIpAddress"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_InstanceDetailsItemItem struct {
	// privateDnsName property.
	//
	// Specify an array of string values to match this event if the actual value of privateDnsName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateDnsName *[]*string `field:"optional" json:"privateDnsName" yaml:"privateDnsName"`
	// privateIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of privateIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivateIpAddress *[]*string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

