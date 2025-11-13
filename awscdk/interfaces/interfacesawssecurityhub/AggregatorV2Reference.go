package interfacesawssecurityhub


// A reference to a AggregatorV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregatorV2Reference := &AggregatorV2Reference{
//   	AggregatorV2Arn: jsii.String("aggregatorV2Arn"),
//   }
//
type AggregatorV2Reference struct {
	// The AggregatorV2Arn of the AggregatorV2 resource.
	AggregatorV2Arn *string `field:"required" json:"aggregatorV2Arn" yaml:"aggregatorV2Arn"`
}

