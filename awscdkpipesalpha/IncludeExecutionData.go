package awscdkpipesalpha


// Log data configuration for a pipe.
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
// Experimental.
type IncludeExecutionData string

const (
	// Specify ALL to include the execution data (specifically, the payload, awsRequest, and awsResponse fields) in the log messages for this pipe.
	// Experimental.
	IncludeExecutionData_ALL IncludeExecutionData = "ALL"
)

