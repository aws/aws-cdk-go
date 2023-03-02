package awspipes


// The overrides that are associated with a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsTaskOverrideProperty := &ecsTaskOverrideProperty{
//   	containerOverrides: []interface{}{
//   		&ecsContainerOverrideProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			environment: []interface{}{
//   				&ecsEnvironmentVariableProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			environmentFiles: []interface{}{
//   				&ecsEnvironmentFileProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			memory: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   			name: jsii.String("name"),
//   			resourceRequirements: []interface{}{
//   				&ecsResourceRequirementProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	cpu: jsii.String("cpu"),
//   	ephemeralStorage: &ecsEphemeralStorageProperty{
//   		sizeInGiB: jsii.Number(123),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	inferenceAcceleratorOverrides: []interface{}{
//   		&ecsInferenceAcceleratorOverrideProperty{
//   			deviceName: jsii.String("deviceName"),
//   			deviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	memory: jsii.String("memory"),
//   	taskRoleArn: jsii.String("taskRoleArn"),
//   }
//
type CfnPipe_EcsTaskOverrideProperty struct {
	// One or more container overrides that are sent to a task.
	ContainerOverrides interface{} `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The cpu override for the task.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ephemeral storage setting override for the task.
	//
	// > This parameter is only supported for tasks hosted on Fargate that use the following platform versions:
	// >
	// > - Linux platform version `1.4.0` or later.
	// > - Windows platform version `1.0.0` or later.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution IAM role override for the task.
	//
	// For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The Elastic Inference accelerator override for the task.
	InferenceAcceleratorOverrides interface{} `field:"optional" json:"inferenceAcceleratorOverrides" yaml:"inferenceAcceleratorOverrides"`
	// The memory override for the task.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The Amazon Resource Name (ARN) of the IAM role that containers in this task can assume.
	//
	// All containers in this task are granted the permissions that are specified in this role. For more information, see [IAM Role for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

