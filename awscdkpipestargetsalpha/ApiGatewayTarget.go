package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an EventBridge API destination.
//
// Example:
//   var sourceQueue Queue
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = e => {}")),
//   })
//
//   restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &LambdaRestApiProps{
//   	Handler: fn,
//   })
//   apiTarget := targets.NewApiGatewayTarget(restApi)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: apiTarget,
//   })
//
// Experimental.
type ApiGatewayTarget interface {
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

// The jsii proxy struct for ApiGatewayTarget
type jsiiProxy_ApiGatewayTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_ApiGatewayTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayTarget(restApi awsapigateway.IRestApi, parameters *ApiGatewayTargetParameters) ApiGatewayTarget {
	_init_.Initialize()

	if err := validateNewApiGatewayTargetParameters(restApi, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.ApiGatewayTarget",
		[]interface{}{restApi, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayTarget_Override(a ApiGatewayTarget, restApi awsapigateway.IRestApi, parameters *ApiGatewayTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.ApiGatewayTarget",
		[]interface{}{restApi, parameters},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
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

func (a *jsiiProxy_ApiGatewayTarget) GrantPush(grantee awsiam.IRole) {
	if err := a.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantPush",
		[]interface{}{grantee},
	)
}

