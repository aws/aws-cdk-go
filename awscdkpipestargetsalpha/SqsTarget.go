package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// A EventBridge Pipes target that sends messages to an SQS queue.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   pipeSource := sources.NewSqsSource(sourceQueue, &SqsSourceParameters{
//   	BatchSize: jsii.Number(10),
//   	MaximumBatchingWindow: cdk.Duration_Seconds(jsii.Number(10)),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type SqsTarget interface {
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

// The jsii proxy struct for SqsTarget
type jsiiProxy_SqsTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_SqsTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsTarget(queue awssqs.IQueue, parameters *SqsTargetParameters) SqsTarget {
	_init_.Initialize()

	if err := validateNewSqsTargetParameters(queue, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTarget",
		[]interface{}{queue, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsTarget_Override(s SqsTarget, queue awssqs.IQueue, parameters *SqsTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.SqsTarget",
		[]interface{}{queue, parameters},
		s,
	)
}

func (s *jsiiProxy_SqsTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
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

func (s *jsiiProxy_SqsTarget) GrantPush(grantee awsiam.IRole) {
	if err := s.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee},
	)
}

