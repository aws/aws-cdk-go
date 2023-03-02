package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awschatbot/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a Slack channel configuration.
// Experimental.
type ISlackChannelConfiguration interface {
	awsiam.IGrantable
	awscodestarnotifications.INotificationRuleTarget
	awscdk.IResource
	// Adds a statement to the IAM role.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Return the given named metric for this SlackChannelConfiguration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The permission role of Slack channel configuration.
	// Experimental.
	Role() awsiam.IRole
	// The ARN of the Slack channel configuration In the form of arn:aws:chatbot:{region}:{account}:chat-configuration/slack-channel/{slackChannelName}.
	// Experimental.
	SlackChannelConfigurationArn() *string
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName() *string
}

// The jsii proxy for ISlackChannelConfiguration
type jsiiProxy_ISlackChannelConfiguration struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscodestarnotificationsINotificationRuleTarget
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_ISlackChannelConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISlackChannelConfiguration) BindAsNotificationRuleTarget(scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	if err := i.validateBindAsNotificationRuleTargetParameters(scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) SlackChannelConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) SlackChannelConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

