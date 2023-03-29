package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An object that represents a group of filter conditions for a webhook.
//
// Every condition in a given FilterGroup must be true in order for the whole group to be true.
// You construct instances of it by calling the `#inEventOf` static factory method,
// and then calling various `andXyz` instance methods to create modified instances of it
// (this class is immutable).
//
// You pass instances of this class to the `webhookFilters` property when constructing a source.
//
// Example:
//   gitHubSource := codebuild.Source_GitHub(&GitHubSourceProps{
//   	Owner: jsii.String("awslabs"),
//   	Repo: jsii.String("aws-cdk"),
//   	Webhook: jsii.Boolean(true),
//   	 // optional, default: true if `webhookFilters` were provided, false otherwise
//   	WebhookTriggersBatchBuild: jsii.Boolean(true),
//   	 // optional, default is false
//   	WebhookFilters: []filterGroup{
//   		codebuild.*filterGroup_InEventOf(codebuild.EventAction_PUSH).AndBranchIs(jsii.String("main")).AndCommitMessageIs(jsii.String("the commit message")),
//   	},
//   })
//
type FilterGroup interface {
	// Create a new FilterGroup with an added condition: the account ID of the actor initiating the event must match the given pattern.
	AndActorAccountIs(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the account ID of the actor initiating the event must not match the given pattern.
	AndActorAccountIsNot(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must target the given base branch.
	//
	// Note that you cannot use this method if this Group contains the `PUSH` event action.
	AndBaseBranchIs(branchName *string) FilterGroup
	// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must not target the given base branch.
	//
	// Note that you cannot use this method if this Group contains the `PUSH` event action.
	AndBaseBranchIsNot(branchName *string) FilterGroup
	// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must target the given Git reference.
	//
	// Note that you cannot use this method if this Group contains the `PUSH` event action.
	AndBaseRefIs(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must not target the given Git reference.
	//
	// Note that you cannot use this method if this Group contains the `PUSH` event action.
	AndBaseRefIsNot(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must affect the given branch.
	AndBranchIs(branchName *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must not affect the given branch.
	AndBranchIsNot(branchName *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must affect a head commit with the given message.
	AndCommitMessageIs(commitMessage *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must not affect a head commit with the given message.
	AndCommitMessageIsNot(commitMessage *string) FilterGroup
	// Create a new FilterGroup with an added condition: the push that is the source of the event must affect a file that matches the given pattern.
	//
	// Note that you can only use this method if this Group contains only the `PUSH` event action,
	// and only for GitHub, Bitbucket and GitHubEnterprise sources.
	AndFilePathIs(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the push that is the source of the event must not affect a file that matches the given pattern.
	//
	// Note that you can only use this method if this Group contains only the `PUSH` event action,
	// and only for GitHub, Bitbucket and GitHubEnterprise sources.
	AndFilePathIsNot(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must affect a Git reference (ie., a branch or a tag) that matches the given pattern.
	AndHeadRefIs(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must not affect a Git reference (ie., a branch or a tag) that matches the given pattern.
	AndHeadRefIsNot(pattern *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must affect the given tag.
	AndTagIs(tagName *string) FilterGroup
	// Create a new FilterGroup with an added condition: the event must not affect the given tag.
	AndTagIsNot(tagName *string) FilterGroup
}

// The jsii proxy struct for FilterGroup
type jsiiProxy_FilterGroup struct {
	_ byte // padding
}

// Creates a new event FilterGroup that triggers on any of the provided actions.
func FilterGroup_InEventOf(actions ...EventAction) FilterGroup {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns FilterGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.FilterGroup",
		"inEventOf",
		args,
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndActorAccountIs(pattern *string) FilterGroup {
	if err := f.validateAndActorAccountIsParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andActorAccountIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndActorAccountIsNot(pattern *string) FilterGroup {
	if err := f.validateAndActorAccountIsNotParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andActorAccountIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBaseBranchIs(branchName *string) FilterGroup {
	if err := f.validateAndBaseBranchIsParameters(branchName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseBranchIs",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBaseBranchIsNot(branchName *string) FilterGroup {
	if err := f.validateAndBaseBranchIsNotParameters(branchName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseBranchIsNot",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBaseRefIs(pattern *string) FilterGroup {
	if err := f.validateAndBaseRefIsParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseRefIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBaseRefIsNot(pattern *string) FilterGroup {
	if err := f.validateAndBaseRefIsNotParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseRefIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBranchIs(branchName *string) FilterGroup {
	if err := f.validateAndBranchIsParameters(branchName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBranchIs",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndBranchIsNot(branchName *string) FilterGroup {
	if err := f.validateAndBranchIsNotParameters(branchName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBranchIsNot",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndCommitMessageIs(commitMessage *string) FilterGroup {
	if err := f.validateAndCommitMessageIsParameters(commitMessage); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andCommitMessageIs",
		[]interface{}{commitMessage},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndCommitMessageIsNot(commitMessage *string) FilterGroup {
	if err := f.validateAndCommitMessageIsNotParameters(commitMessage); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andCommitMessageIsNot",
		[]interface{}{commitMessage},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndFilePathIs(pattern *string) FilterGroup {
	if err := f.validateAndFilePathIsParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andFilePathIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndFilePathIsNot(pattern *string) FilterGroup {
	if err := f.validateAndFilePathIsNotParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andFilePathIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndHeadRefIs(pattern *string) FilterGroup {
	if err := f.validateAndHeadRefIsParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andHeadRefIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndHeadRefIsNot(pattern *string) FilterGroup {
	if err := f.validateAndHeadRefIsNotParameters(pattern); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andHeadRefIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndTagIs(tagName *string) FilterGroup {
	if err := f.validateAndTagIsParameters(tagName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andTagIs",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FilterGroup) AndTagIsNot(tagName *string) FilterGroup {
	if err := f.validateAndTagIsNotParameters(tagName); err != nil {
		panic(err)
	}
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andTagIsNot",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

