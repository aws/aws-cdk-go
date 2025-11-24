package mixinsawssagemaker


// Configuration for uploading output data to Amazon S3 from the processing container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3OutputProperty := &S3OutputProperty{
//   	LocalPath: jsii.String("localPath"),
//   	S3UploadMode: jsii.String("s3UploadMode"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3output.html
//
type CfnProcessingJobPropsMixin_S3OutputProperty struct {
	// The local path of a directory where you want Amazon SageMaker to upload its contents to Amazon S3.
	//
	// `LocalPath` is an absolute path to a directory containing output files. This directory will be created by the platform and exist when your container's entrypoint is invoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3output.html#cfn-sagemaker-processingjob-s3output-localpath
	//
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Whether to upload the results of the processing job continuously or after the job completes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3output.html#cfn-sagemaker-processingjob-s3output-s3uploadmode
	//
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
	// The URI of the Amazon S3 prefix Amazon SageMaker downloads data required to run a processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3output.html#cfn-sagemaker-processingjob-s3output-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

