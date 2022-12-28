package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGameServerGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGameServerGroupProps := &cfnGameServerGroupProps{
//   	gameServerGroupName: jsii.String("gameServerGroupName"),
//   	instanceDefinitions: []interface{}{
//   		&instanceDefinitionProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			weightedCapacity: jsii.String("weightedCapacity"),
//   		},
//   	},
//   	launchTemplate: &launchTemplateProperty{
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   		version: jsii.String("version"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	autoScalingPolicy: &autoScalingPolicyProperty{
//   		targetTrackingConfiguration: &targetTrackingConfigurationProperty{
//   			targetValue: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		estimatedInstanceWarmup: jsii.Number(123),
//   	},
//   	balancingStrategy: jsii.String("balancingStrategy"),
//   	deleteOption: jsii.String("deleteOption"),
//   	gameServerProtectionPolicy: jsii.String("gameServerProtectionPolicy"),
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSubnets: []*string{
//   		jsii.String("vpcSubnets"),
//   	},
//   }
//
type CfnGameServerGroupProps struct {
	// A developer-defined identifier for the game server group.
	//
	// The name is unique for each Region in each AWS account.
	GameServerGroupName *string `field:"required" json:"gameServerGroupName" yaml:"gameServerGroupName"`
	// The set of Amazon EC2 instance types that GameLift FleetIQ can use when balancing and automatically scaling instances in the corresponding Auto Scaling group.
	InstanceDefinitions interface{} `field:"required" json:"instanceDefinitions" yaml:"instanceDefinitions"`
	// The Amazon EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.
	//
	// You can specify the template using either the template name or ID. For help with creating a launch template, see [Creating a Launch Template for an Auto Scaling Group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon Elastic Compute Cloud Auto Scaling User Guide* . After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	//
	// > If you specify network interfaces in your launch template, you must explicitly set the property `AssociatePublicIpAddress` to "true". If no network interface is specified in the launch template, GameLift FleetIQ uses your account's default VPC.
	LaunchTemplate interface{} `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) for an IAM role that allows Amazon GameLift to access your Amazon EC2 Auto Scaling groups.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting.
	//
	// The scaling policy uses the metric `"PercentUtilizedGameServers"` to maintain a buffer of idle game servers that can immediately accommodate new games and players. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	AutoScalingPolicy interface{} `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.
	//
	// Method options include the following:
	//
	// - `SPOT_ONLY` - Only Spot Instances are used in the game server group. If Spot Instances are unavailable or not viable for game hosting, the game server group provides no hosting capacity until Spot Instances can again be used. Until then, no new instances are started, and the existing nonviable Spot Instances are terminated (after current gameplay ends) and are not replaced.
	// - `SPOT_PREFERRED` - (default value) Spot Instances are used whenever available in the game server group. If Spot Instances are unavailable, the game server group continues to provide hosting capacity by falling back to On-Demand Instances. Existing nonviable Spot Instances are terminated (after current gameplay ends) and are replaced with new On-Demand Instances.
	// - `ON_DEMAND_ONLY` - Only On-Demand Instances are used in the game server group. No Spot Instances are used, even when available, while this balancing strategy is in force.
	BalancingStrategy *string `field:"optional" json:"balancingStrategy" yaml:"balancingStrategy"`
	// The type of delete to perform.
	//
	// To delete a game server group, specify the `DeleteOption` . Options include the following:
	//
	// - `SAFE_DELETE` – (default) Terminates the game server group and Amazon EC2 Auto Scaling group only when it has no game servers that are in `UTILIZED` status.
	// - `FORCE_DELETE` – Terminates the game server group, including all active game servers regardless of their utilization status, and the Amazon EC2 Auto Scaling group.
	// - `RETAIN` – Does a safe delete of the game server group but retains the Amazon EC2 Auto Scaling group as is.
	DeleteOption *string `field:"optional" json:"deleteOption" yaml:"deleteOption"`
	// A flag that indicates whether instances in the game server group are protected from early termination.
	//
	// Unprotected instances that have active game servers running might be terminated during a scale-down event, causing players to be dropped from the game. Protected instances cannot be terminated while there are active game servers running except in the event of a forced game server group deletion (see ). An exception to this is with Spot Instances, which can be terminated by AWS regardless of protection status.
	GameServerProtectionPolicy *string `field:"optional" json:"gameServerProtectionPolicy" yaml:"gameServerProtectionPolicy"`
	// The maximum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances allowed in the Amazon EC2 Auto Scaling group.
	//
	// During automatic scaling events, GameLift FleetIQ and Amazon EC2 do not scale down the group below this minimum. In production, this value should be set to at least 1. After the Auto Scaling group is created, update this value directly in the Auto Scaling group using the AWS console or APIs.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A list of labels to assign to the new game server group resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management, and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags, respectively. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of virtual private cloud (VPC) subnets to use with instances in the game server group.
	//
	// By default, all GameLift FleetIQ-supported Availability Zones are used. You can use this parameter to specify VPCs that you've set up. This property cannot be updated after the game server group is created, and the corresponding Auto Scaling group will always use the property value that is set with this request, even if the Auto Scaling group is updated directly.
	VpcSubnets *[]*string `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

