package awscdkpipessourcesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// A source that reads from Kinesis.
//
// Example:
//   var sourceStream Stream
//   var targetQueue Queue
//
//
//   pipeSource := sources.NewKinesisSource(sourceStream, &KinesisSourceParameters{
//   	StartingPosition: sources.KinesisStartingPosition_LATEST,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type KinesisSource interface {
	StreamSource
	// The dead-letter SQS queue or SNS topic.
	// Experimental.
	DeadLetterTarget() interface{}
	// The ARN of the source resource.
	// Experimental.
	SourceArn() *string
	// Base parameters for streaming sources.
	// Experimental.
	SourceParameters() *StreamSourceParameters
	// Bind the source to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig
	// Retrieves the ARN from the dead-letter SQS queue or SNS topic.
	// Experimental.
	GetDeadLetterTargetArn(deadLetterTarget interface{}) *string
	// Grants the pipe role permission to publish to the dead-letter target.
	// Experimental.
	GrantPush(grantee awsiam.IRole, deadLetterTarget interface{})
	// Grant the pipe role read access to the source.
	// Experimental.
	GrantRead(grantee awsiam.IRole)
}

// The jsii proxy struct for KinesisSource
type jsiiProxy_KinesisSource struct {
	jsiiProxy_StreamSource
}

func (j *jsiiProxy_KinesisSource) DeadLetterTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisSource) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisSource) SourceParameters() *StreamSourceParameters {
	var returns *StreamSourceParameters
	_jsii_.Get(
		j,
		"sourceParameters",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisSource(stream awskinesis.IStream, parameters *KinesisSourceParameters) KinesisSource {
	_init_.Initialize()

	if err := validateNewKinesisSourceParameters(stream, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisSource{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisSource",
		[]interface{}{stream, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisSource_Override(k KinesisSource, stream awskinesis.IStream, parameters *KinesisSourceParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisSource",
		[]interface{}{stream, parameters},
		k,
	)
}

// Determines if the source is an instance of SourceWithDeadLetterTarget.
// Experimental.
func KinesisSource_IsSourceWithDeadLetterTarget(source awscdkpipesalpha.ISource) *bool {
	_init_.Initialize()

	if err := validateKinesisSource_IsSourceWithDeadLetterTargetParameters(source); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-sources-alpha.KinesisSource",
		"isSourceWithDeadLetterTarget",
		[]interface{}{source},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisSource) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig {
	if err := k.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.SourceConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisSource) GetDeadLetterTargetArn(deadLetterTarget interface{}) *string {
	if err := k.validateGetDeadLetterTargetArnParameters(deadLetterTarget); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getDeadLetterTargetArn",
		[]interface{}{deadLetterTarget},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisSource) GrantPush(grantee awsiam.IRole, deadLetterTarget interface{}) {
	if err := k.validateGrantPushParameters(grantee, deadLetterTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"grantPush",
		[]interface{}{grantee, deadLetterTarget},
	)
}

func (k *jsiiProxy_KinesisSource) GrantRead(grantee awsiam.IRole) {
	if err := k.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"grantRead",
		[]interface{}{grantee},
	)
}

