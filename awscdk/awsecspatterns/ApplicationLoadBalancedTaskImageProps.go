package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for configuring a new container.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &ApplicationMultipleTargetGroupsFargateServiceProps{
//   	Cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   		Vpc: *Vpc,
//   	}),
//   	MemoryLimitMiB: jsii.Number(256),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	EnableExecuteCommand: jsii.Boolean(true),
//   	LoadBalancers: []ApplicationLoadBalancerProps{
//   		&ApplicationLoadBalancerProps{
//   			Name: jsii.String("lb"),
//   			IdleTimeout: awscdk.Duration_Seconds(jsii.Number(400)),
//   			DomainName: jsii.String("api.example.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("example.com"),
//   			}),
//   			Listeners: []ApplicationListenerProps{
//   				&ApplicationListenerProps{
//   					Name: jsii.String("listener"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   		&ApplicationLoadBalancerProps{
//   			Name: jsii.String("lb2"),
//   			IdleTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
//   			DomainName: jsii.String("frontend.com"),
//   			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   				ZoneName: jsii.String("frontend.com"),
//   			}),
//   			Listeners: []ApplicationListenerProps{
//   				&ApplicationListenerProps{
//   					Name: jsii.String("listener2"),
//   					Protocol: awscdk.ApplicationProtocol_HTTPS,
//   					Certificate: awscdk.Certificate_*FromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
//   					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   	},
//   	TargetGroups: []ApplicationTargetProps{
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			Listener: jsii.String("listener"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(90),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(443),
//   			Listener: jsii.String("listener2"),
//   		},
//   		&ApplicationTargetProps{
//   			ContainerPort: jsii.Number(80),
//   			PathPattern: jsii.String("a/b/c"),
//   			Priority: jsii.Number(10),
//   			Listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
type ApplicationLoadBalancedTaskImageProps struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, not both.
	// Default: - none.
	//
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The container name value to be specified in the task definition.
	// Default: - web.
	//
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
	// Default: - [80].
	//
	ContainerPorts *[]*float64 `field:"optional" json:"containerPorts" yaml:"containerPorts"`
	// A key/value map of labels to add to the container.
	// Default: - No labels.
	//
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// Flag to indicate whether to enable logging.
	// Default: true.
	//
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The environment variables to pass to the container.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the task execution IAM role that grants the Amazon ECS container agent permission to call AWS APIs on your behalf.
	// Default: - No value.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Default: - Automatically generated name.
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The log driver to use.
	// Default: - AwsLogDriver if enableLogging is true.
	//
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secrets to expose to the container as an environment variable.
	// Default: - No secret environment variables.
	//
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the task IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Default: - A task role is automatically created for you.
	//
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

