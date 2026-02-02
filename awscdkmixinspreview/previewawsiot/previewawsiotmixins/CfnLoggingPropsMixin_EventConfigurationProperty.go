package previewawsiotmixins


// Configuration for event-based logging that specifies which event types to log and their logging settings.
//
// Used for account-level logging overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventConfigurationProperty := &EventConfigurationProperty{
//   	EventType: jsii.String("eventType"),
//   	LogDestination: jsii.String("logDestination"),
//   	LogLevel: jsii.String("logLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-logging-eventconfiguration.html
//
type CfnLoggingPropsMixin_EventConfigurationProperty struct {
	// The type of event to log.
	//
	// These include event types like Connect, Publish, and Disconnect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-logging-eventconfiguration.html#cfn-iot-logging-eventconfiguration-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// CloudWatch Log Group for event-based logging.
	//
	// Specifies where log events should be sent. The log destination for event-based logging overrides default Log Group for the specified event type and applies to all resources associated with that event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-logging-eventconfiguration.html#cfn-iot-logging-eventconfiguration-logdestination
	//
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// The logging level for the specified event type.
	//
	// Determines the verbosity of log messages generated for this event type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-logging-eventconfiguration.html#cfn-iot-logging-eventconfiguration-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

