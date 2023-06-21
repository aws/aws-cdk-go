package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an SQS Queue as a target for Amazon EventBridge rules.
//
// Example:
//   // publish to an SQS queue every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &OnCommitOptions{
//   	Target: targets.NewSqsQueue(queue),
//   })
//
type SqsQueue interface {
	awsevents.IRuleTarget
	Queue() awssqs.IQueue
	// Returns a RuleTarget that can be used to trigger this SQS queue as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
	//
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SqsQueue
type jsiiProxy_SqsQueue struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SqsQueue) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}


func NewSqsQueue(queue awssqs.IQueue, props *SqsQueueProps) SqsQueue {
	_init_.Initialize()

	if err := validateNewSqsQueueParameters(queue, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsQueue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsQueue_Override(s SqsQueue, queue awssqs.IQueue, props *SqsQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsQueue) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := s.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

