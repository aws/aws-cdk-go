package awscodecommit

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type IRepository interface {
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	// Grant the given principal identity permissions to perform the actions on this repository.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to pull this repository.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push this repository.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to read this repository.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Defines a CodeStar Notification rule triggered when the project events specified by you are emitted. Similar to `onEvent` API.
	//
	// You can also use the methods to define rules for the specific event emitted.
	// eg: `notifyOnPullRequstCreated`.
	//
	// Returns: CodeStar Notifications rule associated with this repository.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval rule is overridden.
	NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when an approval status is changed.
	NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a new branch or tag is created.
	NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a branch or tag is deleted.
	NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a comment is made on a pull request.
	NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is created.
	NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar Notification rule which triggers when a pull request is merged.
	NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CloudWatch event rule which triggers when a comment is made on a commit.
	OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a comment is made on a pull request.
	OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a commit is pushed to a branch.
	OnCommit(id *string, options *OnCommitOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a pull request state is changed.
	OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is created (i.e. a new branch/tag is created) to the repository.
	OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is delete (i.e. a branch/tag is deleted) from the repository.
	OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a reference is updated (i.e. a commit is pushed to an existing or new branch) from the repository.
	OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers when a "CodeCommit Repository State Change" event occurs.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of this Repository.
	RepositoryArn() *string
	// The HTTPS (GRC) clone URL.
	//
	// HTTPS (GRC) is the protocol to use with git-remote-codecommit (GRC).
	//
	// It is the recommended method for supporting connections made with federated
	// access, identity providers, and temporary credentials.
	// See: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-git-remote-codecommit.html
	//
	RepositoryCloneUrlGrc() *string
	// The HTTP clone URL.
	RepositoryCloneUrlHttp() *string
	// The SSH clone URL.
	RepositoryCloneUrlSsh() *string
	// The human-visible name of this Repository.
	RepositoryName() *string
}

// The jsii proxy for IRepository
type jsiiProxy_IRepository struct {
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRepository) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPullParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPullPushParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) awscodestarnotifications.INotificationRule {
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

func (i *jsiiProxy_IRepository) NotifyOnApprovalRuleOverridden(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnApprovalRuleOverriddenParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnApprovalRuleOverridden",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnApprovalStatusChanged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnApprovalStatusChangedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnApprovalStatusChanged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnBranchOrTagCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnBranchOrTagCreatedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBranchOrTagCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnBranchOrTagDeleted(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnBranchOrTagDeletedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBranchOrTagDeleted",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestComment(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnPullRequestCommentParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestComment",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestCreated(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnPullRequestCreatedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestCreated",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) NotifyOnPullRequestMerged(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := i.validateNotifyOnPullRequestMergedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnPullRequestMerged",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommentOnCommit(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnCommentOnCommitParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommentOnCommit",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommentOnPullRequest(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnCommentOnPullRequestParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommentOnPullRequest",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCommit(id *string, options *OnCommitOptions) awsevents.Rule {
	if err := i.validateOnCommitParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCommit",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
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

func (i *jsiiProxy_IRepository) OnPullRequestStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnPullRequestStateChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onPullRequestStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceCreated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnReferenceCreatedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceCreated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceDeleted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnReferenceDeletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceDeleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnReferenceUpdated(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnReferenceUpdatedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReferenceUpdated",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
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

func (i *jsiiProxy_IRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRepository) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
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

func (j *jsiiProxy_IRepository) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryCloneUrlGrc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlGrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryCloneUrlHttp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryCloneUrlSsh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCloneUrlSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

