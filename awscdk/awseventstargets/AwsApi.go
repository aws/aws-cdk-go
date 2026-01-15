package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Use an AWS Lambda function that makes API calls as an event rule target.
//
// Example:
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewAwsApi(&AwsApiProps{
//   	Service: jsii.String("ECS"),
//   	Action: jsii.String("updateService"),
//   	Parameters: map[string]interface{}{
//   		"service": jsii.String("my-service"),
//   		"forceNewDeployment": jsii.Boolean(true),
//   	},
//   }))
//
type AwsApi interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this AwsApi as a result from an EventBridge event.
	Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for AwsApi
type jsiiProxy_AwsApi struct {
	internal.Type__awseventsIRuleTarget
}

func NewAwsApi(props *AwsApiProps) AwsApi {
	_init_.Initialize()

	if err := validateNewAwsApiParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AwsApi",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAwsApi_Override(a AwsApi, props *AwsApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AwsApi",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsApi) Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig {
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

