package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The properties for a task definition run on an External cluster.
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
//   externalTaskDefinitionProps := &externalTaskDefinitionProps{
//   	executionRole: role,
//   	family: jsii.String("family"),
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	proxyConfiguration: proxyConfiguration,
//   	taskRole: role,
//   	volumes: []volume{
//   		&volume{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			dockerVolumeConfiguration: &dockerVolumeConfiguration{
//   				driver: jsii.String("driver"),
//   				scope: awscdk.*Aws_ecs.scope_TASK,
//
//   				// the properties below are optional
//   				autoprovision: jsii.Boolean(false),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   			},
//   			efsVolumeConfiguration: &efsVolumeConfiguration{
//   				fileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfig{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &host{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   }
//
type ExternalTaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The networking mode to use for the containers in the task.
	//
	// With ECS Anywhere, supported modes are bridge, host and none.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
}

