package previewawsec2events


// Type definition for SpotOptions_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotOptions2 := &SpotOptions2{
//   	MaxPrice: []*string{
//   		jsii.String("maxPrice"),
//   	},
//   	SpotInstanceType: []*string{
//   		jsii.String("spotInstanceType"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions2 struct {
	// MaxPrice property.
	//
	// Specify an array of string values to match this event if the actual value of MaxPrice is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxPrice *[]*string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// SpotInstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of SpotInstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotInstanceType *[]*string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
}

