package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Configures the timeout and maximum number of retries for processing a transform job invocation.
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
// Experimental.
type ModelClientOptions struct {
	// The maximum number of retries when invocation requests are failing.
	// Experimental.
	InvocationsMaxRetries *float64 `field:"optional" json:"invocationsMaxRetries" yaml:"invocationsMaxRetries"`
	// The timeout duration for an invocation request.
	// Experimental.
	InvocationsTimeout awscdk.Duration `field:"optional" json:"invocationsTimeout" yaml:"invocationsTimeout"`
}

