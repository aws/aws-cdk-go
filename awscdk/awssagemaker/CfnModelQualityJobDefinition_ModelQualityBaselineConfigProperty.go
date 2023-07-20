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
//   modelQualityBaselineConfigProperty := &ModelQualityBaselineConfigProperty{
//   	BaseliningJobName: jsii.String("baseliningJobName"),
//   	ConstraintsResource: &ConstraintsResourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig.html
//
type CfnModelQualityJobDefinition_ModelQualityBaselineConfigProperty struct {
	// The name of the job that performs baselining for the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig-baseliningjobname
	//
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig-constraintsresource
	//
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}

