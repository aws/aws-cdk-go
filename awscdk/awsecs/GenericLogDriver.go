package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends logs to the specified driver.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.NewGenericLogDriver(&GenericLogDriverProps{
//   		LogDriver: jsii.String("fluentd"),
//   		Options: map[string]*string{
//   			"tag": jsii.String("example-tag"),
//   		},
//   	}),
//   })
//
type GenericLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for GenericLogDriver
type jsiiProxy_GenericLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the GenericLogDriver class.
func NewGenericLogDriver(props *GenericLogDriverProps) GenericLogDriver {
	_init_.Initialize()

	if err := validateNewGenericLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenericLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the GenericLogDriver class.
func NewGenericLogDriver_Override(g GenericLogDriver, props *GenericLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		[]interface{}{props},
		g,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func GenericLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateGenericLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.GenericLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := g.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

