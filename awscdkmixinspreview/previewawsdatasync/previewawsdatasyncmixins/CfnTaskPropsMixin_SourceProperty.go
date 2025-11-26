package previewawsdatasyncmixins


// Specifies the manifest that you want DataSync to use and where it's hosted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceProperty := &SourceProperty{
//   	S3: &ManifestConfigSourceS3Property{
//   		BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   		ManifestObjectPath: jsii.String("manifestObjectPath"),
//   		ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   		S3BucketArn: jsii.String("s3BucketArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-source.html
//
type CfnTaskPropsMixin_SourceProperty struct {
	// Specifies the S3 bucket where you're hosting the manifest that you want AWS DataSync to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-source.html#cfn-datasync-task-source-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

