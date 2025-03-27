package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to a Kinesis stream.
//
// Example:
//   var sourceQueue queue
//   var targetStream stream
//
//
//   streamTarget := targets.NewKinesisTarget(targetStream, &KinesisTargetParameters{
//   	PartitionKey: jsii.String("pk"),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: streamTarget,
//   })
//
// Experimental.
type KinesisTarget interface {
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

// The jsii proxy struct for KinesisTarget
type jsiiProxy_KinesisTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_KinesisTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisTarget(stream awskinesis.IStream, parameters *KinesisTargetParameters) KinesisTarget {
	_init_.Initialize()

	if err := validateNewKinesisTargetParameters(stream, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.KinesisTarget",
		[]interface{}{stream, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisTarget_Override(k KinesisTarget, stream awskinesis.IStream, parameters *KinesisTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.KinesisTarget",
		[]interface{}{stream, parameters},
		k,
	)
}

func (k *jsiiProxy_KinesisTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := k.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisTarget) GrantPush(grantee awsiam.IRole) {
	if err := k.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"grantPush",
		[]interface{}{grantee},
	)
}

