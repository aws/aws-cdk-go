package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Use an API Gateway REST APIs as a target for Amazon EventBridge rules.
//
// Example:
//   import api "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   })
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
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.AddTarget(
//   targets.NewApiGateway(restApi, &ApiGatewayProps{
//   	Path: jsii.String("/*/test"),
//   	Method: jsii.String("GET"),
//   	Stage: jsii.String("prod"),
//   	PathParameterValues: []*string{
//   		jsii.String("path-value"),
//   	},
//   	HeaderParameters: map[string]*string{
//   		"Header1": jsii.String("header1"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"QueryParam1": jsii.String("query-param-1"),
//   	},
//   	DeadLetterQueue: dlq,
//   }))
//
type ApiGateway interface {
	awsevents.IRuleTarget
	RestApi() awsapigateway.RestApi
	// Returns a RuleTarget that can be used to trigger this API Gateway REST APIs as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
	//
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiGateway
type jsiiProxy_ApiGateway struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_ApiGateway) RestApi() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"restApi",
		&returns,
	)
	return returns
}


func NewApiGateway(restApi awsapigateway.RestApi, props *ApiGatewayProps) ApiGateway {
	_init_.Initialize()

	if err := validateNewApiGatewayParameters(restApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

func NewApiGateway_Override(a ApiGateway, restApi awsapigateway.RestApi, props *ApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		a,
	)
}

func (a *jsiiProxy_ApiGateway) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := a.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

