package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for configuring a new container.
//
// Example:
//   // Two network load balancers, each with their own listener and target group.
//   var cluster cluster
//
//   loadBalancedEc2Service := ecsPatterns.NewNetworkMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &NetworkMultipleTargetGroupsEc2ServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(256),
//   	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	LoadBalancers: []networkLoadBalancerProps{
//   		&networkLoadBalancerProps{
//   			Name: jsii.String("lb1"),
//   			Listeners: []networkListenerProps{
//   				&networkListenerProps{
//   					Name: jsii.String("listener1"),
//   				},
//   			},
//   		},
//   		&networkLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			Listeners: []*networkListenerProps{
//   				&networkListenerProps{
//   					Name: jsii.String("listener2"),
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []networkTargetProps{
//   		&networkTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener1"),
//   		},
//   		&networkTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
type NetworkLoadBalancedTaskImageProps struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The container name value to be specified in the task definition.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// A list of port numbers on the container that is bound to the user-specified or automatically assigned host port.
	//
	// If you are using containers in a task with the awsvpc or host network mode, exposed ports should be specified using containerPort.
	// If you are using containers in a task with the bridge network mode and you specify a container port and not a host port,
	// your container automatically receives a host port in the ephemeral port range.
	//
	// Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.
	//
	// For more information, see
	// [hostPort](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PortMapping.html#ECS-Type-PortMapping-hostPort).
	ContainerPorts *[]*float64 `field:"optional" json:"containerPorts" yaml:"containerPorts"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The log driver to use.
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secrets to expose to the container as an environment variable.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

