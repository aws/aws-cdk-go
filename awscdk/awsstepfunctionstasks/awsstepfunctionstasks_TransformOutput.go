package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// S3 location where you want Amazon SageMaker to save the results from the transform job.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &sageMakerCreateTransformJobProps{
//   	transformJobName: jsii.String("MyTransformJob"),
//   	modelName: jsii.String("MyModelName"),
//   	modelClientOptions: &modelClientOptions{
//   		invocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		invocationsTimeout: awscdk.Duration.minutes(jsii.Number(5)),
//   	},
//   	transformInput: &transformInput{
//   		transformDataSource: &transformDataSource{
//   			s3DataSource: &transformS3DataSource{
//   				s3Uri: jsii.String("s3://inputbucket/train"),
//   				s3DataType: tasks.s3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	transformOutput: &transformOutput{
//   		s3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	transformResources: &transformResources{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_M4, ec2.instanceSize_XLARGE),
//   	},
//   })
//
type TransformOutput struct {
	// S3 path where you want Amazon SageMaker to store the results of the transform job.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// MIME type used to specify the output data.
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Defines how to assemble the results of the transform job as a single S3 object.
	AssembleWith AssembleWith `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// AWS KMS key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

