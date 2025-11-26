package previewawsec2events


// Type definition for SpotOptions_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotOptions1 := &SpotOptions1{
//   	MaxPrice: []*string{
//   		jsii.String("maxPrice"),
//   	},
//   	SpotInstanceType: []*string{
//   		jsii.String("spotInstanceType"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions1 struct {
	// maxPrice property.
	//
	// Specify an array of string values to match this event if the actual value of maxPrice is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxPrice *[]*string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// spotInstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of spotInstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotInstanceType *[]*string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
}

