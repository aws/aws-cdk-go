package awspipes


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
	// `CfnPipe.EcsTaskOverrideProperty.ContainerOverrides`.
	ContainerOverrides interface{} `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// `CfnPipe.EcsTaskOverrideProperty.Cpu`.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// `CfnPipe.EcsTaskOverrideProperty.EphemeralStorage`.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// `CfnPipe.EcsTaskOverrideProperty.ExecutionRoleArn`.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// `CfnPipe.EcsTaskOverrideProperty.InferenceAcceleratorOverrides`.
	InferenceAcceleratorOverrides interface{} `field:"optional" json:"inferenceAcceleratorOverrides" yaml:"inferenceAcceleratorOverrides"`
	// `CfnPipe.EcsTaskOverrideProperty.Memory`.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// `CfnPipe.EcsTaskOverrideProperty.TaskRoleArn`.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

