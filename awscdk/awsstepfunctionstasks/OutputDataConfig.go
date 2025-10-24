package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Configures the S3 bucket where SageMaker will save the result of model training.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
//   	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("BlazingText"),
//   		TrainingInputMode: tasks.InputMode_FILE,
//   	},
//   	InputDataConfig: []Channel{
//   		&Channel{
//   			ChannelName: jsii.String("train"),
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3DataType: tasks.S3DataType_S3_PREFIX,
//   					S3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3OutputLocation: tasks.S3Location_FromBucket(s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket")), jsii.String("myoutputpath")),
//   	},
//   	ResourceConfig: &ResourceConfig{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.NewInstanceType(sfn.JsonPath_*StringAt(jsii.String("$.InstanceType"))),
//   		VolumeSize: awscdk.Size_Gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntime: awscdk.Duration_Hours(jsii.Number(2)),
//   	},
//   })
//
type OutputDataConfig struct {
	// Identifies the S3 path where you want Amazon SageMaker to store the model artifacts.
	S3OutputLocation S3Location `field:"required" json:"s3OutputLocation" yaml:"s3OutputLocation"`
	// Optional KMS encryption key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// Default: - Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

