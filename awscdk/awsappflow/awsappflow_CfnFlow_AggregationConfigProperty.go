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
//   	targetFileSize: jsii.Number(123),
//   }
//
type CfnFlow_AggregationConfigProperty struct {
	// Specifies whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated.
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// The desired file size, in MB, for each output file that Amazon AppFlow writes to the flow destination.
	//
	// For each file, Amazon AppFlow attempts to achieve the size that you specify. The actual file sizes might differ from this target based on the number and size of the records that each file contains.
	TargetFileSize *float64 `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
}

