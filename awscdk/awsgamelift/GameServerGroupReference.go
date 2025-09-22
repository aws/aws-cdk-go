package awsgamelift


// A reference to a GameServerGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gameServerGroupReference := &GameServerGroupReference{
//   	GameServerGroupArn: jsii.String("gameServerGroupArn"),
//   }
//
type GameServerGroupReference struct {
	// The GameServerGroupArn of the GameServerGroup resource.
	GameServerGroupArn *string `field:"required" json:"gameServerGroupArn" yaml:"gameServerGroupArn"`
}

