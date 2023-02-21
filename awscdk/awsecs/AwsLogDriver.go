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
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.NewAwsLogDriver(&AwsLogDriverProps{
//   		StreamPrefix: jsii.String("EventDemo"),
//   		Mode: ecs.AwsLogDriverMode_NON_BLOCKING,
//   	}),
//   })
//
//   // An Rule that describes the event trigger (in this case a scheduled run)
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Expression(jsii.String("rate(1 min)")),
//   })
//
//   // Pass an environment variable to the container 'TheContainer' in the task
//   rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	TaskCount: jsii.Number(1),
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerName: jsii.String("TheContainer"),
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("I_WAS_TRIGGERED"),
//   					Value: jsii.String("From CloudWatch Events"),
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

	if err := validateNewAwsLogDriverParameters(props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_AwsLogDriver)SetLogGroup(val awslogs.ILogGroup) {
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func AwsLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateAwsLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
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
	if err := a.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

