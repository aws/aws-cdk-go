package awsevents


// The custom parameters to be used when the target is an Amazon ECS task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsParametersProperty := &ecsParametersProperty{
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
//   		awsVpcConfiguration: &awsVpcConfigurationProperty{
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
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			expression: jsii.String("expression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	placementStrategies: []interface{}{
//   		&placementStrategyProperty{
//   			field: jsii.String("field"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	referenceId: jsii.String("referenceId"),
//   	tagList: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskCount: jsii.Number(123),
//   }
//
type CfnRule_EcsParametersProperty struct {
	// The ARN of the task definition to use if the event target is an Amazon ECS task.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The capacity provider strategy to use for the task.
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or launchType is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Specifies whether to enable Amazon ECS managed tags for the task.
	//
	// For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon Elastic Container Service Developer Guide.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Whether or not to enable the execute command functionality for the containers in this task.
	//
	// If true, this enables execute command functionality on all containers in the task.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies an ECS task group for the task.
	//
	// The maximum length is 255 characters.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Specifies the launch type on which your task is running.
	//
	// The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The `FARGATE` value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. For more information, see [AWS Fargate on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS-Fargate.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// Use this structure if the Amazon ECS task uses the `awsvpc` network mode.
	//
	// This structure specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. This structure is required if `LaunchType` is `FARGATE` because the `awsvpc` mode is required for Fargate tasks.
	//
	// If you specify `NetworkConfiguration` when the target ECS task does not use the `awsvpc` network mode, the task fails.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for the task.
	//
	// You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for the task.
	//
	// You can specify a maximum of five strategy rules per task.
	PlacementStrategies interface{} `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// Specifies the platform version for the task.
	//
	// Specify only the numeric portion of the platform version, such as `1.1.0` .
	//
	// This structure is used only if `LaunchType` is `FARGATE` . For more information about valid platform versions, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the TagResource API action.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The reference ID to use for the task.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. To learn more, see [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-tags) in the Amazon ECS API Reference.
	TagList interface{} `field:"optional" json:"tagList" yaml:"tagList"`
	// The number of tasks to create based on `TaskDefinition` .
	//
	// The default is 1.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

