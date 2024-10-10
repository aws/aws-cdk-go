package awscdkpipessourcesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/internal"
)

// A source that reads from an SQS queue.
//
// Example:
//   var sourceQueue queue
//   var targetStateMachine iStateMachine
//
//
//   pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
//   	InvocationType: targets.StateMachineInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type SqsSource interface {
	awscdkpipesalpha.ISource
	// The ARN of the source resource.
	// Experimental.
	SourceArn() *string
	// Bind the source to a pipe.
	// Experimental.
	Bind(_pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig
	// Grant the pipe role read access to the source.
	// Experimental.
	GrantRead(grantee awsiam.IRole)
}

// The jsii proxy struct for SqsSource
type jsiiProxy_SqsSource struct {
	internal.Type__awscdkpipesalphaISource
}

func (j *jsiiProxy_SqsSource) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsSource(queue awssqs.IQueue, parameters *SqsSourceParameters) SqsSource {
	_init_.Initialize()

	if err := validateNewSqsSourceParameters(queue, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsSource{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSource",
		[]interface{}{queue, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsSource_Override(s SqsSource, queue awssqs.IQueue, parameters *SqsSourceParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.SqsSource",
		[]interface{}{queue, parameters},
		s,
	)
}

func (s *jsiiProxy_SqsSource) Bind(_pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig {
	if err := s.validateBindParameters(_pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSource) GrantRead(grantee awsiam.IRole) {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantRead",
		[]interface{}{grantee},
	)
}

