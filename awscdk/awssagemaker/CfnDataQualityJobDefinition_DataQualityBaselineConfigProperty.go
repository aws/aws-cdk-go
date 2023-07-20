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
//   dataQualityBaselineConfigProperty := &DataQualityBaselineConfigProperty{
//   	BaseliningJobName: jsii.String("baseliningJobName"),
//   	ConstraintsResource: &ConstraintsResourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	StatisticsResource: &StatisticsResourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html
//
type CfnDataQualityJobDefinition_DataQualityBaselineConfigProperty struct {
	// The name of the job that performs baselining for the data quality monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-baseliningjobname
	//
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-constraintsresource
	//
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// Configuration for monitoring constraints and monitoring statistics.
	//
	// These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-statisticsresource
	//
	StatisticsResource interface{} `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

