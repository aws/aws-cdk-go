package awscodecommit


// Information about a trigger for a repository.
//
// > If you want to receive notifications about repository events, consider using notifications instead of triggers. For more information, see [Configuring notifications for repository events](https://docs.aws.amazon.com/codecommit/latest/userguide/how-to-repository-email.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryTriggerProperty := &repositoryTriggerProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	branches: []*string{
//   		jsii.String("branches"),
//   	},
//   	customData: jsii.String("customData"),
//   }
//
type CfnRepository_RepositoryTriggerProperty struct {
	// The ARN of the resource that is the target for a trigger (for example, the ARN of a topic in Amazon SNS).
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The repository events that cause the trigger to run actions in another service, such as sending a notification through Amazon SNS.
	//
	// > The valid value "all" cannot be used with any other values.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// The name of the trigger.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The branches to be included in the trigger configuration.
	//
	// If you specify an empty array, the trigger applies to all branches.
	//
	// > Although no content is required in the array, you must include the array itself.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Any custom data associated with the trigger to be included in the information sent to the target of the trigger.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
}

