package previewawsec2events


// Type definition for InstanceMarketOptions_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceMarketOptions1 := &InstanceMarketOptions1{
//   	MarketType: []*string{
//   		jsii.String("marketType"),
//   	},
//   	SpotOptions: &SpotOptions2{
//   		MaxPrice: []*string{
//   			jsii.String("maxPrice"),
//   		},
//   		SpotInstanceType: []*string{
//   			jsii.String("spotInstanceType"),
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_InstanceMarketOptions1 struct {
	// MarketType property.
	//
	// Specify an array of string values to match this event if the actual value of MarketType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MarketType *[]*string `field:"optional" json:"marketType" yaml:"marketType"`
	// SpotOptions property.
	//
	// Specify an array of string values to match this event if the actual value of SpotOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotOptions *InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions2 `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

