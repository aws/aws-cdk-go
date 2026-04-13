package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDaemonPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDaemonMixinProps := &CfnDaemonMixinProps{
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
type CfnDaemonMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-capacityproviderarns
	//
	CapacityProviderArns *[]*string `field:"optional" json:"capacityProviderArns" yaml:"capacityProviderArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-daemonname
	//
	DaemonName *string `field:"optional" json:"daemonName" yaml:"daemonName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-daemontaskdefinitionarn
	//
	DaemonTaskDefinitionArn *string `field:"optional" json:"daemonTaskDefinitionArn" yaml:"daemonTaskDefinitionArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-enableecsmanagedtags
	//
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-enableexecutecommand
	//
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemon.html#cfn-ecs-daemon-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

