package awssagemaker


// Specifies options when sharing an Amazon SageMaker Studio notebook.
//
// These settings are specified as part of `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called, and as part of `UserSettings` when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharingSettingsProperty := &SharingSettingsProperty{
//   	NotebookOutputOption: jsii.String("notebookOutputOption"),
//   	S3KmsKeyId: jsii.String("s3KmsKeyId"),
//   	S3OutputPath: jsii.String("s3OutputPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-sharingsettings.html
//
type CfnDomain_SharingSettingsProperty struct {
	// Whether to include the notebook cell output when sharing the notebook.
	//
	// The default is `Disabled` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-sharingsettings.html#cfn-sagemaker-domain-sharingsettings-notebookoutputoption
	//
	NotebookOutputOption *string `field:"optional" json:"notebookOutputOption" yaml:"notebookOutputOption"`
	// When `NotebookOutputOption` is `Allowed` , the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-sharingsettings.html#cfn-sagemaker-domain-sharingsettings-s3kmskeyid
	//
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
	// When `NotebookOutputOption` is `Allowed` , the Amazon S3 bucket used to store the shared notebook snapshots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-sharingsettings.html#cfn-sagemaker-domain-sharingsettings-s3outputpath
	//
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

