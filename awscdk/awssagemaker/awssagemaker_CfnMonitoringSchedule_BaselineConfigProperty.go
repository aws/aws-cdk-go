package awssagemaker


// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineConfigProperty := &baselineConfigProperty{
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   	statisticsResource: &statisticsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnMonitoringSchedule_BaselineConfigProperty struct {
	// The Amazon S3 URI for the constraints resource.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

