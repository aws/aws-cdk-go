package previewawsgameliftmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContainerGroupDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContainerGroupDefinitionMixinProps := &CfnContainerGroupDefinitionMixinProps{
//   	ContainerGroupType: jsii.String("containerGroupType"),
//   	GameServerContainerDefinition: &GameServerContainerDefinitionProperty{
//   		ContainerName: jsii.String("containerName"),
//   		DependsOn: []interface{}{
//   			&ContainerDependencyProperty{
//   				Condition: jsii.String("condition"),
//   				ContainerName: jsii.String("containerName"),
//   			},
//   		},
//   		EnvironmentOverride: []interface{}{
//   			&ContainerEnvironmentProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		MountPoints: []interface{}{
//   			&ContainerMountPointProperty{
//   				AccessLevel: jsii.String("accessLevel"),
//   				ContainerPath: jsii.String("containerPath"),
//   				InstancePath: jsii.String("instancePath"),
//   			},
//   		},
//   		PortConfiguration: &PortConfigurationProperty{
//   			ContainerPortRanges: []interface{}{
//   				&ContainerPortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   		ServerSdkVersion: jsii.String("serverSdkVersion"),
//   	},
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	SourceVersionNumber: jsii.Number(123),
//   	SupportContainerDefinitions: []interface{}{
//   		&SupportContainerDefinitionProperty{
//   			ContainerName: jsii.String("containerName"),
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EnvironmentOverride: []interface{}{
//   				&ContainerEnvironmentProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			HealthCheck: &ContainerHealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			ImageUri: jsii.String("imageUri"),
//   			MemoryHardLimitMebibytes: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&ContainerMountPointProperty{
//   					AccessLevel: jsii.String("accessLevel"),
//   					ContainerPath: jsii.String("containerPath"),
//   					InstancePath: jsii.String("instancePath"),
//   				},
//   			},
//   			PortConfiguration: &PortConfigurationProperty{
//   				ContainerPortRanges: []interface{}{
//   					&ContainerPortRangeProperty{
//   						FromPort: jsii.Number(123),
//   						Protocol: jsii.String("protocol"),
//   						ToPort: jsii.Number(123),
//   					},
//   				},
//   			},
//   			ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   			Vcpu: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TotalMemoryLimitMebibytes: jsii.Number(123),
//   	TotalVcpuLimit: jsii.Number(123),
//   	VersionDescription: jsii.String("versionDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html
//
type CfnContainerGroupDefinitionMixinProps struct {
	// The type of container group.
	//
	// Container group type determines how Amazon GameLift Servers deploys the container group on each fleet instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-containergrouptype
	//
	ContainerGroupType *string `field:"optional" json:"containerGroupType" yaml:"containerGroupType"`
	// The definition for the game server container in this group.
	//
	// This property is used only when the container group type is `GAME_SERVER` . This container definition specifies a container image with the game server build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition
	//
	GameServerContainerDefinition interface{} `field:"optional" json:"gameServerContainerDefinition" yaml:"gameServerContainerDefinition"`
	// A descriptive identifier for the container group definition.
	//
	// The name value is unique in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The platform that all containers in the container group definition run on.
	//
	// > Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in the [Amazon Linux 2 FAQs](https://docs.aws.amazon.com/amazon-linux-2/faqs/) . For game servers that are hosted on AL2 and use server SDK version 4.x for Amazon GameLift Servers, first update the game server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-operatingsystem
	//
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// A specific ContainerGroupDefinition version to be updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-sourceversionnumber
	//
	SourceVersionNumber *float64 `field:"optional" json:"sourceVersionNumber" yaml:"sourceVersionNumber"`
	// The set of definitions for support containers in this group.
	//
	// A container group definition might have zero support container definitions. Support container can be used in any type of container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinitions
	//
	SupportContainerDefinitions interface{} `field:"optional" json:"supportContainerDefinitions" yaml:"supportContainerDefinitions"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The amount of memory (in MiB) on a fleet instance to allocate for the container group.
	//
	// All containers in the group share these resources.
	//
	// You can set a limit for each container definition in the group. If individual containers have limits, this total value must be greater than any individual container's memory limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalmemorylimitmebibytes
	//
	TotalMemoryLimitMebibytes *float64 `field:"optional" json:"totalMemoryLimitMebibytes" yaml:"totalMemoryLimitMebibytes"`
	// The amount of vCPU units on a fleet instance to allocate for the container group (1 vCPU is equal to 1024 CPU units).
	//
	// All containers in the group share these resources. You can set a limit for each container definition in the group. If individual containers have limits, this total value must be equal to or greater than the sum of the limits for each container in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalvcpulimit
	//
	TotalVcpuLimit *float64 `field:"optional" json:"totalVcpuLimit" yaml:"totalVcpuLimit"`
	// An optional description that was provided for a container group definition update.
	//
	// Each version can have a unique description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

