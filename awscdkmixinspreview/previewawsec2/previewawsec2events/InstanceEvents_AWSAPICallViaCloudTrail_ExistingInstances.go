package previewawsec2events


// Type definition for ExistingInstances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   existingInstances := &ExistingInstances{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	MarketOption: []*string{
//   		jsii.String("marketOption"),
//   	},
//   	OperatingSystem: []*string{
//   		jsii.String("operatingSystem"),
//   	},
//   	Tag: []*string{
//   		jsii.String("tag"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_ExistingInstances struct {
	// AvailabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of AvailabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Count property.
	//
	// Specify an array of string values to match this event if the actual value of Count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
	// InstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// MarketOption property.
	//
	// Specify an array of string values to match this event if the actual value of MarketOption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MarketOption *[]*string `field:"optional" json:"marketOption" yaml:"marketOption"`
	// OperatingSystem property.
	//
	// Specify an array of string values to match this event if the actual value of OperatingSystem is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OperatingSystem *[]*string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// tag property.
	//
	// Specify an array of string values to match this event if the actual value of tag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
}

