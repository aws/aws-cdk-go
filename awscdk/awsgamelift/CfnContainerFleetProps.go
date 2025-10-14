package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContainerFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerFleetProps := &CfnContainerFleetProps{
//   	FleetRoleArn: jsii.String("fleetRoleArn"),
//
//   	// the properties below are optional
//   	BillingType: jsii.String("billingType"),
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		ImpairmentStrategy: jsii.String("impairmentStrategy"),
//   		MinimumHealthyPercentage: jsii.Number(123),
//   		ProtectionStrategy: jsii.String("protectionStrategy"),
//   	},
//   	Description: jsii.String("description"),
//   	GameServerContainerGroupDefinitionName: jsii.String("gameServerContainerGroupDefinitionName"),
//   	GameServerContainerGroupsPerInstance: jsii.Number(123),
//   	GameSessionCreationLimitPolicy: &GameSessionCreationLimitPolicyProperty{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriodInMinutes: jsii.Number(123),
//   	},
//   	InstanceConnectionPortRange: &ConnectionPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	InstanceInboundPermissions: []interface{}{
//   		&IpPermissionProperty{
//   			FromPort: jsii.Number(123),
//   			IpRange: jsii.String("ipRange"),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	Locations: []interface{}{
//   		&LocationConfigurationProperty{
//   			Location: jsii.String("location"),
//
//   			// the properties below are optional
//   			LocationCapacity: &LocationCapacityProperty{
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//
//   				// the properties below are optional
//   				DesiredEc2Instances: jsii.Number(123),
//   			},
//   			StoppedActions: []*string{
//   				jsii.String("stoppedActions"),
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDestination: jsii.String("logDestination"),
//   		LogGroupArn: jsii.String("logGroupArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	MetricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	NewGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	PerInstanceContainerGroupDefinitionName: jsii.String("perInstanceContainerGroupDefinitionName"),
//   	ScalingPolicies: []interface{}{
//   		&ScalingPolicyProperty{
//   			MetricName: jsii.String("metricName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			EvaluationPeriods: jsii.Number(123),
//   			PolicyType: jsii.String("policyType"),
//   			ScalingAdjustment: jsii.Number(123),
//   			ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   			TargetConfiguration: &TargetConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//   			},
//   			Threshold: jsii.Number(123),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html
//
type CfnContainerFleetProps struct {
	// The unique identifier for an AWS Identity and Access Management (IAM) role with permissions to run your containers on resources that are managed by Amazon GameLift Servers.
	//
	// See [Set up an IAM service role](https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html) . This fleet property can't be changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-fleetrolearn
	//
	FleetRoleArn *string `field:"required" json:"fleetRoleArn" yaml:"fleetRoleArn"`
	// Indicates whether the fleet uses On-Demand or Spot instances for this fleet.
	//
	// Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot) . You can't update this fleet property.
	//
	// By default, this property is set to `ON_DEMAND` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-billingtype
	//
	BillingType *string `field:"optional" json:"billingType" yaml:"billingType"`
	// Set of rules for processing a deployment for a container fleet update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// A meaningful description of the container fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the fleet's game server container group definition, which describes how to deploy containers with your game server build and support software onto each fleet instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupdefinitionname
	//
	GameServerContainerGroupDefinitionName *string `field:"optional" json:"gameServerContainerGroupDefinitionName" yaml:"gameServerContainerGroupDefinitionName"`
	// The number of times to replicate the game server container group on each fleet instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupsperinstance
	//
	GameServerContainerGroupsPerInstance *float64 `field:"optional" json:"gameServerContainerGroupsPerInstance" yaml:"gameServerContainerGroupsPerInstance"`
	// A policy that limits the number of game sessions that each individual player can create on instances in this fleet.
	//
	// The limit applies for a specified span of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy
	//
	GameSessionCreationLimitPolicy interface{} `field:"optional" json:"gameSessionCreationLimitPolicy" yaml:"gameSessionCreationLimitPolicy"`
	// The set of port numbers to open on each instance in a container fleet.
	//
	// Connection ports are used by inbound traffic to connect with processes that are running in containers on the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceconnectionportrange
	//
	InstanceConnectionPortRange interface{} `field:"optional" json:"instanceConnectionPortRange" yaml:"instanceConnectionPortRange"`
	// The IP address ranges and port settings that allow inbound traffic to access game server processes and other processes on this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceinboundpermissions
	//
	InstanceInboundPermissions interface{} `field:"optional" json:"instanceInboundPermissions" yaml:"instanceInboundPermissions"`
	// The Amazon EC2 instance type to use for all instances in the fleet.
	//
	// Instance type determines the computing resources and processing power that's available to host your game servers. This includes including CPU, memory, storage, and networking capacity. You can't update this fleet property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-locations
	//
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// The method that is used to collect container logs for the fleet.
	//
	// Amazon GameLift Servers saves all standard output for each container in logs, including game session logs.
	//
	// - `CLOUDWATCH` -- Send logs to an Amazon CloudWatch log group that you define. Each container emits a log stream, which is organized in the log group.
	// - `S3` -- Store logs in an Amazon S3 bucket that you define.
	// - `NONE` -- Don't collect container logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// Metric groups aggregate metrics for multiple fleets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-metricgroups
	//
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// Determines whether Amazon GameLift Servers can shut down game sessions on the fleet that are actively running and hosting players.
	//
	// Amazon GameLift Servers might prompt an instance shutdown when scaling down fleet capacity or when retiring unhealthy instances. You can also set game session protection for individual game sessions using [UpdateGameSession](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateGameSession.html) .
	//
	// - *NoProtection* -- Game sessions can be shut down during active gameplay.
	// - *FullProtection* -- Game sessions in `ACTIVE` status can't be shut down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-newgamesessionprotectionpolicy
	//
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// The name of the fleet's per-instance container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-perinstancecontainergroupdefinitionname
	//
	PerInstanceContainerGroupDefinitionName *string `field:"optional" json:"perInstanceContainerGroupDefinitionName" yaml:"perInstanceContainerGroupDefinitionName"`
	// A list of rules that control how a fleet is scaled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-scalingpolicies
	//
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

