package awsdatasync


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-outputtype
	//
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-objectversionids
	//
	ObjectVersionIds *string `field:"optional" json:"objectVersionIds" yaml:"objectVersionIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskreportconfig.html#cfn-datasync-task-taskreportconfig-reportlevel
	//
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

