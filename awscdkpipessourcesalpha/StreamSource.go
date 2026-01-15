package awscdkpipessourcesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipessourcesalpha/v2/internal"
)

// Streaming sources.
// Experimental.
type StreamSource interface {
	awscdkpipesalpha.SourceWithDeadLetterTarget
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
	//
	// [disable-awslint:no-grants].
	// Experimental.
	GrantPush(grantee awsiam.IRole, deadLetterTarget interface{})
	// Grant the pipe role read access to the source.
	// Experimental.
	GrantRead(grantee awsiam.IRole)
}

// The jsii proxy struct for StreamSource
type jsiiProxy_StreamSource struct {
	internal.Type__awscdkpipesalphaSourceWithDeadLetterTarget
}

func (j *jsiiProxy_StreamSource) DeadLetterTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamSource) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamSource) SourceParameters() *StreamSourceParameters {
	var returns *StreamSourceParameters
	_jsii_.Get(
		j,
		"sourceParameters",
		&returns,
	)
	return returns
}


// Experimental.
func NewStreamSource_Override(s StreamSource, sourceArn *string, sourceParameters *StreamSourceParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-sources-alpha.StreamSource",
		[]interface{}{sourceArn, sourceParameters},
		s,
	)
}

// Determines if the source is an instance of SourceWithDeadLetterTarget.
// Experimental.
func StreamSource_IsSourceWithDeadLetterTarget(source awscdkpipesalpha.ISource) *bool {
	_init_.Initialize()

	if err := validateStreamSource_IsSourceWithDeadLetterTargetParameters(source); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-sources-alpha.StreamSource",
		"isSourceWithDeadLetterTarget",
		[]interface{}{source},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamSource) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.SourceConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamSource) GetDeadLetterTargetArn(deadLetterTarget interface{}) *string {
	if err := s.validateGetDeadLetterTargetArnParameters(deadLetterTarget); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getDeadLetterTargetArn",
		[]interface{}{deadLetterTarget},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamSource) GrantPush(grantee awsiam.IRole, deadLetterTarget interface{}) {
	if err := s.validateGrantPushParameters(grantee, deadLetterTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee, deadLetterTarget},
	)
}

func (s *jsiiProxy_StreamSource) GrantRead(grantee awsiam.IRole) {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantRead",
		[]interface{}{grantee},
	)
}

