package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Example:
//   var cluster cluster
//
//   loadBalancedEcsService := ecsPatterns.NewNetworkLoadBalancedEc2Service(this, jsii.String("Service"), &NetworkLoadBalancedEc2ServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
//   		Environment: map[string]*string{
//   			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
//   			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
//   		},
//   	},
//   	DesiredCount: jsii.Number(2),
//   })
//
// Experimental.
type NetworkLoadBalancedTaskImageOptions struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The container name value to be specified in the task definition.
	// Experimental.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port number on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	// Experimental.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// A key/value map of labels to add to the container.
	// Experimental.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

