package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   percentileAggregationProperty := &PercentileAggregationProperty{
//   	PercentileValue: jsii.Number(123),
//   }
//
type CfnAnalysis_PercentileAggregationProperty struct {
	// `CfnAnalysis.PercentileAggregationProperty.PercentileValue`.
	PercentileValue *float64 `field:"optional" json:"percentileValue" yaml:"percentileValue"`
}

