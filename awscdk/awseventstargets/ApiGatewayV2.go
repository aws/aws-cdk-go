package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Use an API Gateway V2 HTTP APIs as a target for Amazon EventBridge rules.
//
// Example:
//   import apigwv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi HttpApi
//   var rule Rule
//
//
//   rule.AddTarget(targets.NewApiGatewayV2(httpApi))
//
type ApiGatewayV2 interface {
	awsevents.IRuleTarget
	// Returns the target IHttpApi.
	IHttpApi() awsapigatewayv2.IHttpApi
	// Returns a RuleTarget that can be used to trigger this API Gateway HTTP APIs as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-use-resource-based.html#eb-api-gateway-permissions
	//
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiGatewayV2
type jsiiProxy_ApiGatewayV2 struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_ApiGatewayV2) IHttpApi() awsapigatewayv2.IHttpApi {
	var returns awsapigatewayv2.IHttpApi
	_jsii_.Get(
		j,
		"iHttpApi",
		&returns,
	)
	return returns
}


func NewApiGatewayV2(httpApi awsapigatewayv2.IHttpApi, props *ApiGatewayProps) ApiGatewayV2 {
	_init_.Initialize()

	if err := validateNewApiGatewayV2Parameters(httpApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGatewayV2",
		[]interface{}{httpApi, props},
		&j,
	)

	return &j
}

func NewApiGatewayV2_Override(a ApiGatewayV2, httpApi awsapigatewayv2.IHttpApi, props *ApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGatewayV2",
		[]interface{}{httpApi, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayV2) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	if err := a.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

