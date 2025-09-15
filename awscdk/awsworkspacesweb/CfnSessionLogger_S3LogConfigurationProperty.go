package awsworkspacesweb


// The S3 log configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogConfigurationProperty := &S3LogConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	FolderStructure: jsii.String("folderStructure"),
//   	LogFileFormat: jsii.String("logFileFormat"),
//
//   	// the properties below are optional
//   	BucketOwner: jsii.String("bucketOwner"),
//   	KeyPrefix: jsii.String("keyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html
//
type CfnSessionLogger_S3LogConfigurationProperty struct {
	// The S3 bucket name where logs are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The folder structure that defines the organizational structure for log files in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-folderstructure
	//
	FolderStructure *string `field:"required" json:"folderStructure" yaml:"folderStructure"`
	// The format of the LogFile that is written to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-logfileformat
	//
	LogFileFormat *string `field:"required" json:"logFileFormat" yaml:"logFileFormat"`
	// The expected bucket owner of the target S3 bucket.
	//
	// The caller must have permissions to write to the target bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The S3 path prefix that determines where log files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

