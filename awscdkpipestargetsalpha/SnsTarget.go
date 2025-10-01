package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// A EventBridge Pipes target that sends messages to an SNS topic.
//
// Example:
//   var sourceQueue queue
//   var targetTopic topic
//
//
//   pipeTarget := targets.NewSnsTarget(targetTopic)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type SnsTarget interface {
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

// The jsii proxy struct for SnsTarget
type jsiiProxy_SnsTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_SnsTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsTarget(topic awssns.ITopic, parameters *SnsTargetParameters) SnsTarget {
	_init_.Initialize()

	if err := validateNewSnsTargetParameters(topic, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SnsTarget",
		[]interface{}{topic, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTarget_Override(s SnsTarget, topic awssns.ITopic, parameters *SnsTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SnsTarget",
		[]interface{}{topic, parameters},
		s,
	)
}

func (s *jsiiProxy_SnsTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
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

func (s *jsiiProxy_SnsTarget) GrantPush(grantee awsiam.IRole) {
	if err := s.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee},
	)
}

