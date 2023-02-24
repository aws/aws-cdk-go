// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of a job definition container.
//
// Example:
//   dbSecret := secretsmanager.NewSecret(this, jsii.String("secret"))
//
//   batch.NewJobDefinition(this, jsii.String("batch-job-def-secrets"), &JobDefinitionProps{
//   	Container: &JobDefinitionContainer{
//   		Image: ecs.EcrImage_FromRegistry(jsii.String("docker/whalesay")),
//   		Secrets: map[string]secret{
//   			"PASSWORD": ecs.*secret_fromSecretsManager(dbSecret, jsii.String("password")),
//   		},
//   	},
//   })
//
// Experimental.
type JobDefinitionContainer struct {
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// Whether or not to assign a public IP to the job.
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The IAM role that AWS Batch can assume.
	//
	// Required when using Fargate.
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all
	// containers in a job should not exceed the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// Currently all node groups in a
	// multi-node parallel job must use the same instance type. This parameter is not valid
	// for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The IAM role that the container can assume for AWS permissions.
	// Experimental.
	JobRole awsiam.IRole `field:"optional" json:"jobRole" yaml:"jobRole"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	//
	// For now, only the `devices` property is supported.
	// Experimental.
	LinuxParams awsecs.LinuxParameters `field:"optional" json:"linuxParams" yaml:"linuxParams"`
	// The log configuration specification for the container.
	// Experimental.
	LogConfiguration *LogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The hard limit (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed
	// the memory specified here, the container is killed. You must specify at least 4 MiB of memory for EC2 and 512 MiB for Fargate.
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The mount points for data volumes in your container.
	// Experimental.
	MountPoints *[]*awsecs.MountPoint `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Fargate platform version.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	// Experimental.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	// Experimental.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The environment variables from secrets manager or ssm parameter store.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// A list of ulimits to set in the container.
	// Experimental.
	Ulimits *[]*awsecs.Ulimit `field:"optional" json:"ulimits" yaml:"ulimits"`
	// The user name to use inside the container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to
	// 1,024 CPU shares. You must specify at least one vCPU for EC2 and 0.25 for Fargate.
	// Experimental.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
	// A list of data volumes used in a job.
	// Experimental.
	Volumes *[]*awsecs.Volume `field:"optional" json:"volumes" yaml:"volumes"`
}

