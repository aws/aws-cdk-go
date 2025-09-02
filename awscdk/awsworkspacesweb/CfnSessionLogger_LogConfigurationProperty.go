package awsworkspacesweb


// The configuration of the log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	S3: &S3LogConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		FolderStructure: jsii.String("folderStructure"),
//   		LogFileFormat: jsii.String("logFileFormat"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-logconfiguration.html
//
type CfnSessionLogger_LogConfigurationProperty struct {
	// The configuration for delivering the logs to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-logconfiguration.html#cfn-workspacesweb-sessionlogger-logconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

