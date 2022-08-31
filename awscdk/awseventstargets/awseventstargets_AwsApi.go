package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
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
//   awsApi := awscdk.Aws_events_targets.NewAwsApi(&awsApiProps{
//   	action: jsii.String("action"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	catchErrorPattern: jsii.String("catchErrorPattern"),
//   	parameters: parameters,
//   	policyStatement: policyStatement,
//   })
//
// Experimental.
type AwsApi interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this AwsApi as a result from an EventBridge event.
	// Experimental.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for AwsApi
type jsiiProxy_AwsApi struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewAwsApi(props *AwsApiProps) AwsApi {
	_init_.Initialize()

	j := jsiiProxy_AwsApi{}

	_jsii_.Create(
		"monocdk.aws_events_targets.AwsApi",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApi_Override(a AwsApi, props *AwsApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.AwsApi",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsApi) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

