package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the CodeBuild Event Target.
//
// Example:
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repo := codecommit.NewRepository(this, jsii.String("MyRepo"), &repositoryProps{
//   	repositoryName: jsii.String("aws-cdk-codebuild-events"),
//   })
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	source: codebuild.source.codeCommit(&codeCommitSourceProps{
//   		repository: repo,
//   	}),
//   })
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   // trigger a build when a commit is pushed to the repo
//   onCommitRule := repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
//   	target: targets.NewCodeBuildProject(project, &codeBuildProjectProps{
//   		deadLetterQueue: deadLetterQueue,
//   	}),
//   	branches: []*string{
//   		jsii.String("master"),
//   	},
//   })
//
type CodeBuildProjectProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to CodeBuild.
	//
	// This will be the payload for the StartBuild API.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The role to assume before invoking the target (i.e., the codebuild) when the given rule is triggered.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
}

