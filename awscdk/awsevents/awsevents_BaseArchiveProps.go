package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
type BaseArchiveProps struct {
	// An event pattern to use to filter events sent to the archive.
	EventPattern *EventPattern `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// The name of the archive.
	ArchiveName *string `field:"optional" json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
}

