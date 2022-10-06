package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The event archive base properties.
//
// Example:
//   bus := events.NewEventBus(this, jsii.String("bus"), &eventBusProps{
//   	eventBusName: jsii.String("MyCustomEventBus"),
//   })
//
//   bus.archive(jsii.String("MyArchive"), &baseArchiveProps{
//   	archiveName: jsii.String("MyCustomEventBusArchive"),
//   	description: jsii.String("MyCustomerEventBus Archive"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			awscdk.*stack.of(this).account,
//   		},
//   	},
//   	retention: awscdk.Duration.days(jsii.Number(365)),
//   })
//
// Experimental.
type BaseArchiveProps struct {
	// An event pattern to use to filter events sent to the archive.
	// Experimental.
	EventPattern *EventPattern `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// The name of the archive.
	// Experimental.
	ArchiveName *string `field:"optional" json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	// Experimental.
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
}

