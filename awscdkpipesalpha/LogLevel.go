package awscdkpipesalpha


// Log configuration for a pipe.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//   var logGroup logGroup
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
//   	LogDestinations: []iLogDestination{
//   		cwlLogDestination,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html#eb-pipes-logs-level
//
// Experimental.
type LogLevel string

const (
	// No logging.
	// Experimental.
	LogLevel_OFF LogLevel = "OFF"
	// Log only errors.
	// Experimental.
	LogLevel_ERROR LogLevel = "ERROR"
	// Log errors, warnings, and info.
	// Experimental.
	LogLevel_INFO LogLevel = "INFO"
	// Log everything.
	// Experimental.
	LogLevel_TRACE LogLevel = "TRACE"
)

