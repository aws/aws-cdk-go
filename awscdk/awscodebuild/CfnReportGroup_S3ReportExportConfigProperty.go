package awscodebuild


// Information about the S3 bucket where the raw data of a report are exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReportExportConfigProperty := &S3ReportExportConfigProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	BucketOwner: jsii.String("bucketOwner"),
//   	EncryptionDisabled: jsii.Boolean(false),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Packaging: jsii.String("packaging"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html
//
type CfnReportGroup_S3ReportExportConfigProperty struct {
	// The name of the S3 bucket where the raw data of a report are exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account identifier of the owner of the Amazon S3 bucket.
	//
	// This allows report data to be exported to an Amazon S3 bucket that is owned by an account other than the account running the build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// A boolean value that specifies if the results of a report are encrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-encryptiondisabled
	//
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// The encryption key for the report's encrypted raw data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The type of build output artifact to create. Valid values include:.
	//
	// - `NONE` : CodeBuild creates the raw data in the output bucket. This is the default if packaging is not specified.
	// - `ZIP` : CodeBuild creates a ZIP file with the raw data in the output bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-packaging
	//
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// The path to the exported report's raw data results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-reportgroup-s3reportexportconfig.html#cfn-codebuild-reportgroup-s3reportexportconfig-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

