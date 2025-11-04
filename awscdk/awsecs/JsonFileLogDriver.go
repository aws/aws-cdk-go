package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to json-file Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFileLogDriver := awscdk.Aws_ecs.NewJsonFileLogDriver(&JsonFileLogDriverProps{
//   	Compress: jsii.Boolean(false),
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	MaxFile: jsii.Number(123),
//   	MaxSize: jsii.String("maxSize"),
//   	Tag: jsii.String("tag"),
//   })
//
type JsonFileLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JsonFileLogDriver
type jsiiProxy_JsonFileLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JsonFileLogDriver class.
func NewJsonFileLogDriver(props *JsonFileLogDriverProps) JsonFileLogDriver {
	_init_.Initialize()

	if err := validateNewJsonFileLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsonFileLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JsonFileLogDriver class.
func NewJsonFileLogDriver_Override(j JsonFileLogDriver, props *JsonFileLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func JsonFileLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateJsonFileLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.JsonFileLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonFileLogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := j.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

