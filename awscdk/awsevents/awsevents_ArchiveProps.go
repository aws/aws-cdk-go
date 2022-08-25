package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The event archive properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var detail interface{}
//   var eventBus eventBus
//
//   archiveProps := &archiveProps{
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			jsii.String("account"),
//   		},
//   		detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		detailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		id: []*string{
//   			jsii.String("id"),
//   		},
//   		region: []*string{
//   			jsii.String("region"),
//   		},
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   		time: []*string{
//   			jsii.String("time"),
//   		},
//   		version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	sourceEventBus: eventBus,
//
//   	// the properties below are optional
//   	archiveName: jsii.String("archiveName"),
//   	description: jsii.String("description"),
//   	retention: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type ArchiveProps struct {
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
	// The event source associated with the archive.
	SourceEventBus IEventBus `field:"required" json:"sourceEventBus" yaml:"sourceEventBus"`
}

