package awsiotsitewise


// Configuration for running a custom container image on managed compute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerTaskConfigurationProperty := &ContainerTaskConfigurationProperty{
//   	EcrUri: jsii.String("ecrUri"),
//   	ProcessingType: jsii.String("processingType"),
//   	ProcessingUnit: jsii.String("processingUnit"),
//   	TaskExecutionRole: jsii.String("taskExecutionRole"),
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	EphemeralStorageConfiguration: &EphemeralStorageConfigurationProperty{
//   		StorageClass: jsii.String("storageClass"),
//   		StorageSizeInGiB: jsii.Number(123),
//   	},
//   	Mounts: []interface{}{
//   		&MountProperty{
//   			Name: jsii.String("name"),
//   			RelativePath: jsii.String("relativePath"),
//   			Source: &MountSourceProperty{
//   				S3AccessPoint: &S3AccessPointSourceProperty{
//   					AccessPointArn: jsii.String("accessPointArn"),
//
//   					// the properties below are optional
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   			StorageType: jsii.String("storageType"),
//   		},
//   	},
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html
//
type CfnTask_ContainerTaskConfigurationProperty struct {
	// The Amazon ECR image URI for the task container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-ecruri
	//
	EcrUri *string `field:"required" json:"ecrUri" yaml:"ecrUri"`
	// The processing type for compute resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-processingtype
	//
	ProcessingType *string `field:"required" json:"processingType" yaml:"processingType"`
	// The processing unit allocation that determines vCPU, memory, and GPU resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-processingunit
	//
	ProcessingUnit *string `field:"required" json:"processingUnit" yaml:"processingUnit"`
	// The ARN of the IAM role that grants the containerized workload permissions to access AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-taskexecutionrole
	//
	TaskExecutionRole *string `field:"required" json:"taskExecutionRole" yaml:"taskExecutionRole"`
	// The command to execute in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// A map of environment variable key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Configuration for ephemeral storage attached to the container task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-ephemeralstorageconfiguration
	//
	EphemeralStorageConfiguration interface{} `field:"optional" json:"ephemeralStorageConfiguration" yaml:"ephemeralStorageConfiguration"`
	// Mounts attached to the container filesystem.
	//
	// Each mount exposes an external data source as a local directory inside the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-mounts
	//
	Mounts interface{} `field:"optional" json:"mounts" yaml:"mounts"`
	// The timeout in seconds for task execution.
	//
	// Default: 3600 (1 hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-containertaskconfiguration.html#cfn-iotsitewise-task-containertaskconfiguration-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

