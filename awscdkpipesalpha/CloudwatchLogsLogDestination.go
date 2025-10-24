package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// CloudWatch Logs log group for delivery of pipe logs.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//   var logGroup LogGroup
//
//
//   cwlLogDestination := pipes.NewCloudwatchLogsLogDestination(logGroup)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   	LogLevel: pipes.LogLevel_TRACE,
//   	LogIncludeExecutionData: []aLL{
//   		pipes.IncludeExecutionData_*aLL,
//   	},
//   	LogDestinations: []ILogDestination{
//   		cwlLogDestination,
//   	},
//   })
//
// Experimental.
type CloudwatchLogsLogDestination interface {
	ILogDestination
	// Bind the log destination to the pipe.
	// Experimental.
	Bind(_pipe IPipe) *LogDestinationConfig
	// Grant the pipe role to push to the log destination.
	// Experimental.
	GrantPush(pipeRole awsiam.IRole)
}

// The jsii proxy struct for CloudwatchLogsLogDestination
type jsiiProxy_CloudwatchLogsLogDestination struct {
	jsiiProxy_ILogDestination
}

// Experimental.
func NewCloudwatchLogsLogDestination(logGroup awslogs.ILogGroup) CloudwatchLogsLogDestination {
	_init_.Initialize()

	if err := validateNewCloudwatchLogsLogDestinationParameters(logGroup); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchLogsLogDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.CloudwatchLogsLogDestination",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudwatchLogsLogDestination_Override(c CloudwatchLogsLogDestination, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.CloudwatchLogsLogDestination",
		[]interface{}{logGroup},
		c,
	)
}

func (c *jsiiProxy_CloudwatchLogsLogDestination) Bind(_pipe IPipe) *LogDestinationConfig {
	if err := c.validateBindParameters(_pipe); err != nil {
		panic(err)
	}
	var returns *LogDestinationConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_pipe},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogsLogDestination) GrantPush(pipeRole awsiam.IRole) {
	if err := c.validateGrantPushParameters(pipeRole); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"grantPush",
		[]interface{}{pipeRole},
	)
}

