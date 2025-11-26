package previewawsworkspaceswebmixins


// The configuration of the log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	S3: &S3LogConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		FolderStructure: jsii.String("folderStructure"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   		LogFileFormat: jsii.String("logFileFormat"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-logconfiguration.html
//
type CfnSessionLoggerPropsMixin_LogConfigurationProperty struct {
	// The configuration for delivering the logs to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-sessionlogger-logconfiguration.html#cfn-workspacesweb-sessionlogger-logconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

