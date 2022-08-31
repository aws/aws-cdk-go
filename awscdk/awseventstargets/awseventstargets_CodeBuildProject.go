package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
)

// Start a CodeBuild build when an Amazon EventBridge rule is triggered.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//   var project pipelineProject
//   var myTopic topic
//
//
//   // starts a CodeBuild project when a commit is pushed to the "master" branch of the repo
//   repo.onCommit(jsii.String("CommitToMaster"), &onCommitOptions{
//   	target: targets.NewCodeBuildProject(project),
//   	branches: []*string{
//   		jsii.String("master"),
//   	},
//   })
//
//   // publishes a message to an Amazon SNS topic when a comment is made on a pull request
//   rule := repo.onCommentOnPullRequest(jsii.String("CommentOnPullRequest"), &onEventOptions{
//   	target: targets.NewSnsTopic(myTopic),
//   })
//
// Experimental.
type CodeBuildProject interface {
	awsevents.IRuleTarget
	// Allows using build projects as event rule targets.
	// Experimental.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodeBuildProject
type jsiiProxy_CodeBuildProject struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewCodeBuildProject(project awscodebuild.IProject, props *CodeBuildProjectProps) CodeBuildProject {
	_init_.Initialize()

	if err := validateNewCodeBuildProjectParameters(project, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildProject{}

	_jsii_.Create(
		"monocdk.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildProject_Override(c CodeBuildProject, project awscodebuild.IProject, props *CodeBuildProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.CodeBuildProject",
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

