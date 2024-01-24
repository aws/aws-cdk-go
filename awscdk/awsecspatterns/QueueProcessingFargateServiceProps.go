package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The properties for the QueueProcessingFargateService service.
//
// Example:
//   var cluster cluster
//
//   cluster.EnableFargateCapacityProviders()
//
//   queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(512),
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("FARGATE_SPOT"),
//   			Weight: jsii.Number(2),
//   		},
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("FARGATE"),
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type QueueProcessingFargateServiceProps struct {
	// A list of Capacity Provider strategies used to place a service.
	// Default: - undefined.
	//
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Default: - disabled.
	//
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Default: - create a new cluster; if both cluster and vpc are omitted, a new VPC will be created for you.
	//
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Default: - CMD value built into container image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Grace period after scaling activity in seconds.
	//
	// Subsequent scale outs during the cooldown period are squashed so that only
	// the biggest scale out happens.
	//
	// Subsequent scale ins during the cooldown period are ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	// Default: 300 seconds.
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The target CPU utilization percentage for CPU based scaling strategy when enabled.
	// Default: - 50.
	//
	CpuTargetUtilizationPercent *float64 `field:"optional" json:"cpuTargetUtilizationPercent" yaml:"cpuTargetUtilizationPercent"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Default: - Rolling update (ECS).
	//
	DeploymentController *awsecs.DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// Flag to disable CPU based auto scaling strategy on the service.
	// Default: - false.
	//
	DisableCpuBasedScaling *bool `field:"optional" json:"disableCpuBasedScaling" yaml:"disableCpuBasedScaling"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Default: false.
	//
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether ECS Exec should be enabled.
	// Default: - false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Flag to indicate whether to enable logging.
	// Default: true.
	//
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The environment variables to pass to the container.
	//
	// The variable `QUEUE_NAME` with value `queue.queueName` will
	// always be passed.
	// Default: 'QUEUE_NAME: queue.queueName'
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The name of a family that the task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Default: - Automatically generated name.
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The image used to start a container.
	//
	// For `QueueProcessingFargateService`, either `image` or `taskDefinition` must be specified, but not both.
	// For `QueueProcessingEc2Service`, `image` is required.
	// Default: - the image of the task definition is used for Fargate, required otherwise.
	//
	Image awsecs.ContainerImage `field:"optional" json:"image" yaml:"image"`
	// The log driver to use.
	// Default: - AwsLogDriver if enableLogging is true.
	//
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Default: - default from underlying service.
	//
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The maximum number of times that a message can be received by consumers.
	//
	// When this value is exceeded for a message the message will be automatically sent to the Dead Letter Queue.
	//
	// If the queue construct is specified, maxReceiveCount should be omitted.
	// Default: 3.
	//
	MaxReceiveCount *float64 `field:"optional" json:"maxReceiveCount" yaml:"maxReceiveCount"`
	// Maximum capacity to scale to.
	// Default: - If the feature flag, ECS_REMOVE_DEFAULT_DESIRED_COUNT is false, the default is (desiredTaskCount * 2); if true, the default is 2.
	//
	MaxScalingCapacity *float64 `field:"optional" json:"maxScalingCapacity" yaml:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - default from underlying service.
	//
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Default: - If the feature flag, ECS_REMOVE_DEFAULT_DESIRED_COUNT is false, the default is the desiredTaskCount; if true, the default is 1.
	//
	MinScalingCapacity *float64 `field:"optional" json:"minScalingCapacity" yaml:"minScalingCapacity"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Default: - none.
	//
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A queue for which to process items from.
	//
	// If specified and this is a FIFO queue, the queue name must end in the string '.fifo'. See
	// [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)
	// Default: 'SQSQueue with CloudFormation-generated name'.
	//
	Queue awssqs.IQueue `field:"optional" json:"queue" yaml:"queue"`
	// The number of seconds that Dead Letter Queue retains a message.
	//
	// If the queue construct is specified, retentionPeriod should be omitted.
	// Default: Duration.days(14)
	//
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The intervals for scaling based on the SQS queue's ApproximateNumberOfMessagesVisible metric.
	//
	// Maps a range of metric values to a particular scaling behavior. See
	// [Simple and Step Scaling Policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
	// Default: [{ upper: 0, change: -1 },{ lower: 100, change: +1 },{ lower: 500, change: +5 }].
	//
	ScalingSteps *[]*awsapplicationautoscaling.ScalingInterval `field:"optional" json:"scalingSteps" yaml:"scalingSteps"`
	// The secret to expose to the container as an environment variable.
	// Default: - No secret environment variables.
	//
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the service.
	// Default: - CloudFormation-generated name.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message and delete it from the queue
	// before it becomes visible again for dequeueing by another processor. Values must be between 0 and (12 hours).
	//
	// If the queue construct is specified, visibilityTimeout should be omitted.
	// Default: Duration.seconds(30)
	//
	VisibilityTimeout awscdk.Duration `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Default: - uses the VPC defined in the cluster or creates a new VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// 8192 (8 vCPU) - Available memory values: Between 16GB and 60GB in 4GB increments
	//
	// 16384 (16 vCPU) - Available memory values: Between 32GB and 120GB in 8GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 256.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB) - Available cpu values: 8192 (8 vCPU)
	//
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Default: 512.
	//
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Default: Latest.
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	// Default: - If the property is undefined, `operatingSystemFamily` is LINUX and `cpuArchitecture` is X86_64.
	//
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - none.
	//
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Optional name for the container added.
	//
	// This name is not used when `taskDefinition` is provided.
	// Default: - QueueProcessingContainer.
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The health check command and associated configuration parameters for the container.
	// Default: - Health check configuration from container.
	//
	HealthCheck *awsecs.HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Default: - A new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	// Default: - Public subnets if `assignPublicIp` is set, otherwise the first available one of Private, Isolated, Public, in that order.
	//
	TaskSubnets *awsec2.SubnetSelection `field:"optional" json:"taskSubnets" yaml:"taskSubnets"`
}

