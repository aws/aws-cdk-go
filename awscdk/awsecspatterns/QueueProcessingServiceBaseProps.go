package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The properties for the base QueueProcessingEc2Service or QueueProcessingFargateService service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//   var containerImage ContainerImage
//   var logDriver LogDriver
//   var queue Queue
//   var secret Secret
//   var vpc Vpc
//
//   queueProcessingServiceBaseProps := &QueueProcessingServiceBaseProps{
//   	CapacityProviderStrategies: []CapacityProviderStrategy{
//   		&CapacityProviderStrategy{
//   			CapacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			Base: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	CircuitBreaker: &DeploymentCircuitBreaker{
//   		Enable: jsii.Boolean(false),
//   		Rollback: jsii.Boolean(false),
//   	},
//   	Cluster: cluster,
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	CpuTargetUtilizationPercent: jsii.Number(123),
//   	DeploymentController: &DeploymentController{
//   		Type: awscdk.Aws_ecs.DeploymentControllerType_ECS,
//   	},
//   	DisableCpuBasedScaling: jsii.Boolean(false),
//   	EnableECSManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	EnableLogging: jsii.Boolean(false),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Family: jsii.String("family"),
//   	Image: containerImage,
//   	LogDriver: logDriver,
//   	MaxHealthyPercent: jsii.Number(123),
//   	MaxReceiveCount: jsii.Number(123),
//   	MaxScalingCapacity: jsii.Number(123),
//   	MinHealthyPercent: jsii.Number(123),
//   	MinScalingCapacity: jsii.Number(123),
//   	PropagateTags: awscdk.*Aws_ecs.PropagatedTagSource_SERVICE,
//   	Queue: queue,
//   	RetentionPeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	ScalingSteps: []ScalingInterval{
//   		&ScalingInterval{
//   			Change: jsii.Number(123),
//
//   			// the properties below are optional
//   			Lower: jsii.Number(123),
//   			Upper: jsii.Number(123),
//   		},
//   	},
//   	Secrets: map[string]Secret{
//   		"secretsKey": secret,
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	VisibilityTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	Vpc: vpc,
//   }
//
type QueueProcessingServiceBaseProps struct {
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
	// Default: 2.
	//
	MaxScalingCapacity *float64 `field:"optional" json:"maxScalingCapacity" yaml:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Default: - default from underlying service.
	//
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Default: 1.
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
}

