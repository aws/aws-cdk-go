package previewawsec2events


// Type definition for NetworkInterface_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterface1 := &NetworkInterface1{
//   	DeviceIndex: []*string{
//   		jsii.String("deviceIndex"),
//   	},
//   	SecurityGroupId: &SecurityGroupId{
//   		Content: []*string{
//   			jsii.String("content"),
//   		},
//   		Tag: []*string{
//   			jsii.String("tag"),
//   		},
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   	Tag: []*string{
//   		jsii.String("tag"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_NetworkInterface1 struct {
	// DeviceIndex property.
	//
	// Specify an array of string values to match this event if the actual value of DeviceIndex is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeviceIndex *[]*string `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// SecurityGroupId property.
	//
	// Specify an array of string values to match this event if the actual value of SecurityGroupId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityGroupId *InstanceEvents_AWSAPICallViaCloudTrail_SecurityGroupId `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// SubnetId property.
	//
	// Specify an array of string values to match this event if the actual value of SubnetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// tag property.
	//
	// Specify an array of string values to match this event if the actual value of tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
}

