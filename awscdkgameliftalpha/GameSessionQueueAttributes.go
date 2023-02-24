// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// A full specification of an gameSessionQueue that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   gameSessionQueueAttributes := &GameSessionQueueAttributes{
//   	GameSessionQueueArn: jsii.String("gameSessionQueueArn"),
//   	GameSessionQueueName: jsii.String("gameSessionQueueName"),
//   }
//
// Experimental.
type GameSessionQueueAttributes struct {
	// The ARN of the gameSessionQueue.
	//
	// At least one of `gameSessionQueueArn` and `gameSessionQueueName` must be provided.
	// Experimental.
	GameSessionQueueArn *string `field:"optional" json:"gameSessionQueueArn" yaml:"gameSessionQueueArn"`
	// The name of the gameSessionQueue.
	//
	// At least one of `gameSessionQueueName` and `gameSessionQueueArn`  must be provided.
	// Experimental.
	GameSessionQueueName *string `field:"optional" json:"gameSessionQueueName" yaml:"gameSessionQueueName"`
}

