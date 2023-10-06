package awsdatasync


// Specifies how you want to configure a task report, which provides detailed information about for your AWS DataSync transfer.
//
// For more information, see [Task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskReportConfigProperty := &TaskReportConfigProperty{
//   	Destination: &DestinationProperty{
//   		S3: &S3Property{
//   			BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   			S3BucketArn: jsii.String("s3BucketArn"),
//   			Subdirectory: jsii.String("subdirectory"),
//   		},
//   	},
//   	OutputType: jsii.String("outputType"),
//
//   	// the properties below are optional
//   	ObjectVersionIds: jsii.String("objectVersionIds"),
//   	Overrides: &OverridesProperty{
//   		Deleted: &DeletedProperty{
//   			ReportLevel: jsii.String("reportLevel"),
//   		},
//   		Skipped: &SkippedProperty{
//   			ReportLevel: jsii.String("reportLevel"),
//   		},
//   		Transferred: &TransferredProperty{
//   			ReportLevel: jsii.String("reportLevel"),
//   		},
//   		Verified: &VerifiedProperty{
//   			ReportLevel: jsii.String("reportLevel"),
//   		},
//   	},
//   	ReportLevel: jsii.String("reportLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html
//
type CfnTask_TaskReportConfigProperty struct {
	// Specifies the Amazon S3 bucket where DataSync uploads your task report.
	//
	// For more information, see [Task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html#task-report-access) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies the type of task report that you want:.
	//
	// - `SUMMARY_ONLY` : Provides necessary details about your task, including the number of files, objects, and directories transferred and transfer duration.
	// - `STANDARD` : Provides complete details about your task, including a full list of files, objects, and directories that were transferred, skipped, verified, and more.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-outputtype
	//
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Specifies whether your task report includes the new version of each object transferred into an S3 bucket.
	//
	// This only applies if you [enable versioning on your bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) . Keep in mind that setting this to `INCLUDE` can increase the duration of your task execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-objectversionids
	//
	ObjectVersionIds *string `field:"optional" json:"objectVersionIds" yaml:"objectVersionIds"`
	// Customizes the reporting level for aspects of your task report.
	//
	// For example, your report might generally only include errors, but you could specify that you want a list of successes and errors just for the files that DataSync attempted to delete in your destination location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
	//
	// - `ERRORS_ONLY` : A report shows what DataSync was unable to transfer, skip, verify, and delete.
	// - `SUCCESSES_AND_ERRORS` : A report shows what DataSync was able and unable to transfer, skip, verify, and delete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-reportlevel
	//
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

