package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// The base class for log drivers.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		token: awscdk.SecretValue.secretsManager(jsii.String("my-splunk-token")),
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
// Experimental.
type LogDriver interface {
	// Called when the log driver is configured on a container.
	// Experimental.
	Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for LogDriver
type jsiiProxy_LogDriver struct {
	_ byte // padding
}

// Experimental.
func NewLogDriver_Override(l LogDriver) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.LogDriver",
		nil, // no parameters
		l,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
// Experimental.
func LogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.LogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogDriver) Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
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

