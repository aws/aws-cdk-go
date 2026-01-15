package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

// Properties for a configuration set event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationSetRef IConfigurationSetRef
//   var eventDestination EventDestination
//
//   configurationSetEventDestinationProps := &ConfigurationSetEventDestinationProps{
//   	ConfigurationSet: configurationSetRef,
//   	Destination: eventDestination,
//
//   	// the properties below are optional
//   	ConfigurationSetEventDestinationName: jsii.String("configurationSetEventDestinationName"),
//   	Enabled: jsii.Boolean(false),
//   	Events: []EmailSendingEvent{
//   		awscdk.Aws_ses.EmailSendingEvent_SEND,
//   	},
//   }
//
type ConfigurationSetEventDestinationProps struct {
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
	// The configuration set that contains the event destination.
	ConfigurationSet interfacesawsses.IConfigurationSetRef `field:"required" json:"configurationSet" yaml:"configurationSet"`
}

