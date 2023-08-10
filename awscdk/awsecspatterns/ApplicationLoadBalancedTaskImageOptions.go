package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	TaskSubnets: &SubnetSelection{
//   		Subnets: []iSubnet{
//   			ec2.Subnet_FromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
//   		},
//   	},
//   	LoadBalancerName: jsii.String("application-lb-name"),
//   })
//
type ApplicationLoadBalancedTaskImageOptions struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, not both.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that's passed to the container.
	//
	// If there are multiple arguments, make sure that each argument is a separated string in the array.
	//
	// This parameter maps to `Cmd` in the [Create a container](https://docs.docker.com/engine/api/v1.38/#operation/ContainerCreate) section
	// of the [Docker Remote API](https://docs.docker.com/engine/api/v1.38/) and the `COMMAND` parameter to
	// [docker run](https://docs.docker.com/engine/reference/commandline/run/).
	//
	// For more information about the Docker `CMD` parameter, see https://docs.docker.com/engine/reference/builder/#cmd.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The container name value to be specified in the task definition.
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
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The entry point that's passed to the container.
	//
	// This parameter maps to `Entrypoint` in the [Create a container](https://docs.docker.com/engine/api/v1.38/#operation/ContainerCreate) section
	// of the [Docker Remote API](https://docs.docker.com/engine/api/v1.38/) and the `--entrypoint` option to
	// [docker run](https://docs.docker.com/engine/reference/commandline/run/).
	//
	// For more information about the Docker `ENTRYPOINT` parameter, see https://docs.docker.com/engine/reference/builder/#entrypoint.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
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
	// The secret to expose to the container as an environment variable.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

