package awsappflow


// The properties that are applied when Amazon Lookout for Metrics is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lookoutMetricsDestinationPropertiesProperty := &lookoutMetricsDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_LookoutMetricsDestinationPropertiesProperty struct {
	// The object specified in the Amazon Lookout for Metrics flow destination.
	Object *string `field:"optional" json:"object" yaml:"object"`
}

