package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// S3 location where you want Amazon SageMaker to save the results from the transform job.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &SageMakerCreateTransformJobProps{
//   	TransformJobName: jsii.String("MyTransformJob"),
//   	ModelName: jsii.String("MyModelName"),
//   	ModelClientOptions: &ModelClientOptions{
//   		InvocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		InvocationsTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   	TransformInput: &TransformInput{
//   		TransformDataSource: &TransformDataSource{
//   			S3DataSource: &TransformS3DataSource{
//   				S3Uri: jsii.String("s3://inputbucket/train"),
//   				S3DataType: tasks.S3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M4, ec2.InstanceSize_XLARGE),
//   	},
//   })
//
type TransformOutput struct {
	// S3 path where you want Amazon SageMaker to store the results of the transform job.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// MIME type used to specify the output data.
	// Default: - None.
	//
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Defines how to assemble the results of the transform job as a single S3 object.
	// Default: - None.
	//
	AssembleWith AssembleWith `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// AWS KMS key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// Default: - default KMS key for Amazon S3 for your role's account.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

