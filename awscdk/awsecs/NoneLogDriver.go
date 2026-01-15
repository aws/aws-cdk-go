package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sets the log driver to `none` (no logs collected).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noneLogDriver := awscdk.Aws_ecs.NewNoneLogDriver()
//
type NoneLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for NoneLogDriver
type jsiiProxy_NoneLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the NoneLogDriver class.
func NewNoneLogDriver() NoneLogDriver {
	_init_.Initialize()

	j := jsiiProxy_NoneLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.NoneLogDriver",
		nil, // no parameters
		&j,
	)

	return &j
}

// Constructs a new instance of the NoneLogDriver class.
func NewNoneLogDriver_Override(n NoneLogDriver) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.NoneLogDriver",
		nil, // no parameters
		n,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func NoneLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateNoneLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.NoneLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NoneLogDriver) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := n.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		n,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

