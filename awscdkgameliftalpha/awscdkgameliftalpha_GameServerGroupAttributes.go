// The CDK Construct Library for AWS::GameLift
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
//   var role role
//
//   gameServerGroupAttributes := &gameServerGroupAttributes{
//   	autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   	// the properties below are optional
//   	gameServerGroupArn: jsii.String("gameServerGroupArn"),
//   	gameServerGroupName: jsii.String("gameServerGroupName"),
//   	role: role,
//   }
//
// Experimental.
type GameServerGroupAttributes struct {
	// The ARN of the generated AutoScaling group.
	// Experimental.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The ARN of the game server group.
	//
	// At least one of `gameServerGroupArn` and `gameServerGroupName` must be provided.
	// Experimental.
	GameServerGroupArn *string `field:"optional" json:"gameServerGroupArn" yaml:"gameServerGroupArn"`
	// The name of the game server group.
	//
	// At least one of `gameServerGroupArn` and `gameServerGroupName` must be provided.
	// Experimental.
	GameServerGroupName *string `field:"optional" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// The IAM role that allows Amazon GameLift to access your Amazon EC2 Auto Scaling groups.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

