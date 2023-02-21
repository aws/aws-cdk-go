package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// The abstract view of an AWS CodePipeline as required and used by Actions.
//
// It extends `events.IRuleTarget`,
// so this interface can be used as a Target for CloudWatch Events.
type IPipeline interface {
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	// Defines a CodeStar notification rule triggered when the pipeline events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnExecutionStateChange`, `notifyOnAnyStageStateChange`,
	// `notifyOnAnyActionStateChange` and `notifyOnAnyManualApprovalStateChange`
	// to define rules for these specific event emitted.
	//
	// Returns: CodeStar notification rule associated with this build project.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline
	//
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an event rule triggered by this CodePipeline.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Define an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of the Pipeline.
	PipelineArn() *string
	// The name of the Pipeline.
	PipelineName() *string
}

// The jsii proxy for IPipeline
type jsiiProxy_IPipeline struct {
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPipeline) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnAnyActionStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyActionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnAnyManualApprovalStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyManualApprovalStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnAnyStageStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnAnyStageStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnExecutionStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnExecutionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnStateChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IPipeline) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	if err := i.validateBindAsNotificationRuleSourceParameters(scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

