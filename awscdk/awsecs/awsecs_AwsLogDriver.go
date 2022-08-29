package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to CloudWatch Logs.
//
// Example:
//   var cluster cluster
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewAwsLogDriver(&awsLogDriverProps{
//   		streamPrefix: jsii.String("EventDemo"),
//   		mode: ecs.awsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.addTarget(targets.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	taskCount: jsii.Number(1),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerName: jsii.String("TheContainer"),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("I_WAS_TRIGGERED"),
//   					value: jsii.String("From CloudWatch Events"),
//   				},
//   			},
//   		},
//   	},
//   }))
//
type AwsLogDriver interface {
	LogDriver
	// The log group to send log streams to.
	//
	// Only available after the LogDriver has been bound to a ContainerDefinition.
	LogGroup() awslogs.ILogGroup
	SetLogGroup(val awslogs.ILogGroup)
	// Called when the log driver is configured on a container.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for AwsLogDriver
type jsiiProxy_AwsLogDriver struct {
	jsiiProxy_LogDriver
}

func (j *jsiiProxy_AwsLogDriver) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}


// Constructs a new instance of the AwsLogDriver class.
func NewAwsLogDriver(props *AwsLogDriverProps) AwsLogDriver {
	_init_.Initialize()

	j := jsiiProxy_AwsLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AwsLogDriver class.
func NewAwsLogDriver_Override(a AwsLogDriver, props *AwsLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AwsLogDriver) SetLogGroup(val awslogs.ILogGroup) {
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func AwsLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AwsLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

