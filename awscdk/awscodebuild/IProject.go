package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type IProject interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	AddToRolePolicy(policyStatement awsiam.PolicyStatement)
	// Enable batch builds.
	//
	// Returns an object contining the batch service role if batch builds
	// could be enabled.
	EnableBatchBuilds() *BatchBuildConfig
	// Returns: a CloudWatch metric associated with this build project.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds triggered.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	// Default: sum over 5 minutes.
	//
	MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the duration of all builds over time.
	//
	// Units: Seconds
	//
	// Valid CloudWatch statistics: Average (recommended), Maximum, Minimum.
	// Default: average over 5 minutes.
	//
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds that failed because of client error or because of a timeout.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	// Default: sum over 5 minutes.
	//
	MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of successful builds.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	// Default: sum over 5 minutes.
	//
	MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CodeStar Notification rule triggered when the project events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnBuildSucceeded` and
	// `notifyOnBuildFailed` to define rules for these specific event emitted.
	//
	// Returns: CodeStar Notifications rule associated with this build project.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build fails.
	NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build completes successfully.
	NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines an event rule which triggers when a build fails.
	OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build starts.
	OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build completes successfully.
	OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when something happens with this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule that triggers upon phase change of this build project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when the build project state changes.
	//
	// You can filter specific build status events using an event
	// pattern filter on the `build-status` detail field:
	//
	//    const rule = project.onStateChange('OnBuildStarted', { target });
	//    rule.addEventPattern({
	//      detail: {
	//        'build-status': [
	//          "IN_PROGRESS",
	//          "SUCCEEDED",
	//          "FAILED",
	//          "STOPPED"
	//        ]
	//      }
	//    });
	//
	// You can also use the methods `onBuildFailed` and `onBuildSucceeded` to define rules for
	// these specific state changes.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of this Project.
	ProjectArn() *string
	// The human-visible name of this Project.
	ProjectName() *string
	// The IAM service Role of this Project.
	//
	// Undefined for imported Projects.
	Role() awsiam.IRole
}

// The jsii proxy for IProject
type jsiiProxy_IProject struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProject) AddToRolePolicy(policyStatement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(policyStatement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{policyStatement},
	)
}

func (i *jsiiProxy_IProject) EnableBatchBuilds() *BatchBuildConfig {
	var returns *BatchBuildConfig

	_jsii_.Invoke(
		i,
		"enableBatchBuilds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IProject) MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFailedBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailedBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSucceededBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSucceededBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule {
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

func (i *jsiiProxy_IProject) NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnBuildFailedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBuildFailed",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnBuildSucceededParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBuildSucceeded",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnBuildFailedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnBuildStartedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnBuildSucceededParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
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

func (i *jsiiProxy_IProject) OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnPhaseChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onPhaseChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
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

func (i *jsiiProxy_IProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IProject) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
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

func (j *jsiiProxy_IProject) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

