package awssagemaker


// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineConfigProperty := &BaselineConfigProperty{
//   	ConstraintsResource: &ConstraintsResourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	StatisticsResource: &StatisticsResourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-baselineconfig.html
//
type CfnMonitoringSchedule_BaselineConfigProperty struct {
	// The Amazon S3 URI for the constraints resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-baselineconfig.html#cfn-sagemaker-monitoringschedule-baselineconfig-constraintsresource
	//
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-baselineconfig.html#cfn-sagemaker-monitoringschedule-baselineconfig-statisticsresource
	//
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

