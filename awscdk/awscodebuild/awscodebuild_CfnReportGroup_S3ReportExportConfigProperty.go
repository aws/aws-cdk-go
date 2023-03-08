package awscodebuild


// Information about the S3 bucket where the raw data of a report are exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReportExportConfigProperty := &s3ReportExportConfigProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketOwner: jsii.String("bucketOwner"),
//   	encryptionDisabled: jsii.Boolean(false),
//   	encryptionKey: jsii.String("encryptionKey"),
//   	packaging: jsii.String("packaging"),
//   	path: jsii.String("path"),
//   }
//
type CfnReportGroup_S3ReportExportConfigProperty struct {
	// The name of the S3 bucket where the raw data of a report are exported.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account identifier of the owner of the Amazon S3 bucket.
	//
	// This allows report data to be exported to an Amazon S3 bucket that is owned by an account other than the account running the build.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// A boolean value that specifies if the results of a report are encrypted.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// The encryption key for the report's encrypted raw data.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The type of build output artifact to create. Valid values include:.
	//
	// - `NONE` : CodeBuild creates the raw data in the output bucket. This is the default if packaging is not specified.
	// - `ZIP` : CodeBuild creates a ZIP file with the raw data in the output bucket.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// The path to the exported report's raw data results.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

