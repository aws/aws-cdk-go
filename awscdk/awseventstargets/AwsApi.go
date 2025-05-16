package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Use an AWS Lambda function that makes API calls as an event rule target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var policyStatement policyStatement
//
//   awsApi := awscdk.Aws_events_targets.NewAwsApi(&AwsApiProps{
//   	Action: jsii.String("action"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
//   	CatchErrorPattern: jsii.String("catchErrorPattern"),
//   	Parameters: parameters,
//   	PolicyStatement: policyStatement,
//   })
//
type AwsApi interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this AwsApi as a result from an EventBridge event.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
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

func (a *jsiiProxy_AwsApi) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
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

