package previewawsec2events


// Type definition for InstanceMarketOptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceMarketOptions := &InstanceMarketOptions{
//   	MarketType: []*string{
//   		jsii.String("marketType"),
//   	},
//   	SpotOptions: &SpotOptions1{
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
type InstanceEvents_AWSAPICallViaCloudTrail_InstanceMarketOptions struct {
	// marketType property.
	//
	// Specify an array of string values to match this event if the actual value of marketType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MarketType *[]*string `field:"optional" json:"marketType" yaml:"marketType"`
	// spotOptions property.
	//
	// Specify an array of string values to match this event if the actual value of spotOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotOptions *InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions1 `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

