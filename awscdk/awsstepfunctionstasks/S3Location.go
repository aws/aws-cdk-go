package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Constructs `IS3Location` objects.
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
type S3Location interface {
	// Called when the S3Location is bound to a StepFunctions task.
	Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig
}

// The jsii proxy struct for S3Location
type jsiiProxy_S3Location struct {
	_ byte // padding
}

func NewS3Location_Override(s S3Location) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3Location",
		nil, // no parameters
		s,
	)
}

// An `IS3Location` built with a determined bucket and key prefix.
func S3Location_FromBucket(bucket awss3.IBucket, keyPrefix *string) S3Location {
	_init_.Initialize()

	if err := validateS3Location_FromBucketParameters(bucket, keyPrefix); err != nil {
		panic(err)
	}
	var returns S3Location

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3Location",
		"fromBucket",
		[]interface{}{bucket, keyPrefix},
		&returns,
	)

	return returns
}

// An `IS3Location` determined fully by a JSONata expression or JSON Path from the task input.
//
// Due to the dynamic nature of those locations, the IAM grants that will be set by `grantRead` and `grantWrite`
// apply to the `*` resource.
func S3Location_FromJsonExpression(expression *string) S3Location {
	_init_.Initialize()

	if err := validateS3Location_FromJsonExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns S3Location

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3Location",
		"fromJsonExpression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Location) Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig {
	if err := s.validateBindParameters(task, opts); err != nil {
		panic(err)
	}
	var returns *S3LocationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{task, opts},
		&returns,
	)

	return returns
}

