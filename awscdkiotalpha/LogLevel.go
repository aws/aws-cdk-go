package awscdkiotalpha


// The log level for the AWS IoT Logging.
//
// Example:
//   iot.NewLogging(this, jsii.String("Logging"), &LoggingProps{
//   	LogLevel: iot.LogLevel_INFO,
//   })
//
// Experimental.
type LogLevel string

const (
	// Any error that causes an operation to fail.
	//
	// Logs include ERROR information only.
	// Experimental.
	LogLevel_ERROR LogLevel = "ERROR"
	// Anything that can potentially cause inconsistencies in the system, but might not cause the operation to fail.
	//
	// Logs include ERROR and WARN information.
	// Experimental.
	LogLevel_WARN LogLevel = "WARN"
	// High-level information about the flow of things.
	//
	// Logs include INFO, ERROR, and WARN information.
	// Experimental.
	LogLevel_INFO LogLevel = "INFO"
	// Information that might be helpful when debugging a problem.
	//
	// Logs include DEBUG, INFO, ERROR, and WARN information.
	// Experimental.
	LogLevel_DEBUG LogLevel = "DEBUG"
	// All logging is disabled.
	// Experimental.
	LogLevel_DISABLED LogLevel = "DISABLED"
)

