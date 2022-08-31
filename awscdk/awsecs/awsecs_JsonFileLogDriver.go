package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A log driver that sends log information to json-file Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFileLogDriver := awscdk.Aws_ecs.NewJsonFileLogDriver(&jsonFileLogDriverProps{
//   	compress: jsii.Boolean(false),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxFile: jsii.Number(123),
//   	maxSize: jsii.String("maxSize"),
//   	tag: jsii.String("tag"),
//   })
//
// Experimental.
type JsonFileLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	// Experimental.
	Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JsonFileLogDriver
type jsiiProxy_JsonFileLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JsonFileLogDriver class.
// Experimental.
func NewJsonFileLogDriver(props *JsonFileLogDriverProps) JsonFileLogDriver {
	_init_.Initialize()

	j := jsiiProxy_JsonFileLogDriver{}

	_jsii_.Create(
		"monocdk.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JsonFileLogDriver class.
// Experimental.
func NewJsonFileLogDriver_Override(j JsonFileLogDriver, props *JsonFileLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.JsonFileLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
// Experimental.
func JsonFileLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	var returns LogDriver

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.JsonFileLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsonFileLogDriver) Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

