package awsnovaact


// Configuration settings for exporting workflow execution data and logs to Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowExportConfigProperty := &WorkflowExportConfigProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-novaact-workflowdefinition-workflowexportconfig.html
//
type CfnWorkflowDefinition_WorkflowExportConfigProperty struct {
	// The name of the Amazon S3 bucket for exporting workflow data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-novaact-workflowdefinition-workflowexportconfig.html#cfn-novaact-workflowdefinition-workflowexportconfig-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// An optional prefix for Amazon S3 object keys to organize exported data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-novaact-workflowdefinition-workflowexportconfig.html#cfn-novaact-workflowdefinition-workflowexportconfig-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

