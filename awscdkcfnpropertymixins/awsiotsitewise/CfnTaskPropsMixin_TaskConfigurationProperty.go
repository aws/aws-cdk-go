package awsiotsitewise


// The task execution configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   taskConfigurationProperty := &TaskConfigurationProperty{
//   	ContainerTaskConfiguration: &ContainerTaskConfigurationProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EcrUri: jsii.String("ecrUri"),
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		EphemeralStorageConfiguration: &EphemeralStorageConfigurationProperty{
//   			StorageClass: jsii.String("storageClass"),
//   			StorageSizeInGiB: jsii.Number(123),
//   		},
//   		Mounts: []interface{}{
//   			&MountProperty{
//   				Name: jsii.String("name"),
//   				RelativePath: jsii.String("relativePath"),
//   				Source: &MountSourceProperty{
//   					S3AccessPoint: &S3AccessPointSourceProperty{
//   						AccessPointArn: jsii.String("accessPointArn"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   				},
//   				StorageType: jsii.String("storageType"),
//   			},
//   		},
//   		ProcessingType: jsii.String("processingType"),
//   		ProcessingUnit: jsii.String("processingUnit"),
//   		TaskExecutionRole: jsii.String("taskExecutionRole"),
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-taskconfiguration.html
//
type CfnTaskPropsMixin_TaskConfigurationProperty struct {
	// Configuration for running a custom container image on managed compute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-taskconfiguration.html#cfn-iotsitewise-task-taskconfiguration-containertaskconfiguration
	//
	ContainerTaskConfiguration interface{} `field:"optional" json:"containerTaskConfiguration" yaml:"containerTaskConfiguration"`
}

