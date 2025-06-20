package awscodebuild


// Information about the location where the run of a report is exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportExportConfigProperty := &ReportExportConfigProperty{
//   	ExportConfigType: jsii.String("exportConfigType"),
//
//   	// the properties below are optional
//   	S3Destination: &S3ReportExportConfigProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		EncryptionDisabled: jsii.Boolean(false),
//   		EncryptionKey: jsii.String("encryptionKey"),
//   		Packaging: jsii.String("packaging"),
//   		Path: jsii.String("path"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-reportexportconfig.html
//
type CfnReportGroup_ReportExportConfigProperty struct {
	// The export configuration type. Valid values are:.
	//
	// - `S3` : The report results are exported to an S3 bucket.
	// - `NO_EXPORT` : The report results are not exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-reportexportconfig.html#cfn-codebuild-reportgroup-reportexportconfig-exportconfigtype
	//
	ExportConfigType *string `field:"required" json:"exportConfigType" yaml:"exportConfigType"`
	// A `S3ReportExportConfig` object that contains information about the S3 bucket where the run of a report is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-reportexportconfig.html#cfn-codebuild-reportgroup-reportexportconfig-s3destination
	//
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

