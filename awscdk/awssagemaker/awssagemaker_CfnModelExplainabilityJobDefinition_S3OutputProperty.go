package awssagemaker


// The Amazon S3 storage location where the results of a monitoring job are saved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &s3OutputProperty{
//   	localPath: jsii.String("localPath"),
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	s3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnModelExplainabilityJobDefinition_S3OutputProperty struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

