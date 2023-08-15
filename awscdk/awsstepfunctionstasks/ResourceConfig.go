package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
//   	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("BlazingText"),
//   		TrainingInputMode: tasks.InputMode_FILE,
//   	},
//   	InputDataConfig: []channel{
//   		&channel{
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
//   		S3OutputLocation: tasks.S3Location_FromBucket(s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
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
type ResourceConfig struct {
	// The number of ML compute instances to use.
	// Default: 1 instance.
	//
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// ML compute instance type.
	//
	// To provide an instance type from the task input, supply an instance type in the following way
	// where the value in the task input is an EC2 instance type prepended with "ml.":
	//
	// ```ts
	// new ec2.InstanceType(sfn.JsonPath.stringAt('$.path.to.instanceType'));
	// ```.
	// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceConfig.html#sagemaker-Type-ResourceConfig-InstanceType
	//
	// Default: ec2.InstanceType(ec2.InstanceClass.M4, ec2.InstanceType.XLARGE)
	//
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Size of the ML storage volume that you want to provision.
	// Default: 10 GB EBS volume.
	//
	VolumeSize awscdk.Size `field:"required" json:"volumeSize" yaml:"volumeSize"`
	// KMS key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the training job.
	// Default: - Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	//
	VolumeEncryptionKey awskms.IKey `field:"optional" json:"volumeEncryptionKey" yaml:"volumeEncryptionKey"`
}

