package awspipes

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetEcsTaskParametersProperty := &pipeTargetEcsTaskParametersProperty{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	capacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	enableEcsManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	group: jsii.String("group"),
//   	launchType: jsii.String("launchType"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsvpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	overrides: &ecsTaskOverrideProperty{
//   		containerOverrides: []interface{}{
//   			&ecsContainerOverrideProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				cpu: jsii.Number(123),
//   				environment: []interface{}{
//   					&ecsEnvironmentVariableProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				environmentFiles: []interface{}{
//   					&ecsEnvironmentFileProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				memory: jsii.Number(123),
//   				memoryReservation: jsii.Number(123),
//   				name: jsii.String("name"),
//   				resourceRequirements: []interface{}{
//   					&ecsResourceRequirementProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		cpu: jsii.String("cpu"),
//   		ephemeralStorage: &ecsEphemeralStorageProperty{
//   			sizeInGiB: jsii.Number(123),
//   		},
//   		executionRoleArn: jsii.String("executionRoleArn"),
//   		inferenceAcceleratorOverrides: []interface{}{
//   			&ecsInferenceAcceleratorOverrideProperty{
//   				deviceName: jsii.String("deviceName"),
//   				deviceType: jsii.String("deviceType"),
//   			},
//   		},
//   		memory: jsii.String("memory"),
//   		taskRoleArn: jsii.String("taskRoleArn"),
//   	},
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			expression: jsii.String("expression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	placementStrategy: []interface{}{
//   		&placementStrategyProperty{
//   			field: jsii.String("field"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	referenceId: jsii.String("referenceId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskCount: jsii.Number(123),
//   }
//
type CfnPipe_PipeTargetEcsTaskParametersProperty struct {
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.TaskDefinitionArn`.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.CapacityProviderStrategy`.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.EnableECSManagedTags`.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.EnableExecuteCommand`.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.Group`.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.LaunchType`.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.NetworkConfiguration`.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.Overrides`.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.PlacementConstraints`.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.PlacementStrategy`.
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.PlatformVersion`.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.PropagateTags`.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.ReferenceId`.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `CfnPipe.PipeTargetEcsTaskParametersProperty.TaskCount`.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

