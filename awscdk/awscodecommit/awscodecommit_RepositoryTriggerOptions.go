package awscodecommit


// Creates for a repository trigger to an SNS topic or Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryTriggerOptions := &repositoryTriggerOptions{
//   	branches: []*string{
//   		jsii.String("branches"),
//   	},
//   	customData: jsii.String("customData"),
//   	events: []repositoryEventTrigger{
//   		awscdk.Aws_codecommit.*repositoryEventTrigger_ALL,
//   	},
//   	name: jsii.String("name"),
//   }
//
type RepositoryTriggerOptions struct {
	// The names of the branches in the AWS CodeCommit repository that contain events that you want to include in the trigger.
	//
	// If you don't specify at
	// least one branch, the trigger applies to all branches.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// When an event is triggered, additional information that AWS CodeCommit includes when it sends information to the target.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// The repository events for which AWS CodeCommit sends information to the target, which you specified in the DestinationArn property.If you don't specify events, the trigger runs for all repository events.
	Events *[]RepositoryEventTrigger `field:"optional" json:"events" yaml:"events"`
	// A name for the trigger.Triggers on a repository must have unique names.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

