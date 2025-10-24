package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a GameServerGroup content defined outside of this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   gameServerGroupAttributes := &GameServerGroupAttributes{
//   	AutoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   	// the properties below are optional
//   	GameServerGroupArn: jsii.String("gameServerGroupArn"),
//   	GameServerGroupName: jsii.String("gameServerGroupName"),
//   	Role: role,
//   }
//
// Experimental.
type GameServerGroupAttributes struct {
	// The ARN of the generated AutoScaling group.
	// Default: the imported game server group does not have autoscaling group information.
	//
	// Experimental.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The ARN of the game server group.
	//
	// At least one of `gameServerGroupArn` and `gameServerGroupName` must be provided.
	// Default: derived from `gameServerGroupName`.
	//
	// Experimental.
	GameServerGroupArn *string `field:"optional" json:"gameServerGroupArn" yaml:"gameServerGroupArn"`
	// The name of the game server group.
	//
	// At least one of `gameServerGroupArn` and `gameServerGroupName` must be provided.
	// Default: derived from `gameServerGroupArn`.
	//
	// Experimental.
	GameServerGroupName *string `field:"optional" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// The IAM role that allows Amazon GameLift to access your Amazon EC2 Auto Scaling groups.
	// Default: the imported game server group cannot be granted access to other resources as an `iam.IGrantable`.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

