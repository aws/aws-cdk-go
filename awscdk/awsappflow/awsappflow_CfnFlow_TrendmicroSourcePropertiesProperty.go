package awsappflow


// The properties that are applied when using Trend Micro as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trendmicroSourcePropertiesProperty := &trendmicroSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_TrendmicroSourcePropertiesProperty struct {
	// The object specified in the Trend Micro flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

