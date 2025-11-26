package previewawsschedulermixins


// The templated target type for the Amazon ECS [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   ecsParametersProperty := &EcsParametersProperty{
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
//   	Tags: tags,
//   	TaskCount: jsii.Number(123),
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html
//
type CfnSchedulePropsMixin_EcsParametersProperty struct {
	// The capacity provider strategy to use for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-capacityproviderstrategy
	//
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Specifies whether to enable Amazon ECS managed tags for the task.
	//
	// For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon ECS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-enableecsmanagedtags
	//
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Whether or not to enable the execute command functionality for the containers in this task.
	//
	// If true, this enables execute command functionality on all containers in the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-enableexecutecommand
	//
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies an Amazon ECS task group for the task.
	//
	// The maximum length is 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-group
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Specifies the launch type on which your task is running.
	//
	// The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The `FARGATE` value is supported only in the Regions where Fargate with Amazon ECS is supported. For more information, see [AWS Fargate on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS_Fargate.html) in the *Amazon ECS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-launchtype
	//
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// This structure specifies the network configuration for an ECS task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for the task.
	//
	// You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-placementconstraints
	//
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The task placement strategy for a task or service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-placementstrategy
	//
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// Specifies the platform version for the task.
	//
	// Specify only the numeric portion of the platform version, such as `1.1.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-platformversion
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the Amazon ECS [`TagResource`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The reference ID to use for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-referenceid
	//
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. For more information, see [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) in the *Amazon ECS API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The number of tasks to create based on `TaskDefinition` .
	//
	// The default is `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-taskcount
	//
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
	// The Amazon Resource Name (ARN) of the task definition to use if the event target is an Amazon ECS task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-ecsparameters.html#cfn-scheduler-schedule-ecsparameters-taskdefinitionarn
	//
	TaskDefinitionArn *string `field:"optional" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
}

