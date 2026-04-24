package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDaemon`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDaemonProps := &CfnDaemonProps{
//   	CapacityProviderArns: []*string{
//   		jsii.String("capacityProviderArns"),
//   	},
//   	ClusterArn: jsii.String("clusterArn"),
//   	DaemonName: jsii.String("daemonName"),
//   	DaemonTaskDefinitionArn: jsii.String("daemonTaskDefinitionArn"),
//   	DeploymentConfiguration: &DaemonDeploymentConfigurationProperty{
//   		Alarms: &DaemonAlarmConfigurationProperty{
//   			AlarmNames: []*string{
//   				jsii.String("alarmNames"),
//   			},
//   			Enable: jsii.Boolean(false),
//   		},
//   		BakeTimeInMinutes: jsii.Number(123),
//   		DrainPercent: jsii.Number(123),
//   	},
//   	EnableEcsManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	PropagateTags: jsii.String("propagateTags"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html
//
type CfnDaemonProps struct {
	// The Amazon Resource Names (ARNs) of the capacity providers associated with the daemon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-capacityproviderarns
	//
	CapacityProviderArns *[]*string `field:"optional" json:"capacityProviderArns" yaml:"capacityProviderArns"`
	// The Amazon Resource Name (ARN) of the cluster that the daemon is running in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-daemonname
	//
	DaemonName *string `field:"optional" json:"daemonName" yaml:"daemonName"`
	// The Amazon Resource Name (ARN) of the daemon task definition used by this revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-daemontaskdefinitionarn
	//
	DaemonTaskDefinitionArn *string `field:"optional" json:"daemonTaskDefinitionArn" yaml:"daemonTaskDefinitionArn"`
	// Optional deployment parameters that control how a daemon rolls out updates across container instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// Specifies whether Amazon ECS managed tags are turned on for the daemon tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-enableecsmanagedtags
	//
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Specifies whether the execute command functionality is turned on for the daemon tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-enableexecutecommand
	//
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies whether tags are propagated from the daemon to the daemon tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

