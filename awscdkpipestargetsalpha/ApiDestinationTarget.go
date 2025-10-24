package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an EventBridge API destination.
//
// Example:
//   var sourceQueue Queue
//   var dest ApiDestination
//
//
//   apiTarget := targets.NewApiDestinationTarget(dest)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: apiTarget,
//   })
//
// Experimental.
type ApiDestinationTarget interface {
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

// The jsii proxy struct for ApiDestinationTarget
type jsiiProxy_ApiDestinationTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_ApiDestinationTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiDestinationTarget(destination awsevents.IApiDestination, parameters *ApiDestinationTargetParameters) ApiDestinationTarget {
	_init_.Initialize()

	if err := validateNewApiDestinationTargetParameters(destination, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiDestinationTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.ApiDestinationTarget",
		[]interface{}{destination, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewApiDestinationTarget_Override(a ApiDestinationTarget, destination awsevents.IApiDestination, parameters *ApiDestinationTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.ApiDestinationTarget",
		[]interface{}{destination, parameters},
		a,
	)
}

func (a *jsiiProxy_ApiDestinationTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := a.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDestinationTarget) GrantPush(grantee awsiam.IRole) {
	if err := a.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantPush",
		[]interface{}{grantee},
	)
}

