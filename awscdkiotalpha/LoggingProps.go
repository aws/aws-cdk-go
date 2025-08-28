package awscdkiotalpha


// Properties for defining AWS IoT Logging.
//
// Example:
//   iot.NewLogging(this, jsii.String("Logging"), &LoggingProps{
//   	LogLevel: iot.LogLevel_INFO,
//   })
//
// Experimental.
type LoggingProps struct {
	// The log level for the AWS IoT Logging.
	// Default: LogLevel.ERROR
	//
	// Experimental.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
}

