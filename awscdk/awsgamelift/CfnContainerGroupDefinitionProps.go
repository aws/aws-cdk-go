package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContainerGroupDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerGroupDefinitionProps := &CfnContainerGroupDefinitionProps{
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	TotalMemoryLimitMebibytes: jsii.Number(123),
//   	TotalVcpuLimit: jsii.Number(123),
//
//   	// the properties below are optional
//   	ContainerGroupType: jsii.String("containerGroupType"),
//   	GameServerContainerDefinition: &GameServerContainerDefinitionProperty{
//   		ContainerName: jsii.String("containerName"),
//   		ImageUri: jsii.String("imageUri"),
//   		ServerSdkVersion: jsii.String("serverSdkVersion"),
//
//   		// the properties below are optional
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
//   		MountPoints: []interface{}{
//   			&ContainerMountPointProperty{
//   				InstancePath: jsii.String("instancePath"),
//
//   				// the properties below are optional
//   				AccessLevel: jsii.String("accessLevel"),
//   				ContainerPath: jsii.String("containerPath"),
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
//   	},
//   	SourceVersionNumber: jsii.Number(123),
//   	SupportContainerDefinitions: []interface{}{
//   		&SupportContainerDefinitionProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ImageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
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
//
//   				// the properties below are optional
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			MemoryHardLimitMebibytes: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&ContainerMountPointProperty{
//   					InstancePath: jsii.String("instancePath"),
//
//   					// the properties below are optional
//   					AccessLevel: jsii.String("accessLevel"),
//   					ContainerPath: jsii.String("containerPath"),
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionDescription: jsii.String("versionDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html
//
type CfnContainerGroupDefinitionProps struct {
	// A descriptive identifier for the container group definition.
	//
	// The name value is unique in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform that all containers in the container group definition run on.
	//
	// > Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in the [Amazon Linux 2 FAQs](https://docs.aws.amazon.com/https://aws.amazon.com/amazon-linux-2/faqs/) . For game servers that are hosted on AL2 and use Amazon GameLift server SDK 4.x, first update the game server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to Amazon GameLift server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-operatingsystem
	//
	OperatingSystem *string `field:"required" json:"operatingSystem" yaml:"operatingSystem"`
	// The total memory limit of container groups following this definition in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalmemorylimitmebibytes
	//
	TotalMemoryLimitMebibytes *float64 `field:"required" json:"totalMemoryLimitMebibytes" yaml:"totalMemoryLimitMebibytes"`
	// The total amount of virtual CPUs on the container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalvcpulimit
	//
	TotalVcpuLimit *float64 `field:"required" json:"totalVcpuLimit" yaml:"totalVcpuLimit"`
	// The scope of the container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-containergrouptype
	//
	ContainerGroupType *string `field:"optional" json:"containerGroupType" yaml:"containerGroupType"`
	// Specifies the information required to run game servers with this container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition
	//
	GameServerContainerDefinition interface{} `field:"optional" json:"gameServerContainerDefinition" yaml:"gameServerContainerDefinition"`
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
	// The description of this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

