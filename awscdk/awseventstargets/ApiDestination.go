package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Use an API Destination rule target.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_ApiKey(jsii.String("x-api-key"), awscdk.SecretValue_SecretsManager(jsii.String("ApiSecretName"))),
//   	Description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &ApiDestinationProps{
//   	Connection: Connection,
//   	Endpoint: jsii.String("https://example.com"),
//   	Description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   	Targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type ApiDestination interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger API destinations from an EventBridge event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiDestination
type jsiiProxy_ApiDestination struct {
	internal.Type__awseventsIRuleTarget
}

func NewApiDestination(apiDestination awsevents.IApiDestination, props *ApiDestinationProps) ApiDestination {
	_init_.Initialize()

	if err := validateNewApiDestinationParameters(apiDestination, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		&j,
	)

	return &j
}

func NewApiDestination_Override(a ApiDestination, apiDestination awsevents.IApiDestination, props *ApiDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		a,
	)
}

func (a *jsiiProxy_ApiDestination) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := a.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

