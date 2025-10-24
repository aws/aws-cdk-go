package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a new Gamelift FleetIQ Game server group.
//
// Example:
//   var launchTemplate ILaunchTemplate
//   var vpc IVpc
//
//
//   gamelift.NewGameServerGroup(this, jsii.String("GameServerGroup"), &GameServerGroupProps{
//   	GameServerGroupName: jsii.String("sample-gameservergroup-name"),
//   	InstanceDefinitions: []InstanceDefinition{
//   		&InstanceDefinition{
//   			InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C5, ec2.InstanceSize_LARGE),
//   		},
//   		&InstanceDefinition{
//   			InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   		},
//   	},
//   	LaunchTemplate: launchTemplate,
//   	Vpc: vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
// Experimental.
type GameServerGroupProps struct {
	// A developer-defined identifier for the game server group.
	//
	// The name is unique for each Region in each AWS account.
	// Experimental.
	GameServerGroupName *string `field:"required" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// The set of Amazon EC2 instance types that GameLift FleetIQ can use when balancing and automatically scaling instances in the corresponding Auto Scaling group.
	// Experimental.
	InstanceDefinitions *[]*InstanceDefinition `field:"required" json:"instanceDefinitions" yaml:"instanceDefinitions"`
	// The Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.
	//
	// After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	//
	// NOTE:
	// If you specify network interfaces in your launch template, you must explicitly set the property AssociatePublicIpAddress to `true`.
	// If no network interface is specified in the launch template, GameLift FleetIQ uses your account's default VPC.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html
	//
	// Experimental.
	LaunchTemplate awsec2.ILaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// The VPC network to place the game server group in.
	//
	// By default, all GameLift FleetIQ-supported Availability Zones are used.
	//
	// You can use this parameter to specify VPCs that you've set up.
	//
	// This property cannot be updated after the game server group is created,
	// and the corresponding Auto Scaling group will always use the property value that is set with this request,
	// even if the Auto Scaling group is updated directly.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting.
	//
	// The scaling policy uses the metric `PercentUtilizedGameServers` to maintain a buffer of idle game servers that can immediately accommodate new games and players.
	//
	// After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	// Default: no autoscaling policy settled.
	//
	// Experimental.
	AutoScalingPolicy *AutoScalingPolicy `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.
	// Default: SPOT_PREFERRED.
	//
	// Experimental.
	BalancingStrategy BalancingStrategy `field:"optional" json:"balancingStrategy" yaml:"balancingStrategy"`
	// The type of delete to perform.
	//
	// To delete a game server group, specify the DeleteOption.
	// Default: SAFE_DELETE.
	//
	// Experimental.
	DeleteOption DeleteOption `field:"optional" json:"deleteOption" yaml:"deleteOption"`
	// The maximum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
	//
	// After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	// Default: the default is 1.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and Amazon EC2 do not scale down the group below this minimum.
	//
	// In production, this value should be set to at least 1.
	//
	// After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	// Default: the default is 0.
	//
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A flag that indicates whether instances in the game server group are protected from early termination.
	//
	// Unprotected instances that have active game servers running might be terminated during a scale-down event, causing players to be dropped from the game.
	// Protected instances cannot be terminated while there are active game servers running except in the event of a forced game server group deletion.
	//
	// An exception to this is with Spot Instances, which can be terminated by AWS regardless of protection status.
	// Default: game servers running might be terminated during a scale-down event.
	//
	// Experimental.
	ProtectGameServer *bool `field:"optional" json:"protectGameServer" yaml:"protectGameServer"`
	// The IAM role that allows Amazon GameLift to access your Amazon EC2 Auto Scaling groups.
	// See: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-iam-permissions-roles.html
	//
	// Default: - a role will be created with default trust to Gamelift and Autoscaling service principal with a default policy `GameLiftGameServerGroupPolicy` attached.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Game server group subnet selection.
	// Default: all GameLift FleetIQ-supported Availability Zones are used.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

