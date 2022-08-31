package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Constructs `IS3Location` objects.
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
type S3Location interface {
	// Called when the S3Location is bound to a StepFunctions task.
	// Experimental.
	Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig
}

// The jsii proxy struct for S3Location
type jsiiProxy_S3Location struct {
	_ byte // padding
}

// Experimental.
func NewS3Location_Override(s S3Location) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		nil, // no parameters
		s,
	)
}

// An `IS3Location` built with a determined bucket and key prefix.
// Experimental.
func S3Location_FromBucket(bucket awss3.IBucket, keyPrefix *string) S3Location {
	_init_.Initialize()

	var returns S3Location

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		"fromBucket",
		[]interface{}{bucket, keyPrefix},
		&returns,
	)

	return returns
}

// An `IS3Location` determined fully by a JSON Path from the task input.
//
// Due to the dynamic nature of those locations, the IAM grants that will be set by `grantRead` and `grantWrite`
// apply to the `*` resource.
// Experimental.
func S3Location_FromJsonExpression(expression *string) S3Location {
	_init_.Initialize()

	var returns S3Location

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		"fromJsonExpression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Location) Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig {
	var returns *S3LocationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{task, opts},
		&returns,
	)

	return returns
}

