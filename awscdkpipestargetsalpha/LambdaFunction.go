package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an AWS Lambda Function.
//
// Example:
//   var sourceQueue queue
//   var targetFunction iFunction
//
//
//   pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
//   	InvocationType: targets.LambdaFunctionInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type LambdaFunction interface {
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

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_LambdaFunction) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunction(lambdaFunction awslambda.IFunction, parameters *LambdaFunctionParameters) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(lambdaFunction, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.LambdaFunction",
		[]interface{}{lambdaFunction, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunction_Override(l LambdaFunction, lambdaFunction awslambda.IFunction, parameters *LambdaFunctionParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.LambdaFunction",
		[]interface{}{lambdaFunction, parameters},
		l,
	)
}

func (l *jsiiProxy_LambdaFunction) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := l.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GrantPush(grantee awsiam.IRole) {
	if err := l.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"grantPush",
		[]interface{}{grantee},
	)
}

