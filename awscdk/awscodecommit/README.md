# AWS CodeCommit Construct Library

AWS CodeCommit is a version control service that enables you to privately store and manage Git repositories in the AWS cloud.

For further information on CodeCommit,
see the [AWS CodeCommit documentation](https://docs.aws.amazon.com/codecommit).

To add a CodeCommit Repository to your stack:

```go
repo := codecommit.NewRepository(this, jsii.String("Repository"), &RepositoryProps{
	RepositoryName: jsii.String("MyRepositoryName"),
	Description: jsii.String("Some description."),
})
```

Use the `repositoryCloneUrlHttp`, `repositoryCloneUrlSsh` or `repositoryCloneUrlGrc`
property to clone your repository.

To add an Amazon SNS trigger to your repository:

```go
var repo repository


// trigger is established for all repository actions on all branches by default.
repo.Notify(jsii.String("arn:aws:sns:*:123456789012:my_topic"))
```

## Add initial commit

It is possible to initialize the Repository via the `Code` class.
It provides methods for loading code from a directory, `.zip` file and from a pre-created CDK Asset.

Example:

```go
repo := codecommit.NewRepository(this, jsii.String("Repository"), &RepositoryProps{
	RepositoryName: jsii.String("MyRepositoryName"),
	Code: codecommit.Code_FromDirectory(path.join(__dirname, jsii.String("directory/")), jsii.String("develop")),
})
```

## Use a customer managed key

CodeCommit repositories are automatically encrypted with an AWS managed key. To use
a customer managed key, specify the `kmsKey` property.

For more information, see
[AWS Key Management Service and encryption for AWS CodeCommit repositories](https://docs.aws.amazon.com/cdk/latest/guide/reference.html#versioning).

```go
import kms "github.com/aws/aws-cdk-go/awscdk"

var kmsKey iKey


repo := codecommit.NewRepository(this, jsii.String("Repository"), &RepositoryProps{
	RepositoryName: jsii.String("MyRepositoryName"),
	KmsKey: KmsKey,
})
```

## Events

CodeCommit repositories emit Amazon CloudWatch events for certain activities.
Use the `repo.onXxx` methods to define rules that trigger on these events
and invoke targets as a result:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var repo repository
var project pipelineProject
var myTopic topic


// starts a CodeBuild project when a commit is pushed to the "main" branch of the repo
repo.onCommit(jsii.String("CommitToMain"), &OnCommitOptions{
	Target: targets.NewCodeBuildProject(project),
	Branches: []*string{
		jsii.String("main"),
	},
})

// publishes a message to an Amazon SNS topic when a comment is made on a pull request
rule := repo.onCommentOnPullRequest(jsii.String("CommentOnPullRequest"), &OnEventOptions{
	Target: targets.NewSnsTopic(myTopic),
})
```

## CodeStar Notifications

To define CodeStar Notification rules for Repositories, use one of the `notifyOnXxx()` methods.
They are very similar to `onXxx()` methods for CloudWatch events:

```go
import chatbot "github.com/aws/aws-cdk-go/awscdk"

var repository repository

target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
})
rule := repository.notifyOnPullRequestCreated(jsii.String("NotifyOnPullRequestCreated"), target)
```
