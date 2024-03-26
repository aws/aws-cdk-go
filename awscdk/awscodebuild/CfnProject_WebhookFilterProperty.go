package awscodebuild


// `WebhookFilter` is a structure of the `FilterGroups` property on the [AWS CodeBuild Project ProjectTriggers](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html) property type that specifies which webhooks trigger an AWS CodeBuild build.
//
// > The Webhook feature isn't available in AWS CloudFormation for GitHub Enterprise projects. Use the AWS CLI or AWS CodeBuild console to create the webhook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookFilterProperty := &WebhookFilterProperty{
//   	Pattern: jsii.String("pattern"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ExcludeMatchedPattern: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-webhookfilter.html
//
type CfnProject_WebhookFilterProperty struct {
	// For a `WebHookFilter` that uses `EVENT` type, a comma-separated string that specifies one or more events.
	//
	// For example, the webhook filter `PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED` allows all push, pull request created, and pull request updated events to trigger a build.
	//
	// For a `WebHookFilter` that uses any of the other filter types, a regular expression pattern. For example, a `WebHookFilter` that uses `HEAD_REF` for its `type` and the pattern `^refs/heads/` triggers a build when the head reference is a branch with a reference name `refs/heads/branch-name` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-webhookfilter.html#cfn-codebuild-project-webhookfilter-pattern
	//
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The type of webhook filter.
	//
	// There are eight webhook filter types: `EVENT` , `ACTOR_ACCOUNT_ID` , `HEAD_REF` , `BASE_REF` , `FILE_PATH` , `COMMIT_MESSAGE` , `TAG_NAME` , and `RELEASE_NAME` .
	//
	// - EVENT
	//
	// - A webhook event triggers a build when the provided `pattern` matches one of eight event types: `PUSH` , `PULL_REQUEST_CREATED` , `PULL_REQUEST_UPDATED` , `PULL_REQUEST_CLOSED` , `PULL_REQUEST_REOPENED` , `PULL_REQUEST_MERGED` , `RELEASED` , and `PRERELEASED` . The `EVENT` patterns are specified as a comma-separated string. For example, `PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED` filters all push, pull request created, and pull request updated events.
	//
	// > The `PULL_REQUEST_REOPENED` works with GitHub and GitHub Enterprise only. The `RELEASED` and `PRERELEASED` work with GitHub only.
	// - ACTOR_ACCOUNT_ID
	//
	// - A webhook event triggers a build when a GitHub, GitHub Enterprise, or Bitbucket account ID matches the regular expression `pattern` .
	// - HEAD_REF
	//
	// - A webhook event triggers a build when the head reference matches the regular expression `pattern` . For example, `refs/heads/branch-name` and `refs/tags/tag-name` .
	//
	// > Works with GitHub and GitHub Enterprise push, GitHub and GitHub Enterprise pull request, Bitbucket push, and Bitbucket pull request events.
	// - BASE_REF
	//
	// - A webhook event triggers a build when the base reference matches the regular expression `pattern` . For example, `refs/heads/branch-name` .
	//
	// > Works with pull request events only.
	// - FILE_PATH
	//
	// - A webhook triggers a build when the path of a changed file matches the regular expression `pattern` .
	//
	// > Works with GitHub and Bitbucket events push and pull requests events. Also works with GitHub Enterprise push events, but does not work with GitHub Enterprise pull request events.
	// - COMMIT_MESSAGE
	//
	// - A webhook triggers a build when the head commit message matches the regular expression `pattern` .
	//
	// > Works with GitHub and Bitbucket events push and pull requests events. Also works with GitHub Enterprise push events, but does not work with GitHub Enterprise pull request events.
	// - TAG_NAME
	//
	// - A webhook triggers a build when the tag name of the release matches the regular expression `pattern` .
	//
	// > Works with `RELEASED` and `PRERELEASED` events only.
	// - RELEASE_NAME
	//
	// - A webhook triggers a build when the release name matches the regular expression `pattern` .
	//
	// > Works with `RELEASED` and `PRERELEASED` events only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-webhookfilter.html#cfn-codebuild-project-webhookfilter-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Used to indicate that the `pattern` determines which webhook events do not trigger a build.
	//
	// If true, then a webhook event that does not match the `pattern` triggers a build. If false, then a webhook event that matches the `pattern` triggers a build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-webhookfilter.html#cfn-codebuild-project-webhookfilter-excludematchedpattern
	//
	ExcludeMatchedPattern interface{} `field:"optional" json:"excludeMatchedPattern" yaml:"excludeMatchedPattern"`
}

