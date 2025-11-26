package previewawsdatasyncmixins


// Configures a manifest, which is a list of files or objects that you want AWS DataSync to transfer.
//
// For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   manifestConfigProperty := &ManifestConfigProperty{
//   	Action: jsii.String("action"),
//   	Format: jsii.String("format"),
//   	Source: &SourceProperty{
//   		S3: &ManifestConfigSourceS3Property{
//   			BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   			ManifestObjectPath: jsii.String("manifestObjectPath"),
//   			ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   			S3BucketArn: jsii.String("s3BucketArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfig.html
//
type CfnTaskPropsMixin_ManifestConfigProperty struct {
	// Specifies what DataSync uses the manifest for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfig.html#cfn-datasync-task-manifestconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies the file format of your manifest.
	//
	// For more information, see [Creating a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html#transferring-with-manifest-create) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfig.html#cfn-datasync-task-manifestconfig-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Specifies the manifest that you want DataSync to use and where it's hosted.
	//
	// > You must specify this parameter if you're configuring a new manifest on or after February 7, 2024.
	// >
	// > If you don't, you'll get a 400 status code and `ValidationException` error stating that you're missing the IAM role for DataSync to access the S3 bucket where you're hosting your manifest. For more information, see [Providing DataSync access to your manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html#transferring-with-manifest-access) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfig.html#cfn-datasync-task-manifestconfig-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

