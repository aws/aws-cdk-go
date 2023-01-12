// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// A set of custom properties for a game session, formatted as key-value pairs.
//
// These properties are passed to a game server process with a request to start a new game session.
//
// This parameter is not used for Standalone FlexMatch mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   gameProperty := &gameProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession
//
// Experimental.
type GameProperty struct {
	// The game property identifier.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The game property value.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

