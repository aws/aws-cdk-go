package awsses


// Options for a configuration set event destination.
//
// Example:
//   var myConfigurationSet configurationSet
//   var myTopic topic
//
//
//   myConfigurationSet.AddEventDestination(jsii.String("ToSns"), &ConfigurationSetEventDestinationOptions{
//   	Destination: ses.EventDestination_SnsTopic(myTopic),
//   })
//
type ConfigurationSetEventDestinationOptions struct {
	// The event destination.
	Destination EventDestination `field:"required" json:"destination" yaml:"destination"`
	// A name for the configuration set event destination.
	// Default: - a CloudFormation generated name.
	//
	ConfigurationSetEventDestinationName *string `field:"optional" json:"configurationSetEventDestinationName" yaml:"configurationSetEventDestinationName"`
	// Whether Amazon SES publishes events to this destination.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The type of email sending events to publish to the event destination.
	// Default: - send all event types.
	//
	Events *[]EmailSendingEvent `field:"optional" json:"events" yaml:"events"`
}

