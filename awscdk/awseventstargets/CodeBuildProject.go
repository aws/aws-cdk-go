package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Start a CodeBuild build when an Amazon EventBridge rule is triggered.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//   var project pipelineProject
//   var myTopic topic
//
//
//   // starts a CodeBuild project when a commit is pushed to the "main" branch of the repo
//   repo.onCommit(jsii.String("CommitToMain"), &OnCommitOptions{
//   	Target: targets.NewCodeBuildProject(project),
//   	Branches: []*string{
//   		jsii.String("main"),
//   	},
//   })
//
//   // publishes a message to an Amazon SNS topic when a comment is made on a pull request
//   rule := repo.onCommentOnPullRequest(jsii.String("CommentOnPullRequest"), &OnEventOptions{
//   	Target: targets.NewSnsTopic(myTopic),
//   })
//
type CodeBuildProject interface {
	awsevents.IRuleTarget
	// Allows using build projects as event rule targets.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodeBuildProject
type jsiiProxy_CodeBuildProject struct {
	internal.Type__awseventsIRuleTarget
}

func NewCodeBuildProject(project awscodebuild.IProject, props *CodeBuildProjectProps) CodeBuildProject {
	_init_.Initialize()

	if err := validateNewCodeBuildProjectParameters(project, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

func NewCodeBuildProject_Override(c CodeBuildProject, project awscodebuild.IProject, props *CodeBuildProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildProject) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
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

