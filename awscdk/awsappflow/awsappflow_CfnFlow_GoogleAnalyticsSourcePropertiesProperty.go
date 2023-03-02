package awsappflow


// The properties that are applied when Google Analytics is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   googleAnalyticsSourcePropertiesProperty := &googleAnalyticsSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_GoogleAnalyticsSourcePropertiesProperty struct {
	// The object specified in the Google Analytics flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

