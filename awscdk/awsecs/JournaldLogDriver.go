package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   journaldLogDriver := awscdk.Aws_ecs.NewJournaldLogDriver(&JournaldLogDriverProps{
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	Tag: jsii.String("tag"),
//   })
//
type JournaldLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JournaldLogDriver
type jsiiProxy_JournaldLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JournaldLogDriver class.
func NewJournaldLogDriver(props *JournaldLogDriverProps) JournaldLogDriver {
	_init_.Initialize()

	if err := validateNewJournaldLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JournaldLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JournaldLogDriver class.
func NewJournaldLogDriver_Override(j JournaldLogDriver, props *JournaldLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func JournaldLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateJournaldLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.JournaldLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JournaldLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := j.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

