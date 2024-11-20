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
//   				DesiredEc2Instances: jsii.Number(123),
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//   			},
//   			StoppedActions: []*string{
//   				jsii.String("stoppedActions"),
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDestination: jsii.String("logDestination"),
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
	// A unique identifier for an AWS IAM role that manages access to your AWS services.
	//
	// Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-fleetrolearn
	//
	FleetRoleArn *string `field:"required" json:"fleetRoleArn" yaml:"fleetRoleArn"`
	// Indicates whether to use On-Demand instances or Spot instances for this fleet.
	//
	// If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-billingtype
	//
	BillingType *string `field:"optional" json:"billingType" yaml:"billingType"`
	// Provides details about how to drain old tasks and replace them with new updated tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// A human-readable description of a fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the container group definition that will be created per game server.
	//
	// You must specify GAME_SERVER container group. You have the option to also specify one PER_INSTANCE container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupdefinitionname
	//
	GameServerContainerGroupDefinitionName *string `field:"optional" json:"gameServerContainerGroupDefinitionName" yaml:"gameServerContainerGroupDefinitionName"`
	// The number of desired game server container groups per instance, a number between 1-5000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupsperinstance
	//
	GameServerContainerGroupsPerInstance *float64 `field:"optional" json:"gameServerContainerGroupsPerInstance" yaml:"gameServerContainerGroupsPerInstance"`
	// A policy that limits the number of game sessions a player can create on the same fleet.
	//
	// This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period".
	//
	// The policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a CreateGameSession request, Amazon GameLift checks that the player (identified by CreatorId) has created fewer than 10 game sessions in the past 60 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy
	//
	GameSessionCreationLimitPolicy interface{} `field:"optional" json:"gameSessionCreationLimitPolicy" yaml:"gameSessionCreationLimitPolicy"`
	// Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceconnectionportrange
	//
	InstanceConnectionPortRange interface{} `field:"optional" json:"instanceConnectionPortRange" yaml:"instanceConnectionPortRange"`
	// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceinboundpermissions
	//
	InstanceInboundPermissions interface{} `field:"optional" json:"instanceInboundPermissions" yaml:"instanceInboundPermissions"`
	// The name of an EC2 instance type that is supported in Amazon GameLift.
	//
	// A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-locations
	//
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// A policy the location and provider of logs from the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The name of an Amazon CloudWatch metric group.
	//
	// A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-metricgroups
	//
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// A game session protection policy to apply to all game sessions hosted on instances in this fleet.
	//
	// When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-newgamesessionprotectionpolicy
	//
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// The name of the container group definition that will be created per instance.
	//
	// This field is optional if you specify GameServerContainerGroupDefinitionName.
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

