package previewawspipesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The parameters for using an Amazon ECS task as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeTargetEcsTaskParametersProperty := &PipeTargetEcsTaskParametersProperty{
//   	CapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	EnableEcsManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	Group: jsii.String("group"),
//   	LaunchType: jsii.String("launchType"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	Overrides: &EcsTaskOverrideProperty{
//   		ContainerOverrides: []interface{}{
//   			&EcsContainerOverrideProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Cpu: jsii.Number(123),
//   				Environment: []interface{}{
//   					&EcsEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				EnvironmentFiles: []interface{}{
//   					&EcsEnvironmentFileProperty{
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Memory: jsii.Number(123),
//   				MemoryReservation: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				ResourceRequirements: []interface{}{
//   					&EcsResourceRequirementProperty{
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Cpu: jsii.String("cpu"),
//   		EphemeralStorage: &EcsEphemeralStorageProperty{
//   			SizeInGiB: jsii.Number(123),
//   		},
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		InferenceAcceleratorOverrides: []interface{}{
//   			&EcsInferenceAcceleratorOverrideProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				DeviceType: jsii.String("deviceType"),
//   			},
//   		},
//   		Memory: jsii.String("memory"),
//   		TaskRoleArn: jsii.String("taskRoleArn"),
//   	},
//   	PlacementConstraints: []interface{}{
//   		&PlacementConstraintProperty{
//   			Expression: jsii.String("expression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	PlacementStrategy: []interface{}{
//   		&PlacementStrategyProperty{
//   			Field: jsii.String("field"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	PlatformVersion: jsii.String("platformVersion"),
//   	PropagateTags: jsii.String("propagateTags"),
//   	ReferenceId: jsii.String("referenceId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskCount: jsii.Number(123),
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html
//
type CfnPipePropsMixin_PipeTargetEcsTaskParametersProperty struct {
	// The capacity provider strategy to use for the task.
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or launchType is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-capacityproviderstrategy
	//
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Specifies whether to enable Amazon ECS managed tags for the task.
	//
	// For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon Elastic Container Service Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-enableecsmanagedtags
	//
	// Default: - false.
	//
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Whether or not to enable the execute command functionality for the containers in this task.
	//
	// If true, this enables execute command functionality on all containers in the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-enableexecutecommand
	//
	// Default: - false.
	//
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies an Amazon ECS task group for the task.
	//
	// The maximum length is 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-group
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Specifies the launch type on which your task is running.
	//
	// The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The `FARGATE` value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. For more information, see [AWS Fargate on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS-Fargate.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-launchtype
	//
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// Use this structure if the Amazon ECS task uses the `awsvpc` network mode.
	//
	// This structure specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. This structure is required if `LaunchType` is `FARGATE` because the `awsvpc` mode is required for Fargate tasks.
	//
	// If you specify `NetworkConfiguration` when the target ECS task does not use the `awsvpc` network mode, the task fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The overrides that are associated with a task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// An array of placement constraint objects to use for the task.
	//
	// You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-placementconstraints
	//
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for the task.
	//
	// You can specify a maximum of five strategy rules per task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-placementstrategy
	//
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// Specifies the platform version for the task.
	//
	// Specify only the numeric portion of the platform version, such as `1.1.0` .
	//
	// This structure is used only if `LaunchType` is `FARGATE` . For more information about valid platform versions, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-platformversion
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the `TagResource` API action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The reference ID to use for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-referenceid
	//
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. To learn more, see [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-tags) in the Amazon ECS API Reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The number of tasks to create based on `TaskDefinition` .
	//
	// The default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-taskcount
	//
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
	// The ARN of the task definition to use if the event target is an Amazon ECS task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetecstaskparameters.html#cfn-pipes-pipe-pipetargetecstaskparameters-taskdefinitionarn
	//
	TaskDefinitionArn *string `field:"optional" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
}

