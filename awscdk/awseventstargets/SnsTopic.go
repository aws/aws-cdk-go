package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Use an SNS topic as a target for Amazon EventBridge rules.
//
// If the topic is imported the required permissions to publish to that topic need to be set manually.
//
// Example:
//   // publish to an SNS topic every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &OnCommitOptions{
//   	Target: targets.NewSnsTopic(topic),
//   })
//
type SnsTopic interface {
	awsevents.IRuleTarget
	Topic() awssns.ITopic
	// Returns a RuleTarget that can be used to trigger this SNS topic as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
	//
	Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SnsTopic
type jsiiProxy_SnsTopic struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SnsTopic) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


func NewSnsTopic(topic awssns.ITopic, props *SnsTopicProps) SnsTopic {
	_init_.Initialize()

	if err := validateNewSnsTopicParameters(topic, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

func NewSnsTopic_Override(s SnsTopic, topic awssns.ITopic, props *SnsTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsTopic) Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig {
	if err := s.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

