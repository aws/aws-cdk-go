package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for log drivers.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_AwsLogs(&AwsLogDriverProps{
//   		StreamPrefix: jsii.String("EventDemo"),
//   		Mode: ecs.AwsLogDriverMode_NON_BLOCKING,
//   		MaxBufferSize: awscdk.Size_Mebibytes(jsii.Number(25)),
//   	}),
//   })
//
type LogDriver interface {
	// Called when the log driver is configured on a container.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for LogDriver
type jsiiProxy_LogDriver struct {
	_ byte // padding
}

func NewLogDriver_Override(l LogDriver) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.LogDriver",
		nil, // no parameters
		l,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func LogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.LogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := l.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

