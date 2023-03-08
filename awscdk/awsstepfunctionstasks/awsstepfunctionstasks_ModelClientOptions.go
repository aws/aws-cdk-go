package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Configures the timeout and maximum number of retries for processing a transform job invocation.
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
type ModelClientOptions struct {
	// The maximum number of retries when invocation requests are failing.
	// Experimental.
	InvocationsMaxRetries *float64 `field:"optional" json:"invocationsMaxRetries" yaml:"invocationsMaxRetries"`
	// The timeout duration for an invocation request.
	// Experimental.
	InvocationsTimeout awscdk.Duration `field:"optional" json:"invocationsTimeout" yaml:"invocationsTimeout"`
}

