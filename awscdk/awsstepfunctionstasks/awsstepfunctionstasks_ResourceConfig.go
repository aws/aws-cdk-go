package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &sageMakerCreateTrainingJobProps{
//   	trainingJobName: sfn.jsonPath.stringAt(jsii.String("$.JobName")),
//   	algorithmSpecification: &algorithmSpecification{
//   		algorithmName: jsii.String("BlazingText"),
//   		trainingInputMode: tasks.inputMode_FILE,
//   	},
//   	inputDataConfig: []channel{
//   		&channel{
//   			channelName: jsii.String("train"),
//   			dataSource: &dataSource{
//   				s3DataSource: &s3DataSource{
//   					s3DataType: tasks.s3DataType_S3_PREFIX,
//   					s3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	outputDataConfig: &outputDataConfig{
//   		s3OutputLocation: tasks.*s3Location.fromBucket(s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
//   	},
//   	resourceConfig: &resourceConfig{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.NewInstanceType(sfn.*jsonPath.stringAt(jsii.String("$.InstanceType"))),
//   		volumeSize: awscdk.Size.gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	stoppingCondition: &stoppingCondition{
//   		maxRuntime: awscdk.Duration.hours(jsii.Number(2)),
//   	},
//   })
//
// Experimental.
type ResourceConfig struct {
	// The number of ML compute instances to use.
	// Experimental.
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
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Size of the ML storage volume that you want to provision.
	// Experimental.
	VolumeSize awscdk.Size `field:"required" json:"volumeSize" yaml:"volumeSize"`
	// KMS key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the training job.
	// Experimental.
	VolumeEncryptionKey awskms.IKey `field:"optional" json:"volumeEncryptionKey" yaml:"volumeEncryptionKey"`
}

