package awsappflow


// The aggregation settings that you can use to customize the output format of your flow data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationConfigProperty := &aggregationConfigProperty{
//   	aggregationType: jsii.String("aggregationType"),
//   }
//
type CfnFlow_AggregationConfigProperty struct {
	// Specifies whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated.
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
}

