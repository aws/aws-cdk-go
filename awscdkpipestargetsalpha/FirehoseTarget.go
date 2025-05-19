package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an Amazon Data Firehose delivery stream.
//
// Example:
//   var sourceQueue queue
//   var targetDeliveryStream deliveryStream
//
//
//   deliveryStreamTarget := targets.NewFirehoseTarget(targetDeliveryStream)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: deliveryStreamTarget,
//   })
//
// Experimental.
type FirehoseTarget interface {
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

// The jsii proxy struct for FirehoseTarget
type jsiiProxy_FirehoseTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_FirehoseTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirehoseTarget(deliveryStream awskinesisfirehose.IDeliveryStream, parameters *FirehoseTargetParameters) FirehoseTarget {
	_init_.Initialize()

	if err := validateNewFirehoseTargetParameters(deliveryStream, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.FirehoseTarget",
		[]interface{}{deliveryStream, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehoseTarget_Override(f FirehoseTarget, deliveryStream awskinesisfirehose.IDeliveryStream, parameters *FirehoseTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.FirehoseTarget",
		[]interface{}{deliveryStream, parameters},
		f,
	)
}

func (f *jsiiProxy_FirehoseTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := f.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirehoseTarget) GrantPush(grantee awsiam.IRole) {
	if err := f.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"grantPush",
		[]interface{}{grantee},
	)
}

