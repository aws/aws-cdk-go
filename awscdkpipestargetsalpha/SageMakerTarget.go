package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to a SageMaker pipeline.
//
// Example:
//   var sourceQueue Queue
//   var targetPipeline IPipeline
//
//
//   pipelineTarget := targets.NewSageMakerTarget(targetPipeline)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipelineTarget,
//   })
//
// Experimental.
type SageMakerTarget interface {
	awscdkpipesalpha.ITarget
	// The ARN of the target resource.
	// Experimental.
	TargetArn() *string
	// Bind this target to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig
	// Grant the pipe role to push to the target.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for SageMakerTarget
type jsiiProxy_SageMakerTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_SageMakerTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerTarget(pipeline awssagemaker.IPipeline, parameters *SageMakerTargetParameters) SageMakerTarget {
	_init_.Initialize()

	if err := validateNewSageMakerTargetParameters(pipeline, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SageMakerTarget",
		[]interface{}{pipeline, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerTarget_Override(s SageMakerTarget, pipeline awssagemaker.IPipeline, parameters *SageMakerTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SageMakerTarget",
		[]interface{}{pipeline, parameters},
		s,
	)
}

func (s *jsiiProxy_SageMakerTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerTarget) GrantPush(grantee awsiam.IRole) {
	if err := s.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee},
	)
}

