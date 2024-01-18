package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The common properties for all task definitions.
//
// For more information, see
// [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var proxyConfiguration proxyConfiguration
//   var role role
//
//   commonTaskDefinitionProps := &CommonTaskDefinitionProps{
//   	ExecutionRole: role,
//   	Family: jsii.String("family"),
//   	ProxyConfiguration: proxyConfiguration,
//   	TaskRole: role,
//   	Volumes: []volume{
//   		&volume{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ConfiguredAtLaunch: jsii.Boolean(false),
//   			DockerVolumeConfiguration: &DockerVolumeConfiguration{
//   				Driver: jsii.String("driver"),
//   				Scope: awscdk.Aws_ecs.Scope_TASK,
//
//   				// the properties below are optional
//   				Autoprovision: jsii.Boolean(false),
//   				DriverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				Labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   			},
//   			EfsVolumeConfiguration: &EfsVolumeConfiguration{
//   				FileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				AuthorizationConfig: &AuthorizationConfig{
//   					AccessPointId: jsii.String("accessPointId"),
//   					Iam: jsii.String("iam"),
//   				},
//   				RootDirectory: jsii.String("rootDirectory"),
//   				TransitEncryption: jsii.String("transitEncryption"),
//   				TransitEncryptionPort: jsii.Number(123),
//   			},
//   			Host: &Host{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   }
//
type CommonTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	// Default: - An execution role will be automatically created if you use ECR images in your task definition.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Default: - Automatically generated name.
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	// Default: - No proxy configuration.
	//
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Default: - A task role is automatically created for you.
	//
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	// Default: - No volumes are passed to the Docker daemon on a container instance.
	//
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

