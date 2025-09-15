package awsgamelift


// A reference to a GameSessionQueue resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gameSessionQueueReference := &GameSessionQueueReference{
//   	GameSessionQueueArn: jsii.String("gameSessionQueueArn"),
//   	GameSessionQueueName: jsii.String("gameSessionQueueName"),
//   }
//
type GameSessionQueueReference struct {
	// The ARN of the GameSessionQueue resource.
	GameSessionQueueArn *string `field:"required" json:"gameSessionQueueArn" yaml:"gameSessionQueueArn"`
	// The Name of the GameSessionQueue resource.
	GameSessionQueueName *string `field:"required" json:"gameSessionQueueName" yaml:"gameSessionQueueName"`
}

