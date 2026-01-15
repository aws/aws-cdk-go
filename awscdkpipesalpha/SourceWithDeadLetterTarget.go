package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Sources that support a dead-letter target.
// Experimental.
type SourceWithDeadLetterTarget interface {
	ISource
	// The dead-letter SQS queue or SNS topic.
	// Experimental.
	DeadLetterTarget() interface{}
	// The ARN of the source resource.
	// Experimental.
	SourceArn() *string
	// Bind the source to a pipe.
	// Experimental.
	Bind(pipe IPipe) *SourceConfig
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

// The jsii proxy struct for SourceWithDeadLetterTarget
type jsiiProxy_SourceWithDeadLetterTarget struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_SourceWithDeadLetterTarget) DeadLetterTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceWithDeadLetterTarget) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewSourceWithDeadLetterTarget_Override(s SourceWithDeadLetterTarget, sourceArn *string, deadLetterTarget interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.SourceWithDeadLetterTarget",
		[]interface{}{sourceArn, deadLetterTarget},
		s,
	)
}

// Determines if the source is an instance of SourceWithDeadLetterTarget.
// Experimental.
func SourceWithDeadLetterTarget_IsSourceWithDeadLetterTarget(source ISource) *bool {
	_init_.Initialize()

	if err := validateSourceWithDeadLetterTarget_IsSourceWithDeadLetterTargetParameters(source); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.SourceWithDeadLetterTarget",
		"isSourceWithDeadLetterTarget",
		[]interface{}{source},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) Bind(pipe IPipe) *SourceConfig {
	if err := s.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) GetDeadLetterTargetArn(deadLetterTarget interface{}) *string {
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

func (s *jsiiProxy_SourceWithDeadLetterTarget) GrantPush(grantee awsiam.IRole, deadLetterTarget interface{}) {
	if err := s.validateGrantPushParameters(grantee, deadLetterTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPush",
		[]interface{}{grantee, deadLetterTarget},
	)
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) GrantRead(grantee awsiam.IRole) {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantRead",
		[]interface{}{grantee},
	)
}

