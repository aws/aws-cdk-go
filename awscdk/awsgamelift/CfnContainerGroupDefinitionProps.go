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
//   var supportContainerDefinitions interface{}
//
//   cfnContainerGroupDefinitionProps := &CfnContainerGroupDefinitionProps{
//   	ContainerDefinitions: []interface{}{
//   		&ContainerDefinitionProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ImageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Environment: []interface{}{
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
//   			MemoryLimits: &MemoryLimitsProperty{
//   				HardLimit: jsii.Number(123),
//   				SoftLimit: jsii.Number(123),
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
//   			WorkingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	TotalCpuLimit: jsii.Number(123),
//   	TotalMemoryLimit: jsii.Number(123),
//
//   	// the properties below are optional
//   	SchedulingStrategy: jsii.String("schedulingStrategy"),
//   	SourceVersionNumber: jsii.Number(123),
//   	SupportContainerDefinitions: []interface{}{
//   		supportContainerDefinitions,
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html
//
type CfnContainerGroupDefinitionProps struct {
	// The set of container definitions that are included in the container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinitions
	//
	ContainerDefinitions interface{} `field:"required" json:"containerDefinitions" yaml:"containerDefinitions"`
	// A descriptive identifier for the container group definition.
	//
	// The name value is unique in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform required for all containers in the container group definition.
	//
	// > Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in the [Amazon Linux 2 FAQs](https://docs.aws.amazon.com/https://aws.amazon.com/amazon-linux-2/faqs/) . For game servers that are hosted on AL2 and use Amazon GameLift server SDK 4.x., first update the game server build to server SDK 5.x, and then deploy to AL2023 instances. See [Migrate to Amazon GameLift server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-operatingsystem
	//
	OperatingSystem *string `field:"required" json:"operatingSystem" yaml:"operatingSystem"`
	// The amount of CPU units on a fleet instance to allocate for the container group.
	//
	// All containers in the group share these resources. This property is an integer value in CPU units (1 vCPU is equal to 1024 CPU units).
	//
	// You can set additional limits for each `ContainerDefinition` in the group. If individual containers have limits, this value must be equal to or greater than the sum of all container-specific CPU limits in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalcpulimit
	//
	TotalCpuLimit *float64 `field:"required" json:"totalCpuLimit" yaml:"totalCpuLimit"`
	// The amount of memory (in MiB) on a fleet instance to allocate for the container group.
	//
	// All containers in the group share these resources.
	//
	// You can set additional limits for each `ContainerDefinition` in the group. If individual containers have limits, this value must meet the following requirements:
	//
	// - Equal to or greater than the sum of all container-specific soft memory limits in the group.
	// - Equal to or greater than any container-specific hard limits in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-totalmemorylimit
	//
	TotalMemoryLimit *float64 `field:"required" json:"totalMemoryLimit" yaml:"totalMemoryLimit"`
	// The method for deploying the container group across fleet instances.
	//
	// A replica container group might have multiple copies on each fleet instance. A daemon container group maintains only one copy per fleet instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-schedulingstrategy
	//
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// A specific ContainerGroupDefinition version to be updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-sourceversionnumber
	//
	SourceVersionNumber *float64 `field:"optional" json:"sourceVersionNumber" yaml:"sourceVersionNumber"`
	// A collection of support container definitions that define the containers in this group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinitions
	//
	SupportContainerDefinitions interface{} `field:"optional" json:"supportContainerDefinitions" yaml:"supportContainerDefinitions"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html#cfn-gamelift-containergroupdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

