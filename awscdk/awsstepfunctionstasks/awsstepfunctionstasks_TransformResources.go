package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// ML compute instances for the transform job.
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
// Experimental.
type TransformResources struct {
	// Number of ML compute instances to use in the transform job.
	// Experimental.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// ML compute instance type for the transform job.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// AWS KMS key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s).
	// Experimental.
	VolumeEncryptionKey awskms.IKey `field:"optional" json:"volumeEncryptionKey" yaml:"volumeEncryptionKey"`
}

