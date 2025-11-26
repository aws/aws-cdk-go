package previewawspipesmixins


// The overrides that are associated with a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecsTaskOverrideProperty := &EcsTaskOverrideProperty{
//   	ContainerOverrides: []interface{}{
//   		&EcsContainerOverrideProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			Environment: []interface{}{
//   				&EcsEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EnvironmentFiles: []interface{}{
//   				&EcsEnvironmentFileProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Memory: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			ResourceRequirements: []interface{}{
//   				&EcsResourceRequirementProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Cpu: jsii.String("cpu"),
//   	EphemeralStorage: &EcsEphemeralStorageProperty{
//   		SizeInGiB: jsii.Number(123),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	InferenceAcceleratorOverrides: []interface{}{
//   		&EcsInferenceAcceleratorOverrideProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			DeviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	Memory: jsii.String("memory"),
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html
//
type CfnPipePropsMixin_EcsTaskOverrideProperty struct {
	// One or more container overrides that are sent to a task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-containeroverrides
	//
	ContainerOverrides interface{} `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The cpu override for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-cpu
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ephemeral storage setting override for the task.
	//
	// > This parameter is only supported for tasks hosted on Fargate that use the following platform versions:
	// >
	// > - Linux platform version `1.4.0` or later.
	// > - Windows platform version `1.0.0` or later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution IAM role override for the task.
	//
	// For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The Elastic Inference accelerator override for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-inferenceacceleratoroverrides
	//
	InferenceAcceleratorOverrides interface{} `field:"optional" json:"inferenceAcceleratorOverrides" yaml:"inferenceAcceleratorOverrides"`
	// The memory override for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-memory
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The Amazon Resource Name (ARN) of the IAM role that containers in this task can assume.
	//
	// All containers in this task are granted the permissions that are specified in this role. For more information, see [IAM Role for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecstaskoverride.html#cfn-pipes-pipe-ecstaskoverride-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

