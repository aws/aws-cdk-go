package awssagemaker


// Configuration for monitoring constraints and monitoring statistics.
//
// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityBaselineConfigProperty := &dataQualityBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   	statisticsResource: &statisticsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_DataQualityBaselineConfigProperty struct {
	// The name of the job that performs baselining for the data quality monitoring job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// Configuration for monitoring constraints and monitoring statistics.
	//
	// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

