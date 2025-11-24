package mixinsawsworkspacesweb


// The S3 log configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LogConfigurationProperty := &S3LogConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	BucketOwner: jsii.String("bucketOwner"),
//   	FolderStructure: jsii.String("folderStructure"),
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	LogFileFormat: jsii.String("logFileFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html
//
type CfnSessionLoggerPropsMixin_S3LogConfigurationProperty struct {
	// The S3 bucket name where logs are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The expected bucket owner of the target S3 bucket.
	//
	// The caller must have permissions to write to the target bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The folder structure that defines the organizational structure for log files in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-folderstructure
	//
	FolderStructure *string `field:"optional" json:"folderStructure" yaml:"folderStructure"`
	// The S3 path prefix that determines where log files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The format of the LogFile that is written to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html#cfn-workspacesweb-sessionlogger-s3logconfiguration-logfileformat
	//
	LogFileFormat *string `field:"optional" json:"logFileFormat" yaml:"logFileFormat"`
}

