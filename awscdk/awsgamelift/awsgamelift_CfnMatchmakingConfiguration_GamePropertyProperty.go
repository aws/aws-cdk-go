package awsgamelift


// Set of key-value pairs that contain information about a game session.
//
// When included in a game session request, these properties communicate details to be used when setting up the new game session. For example, a game property might specify a game mode, level, or map. Game properties are passed to the game server process when initiating a new game session. For more information, see the [GameLift Developer Guide](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-client-api.html#gamelift-sdk-client-api-create) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gamePropertyProperty := &gamePropertyProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnMatchmakingConfiguration_GamePropertyProperty struct {
	// The game property identifier.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The game property value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

