package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Use an AWS CloudWatch LogGroup as an event rule target.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
//   	LogEvent: targets.LogGroupTargetInput_FromObject(&LogGroupTargetInputOptions{
//   		Message: jSON.stringify(map[string]*string{
//   			"CustomField": jsii.String("CustomValue"),
//   		}),
//   	}),
//   }))
//
type CloudWatchLogGroup interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to log an event into a CloudWatch LogGroup.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CloudWatchLogGroup
type jsiiProxy_CloudWatchLogGroup struct {
	internal.Type__awseventsIRuleTarget
}

func NewCloudWatchLogGroup(logGroup awslogs.ILogGroup, props *LogGroupProps) CloudWatchLogGroup {
	_init_.Initialize()

	if err := validateNewCloudWatchLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchLogGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

func NewCloudWatchLogGroup_Override(c CloudWatchLogGroup, logGroup awslogs.ILogGroup, props *LogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		c,
	)
}

func (c *jsiiProxy_CloudWatchLogGroup) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := c.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

