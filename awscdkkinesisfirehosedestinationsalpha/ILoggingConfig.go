package awscdkkinesisfirehosedestinationsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configuration interface for logging errors when data transformation or delivery fails.
//
// This interface defines whether logging is enabled and optionally allows specifying a
// CloudWatch Log Group for storing error logs.
// Deprecated.
type ILoggingConfig interface {
	// If true, log errors when data transformation or data delivery fails.
	//
	// `true` when using `EnableLogging`, `false` when using `DisableLogging`.
	// Deprecated.
	Logging() *bool
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Default: - if `logging` is set to `true`, a log group will be created for you.
	//
	// Deprecated.
	LogGroup() awslogs.ILogGroup
}

// The jsii proxy for ILoggingConfig
type jsiiProxy_ILoggingConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ILoggingConfig) Logging() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggingConfig) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

