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
//   archiveProps := &ArchiveProps{
//   	EventPattern: &EventPattern{
//   		Account: []*string{
//   			jsii.String("account"),
//   		},
//   		Detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		DetailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Source: []*string{
//   			jsii.String("source"),
//   		},
//   		Time: []*string{
//   			jsii.String("time"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	SourceEventBus: eventBus,
//
//   	// the properties below are optional
//   	ArchiveName: jsii.String("archiveName"),
//   	Description: jsii.String("description"),
//   	Retention: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type ArchiveProps struct {
	// An event pattern to use to filter events sent to the archive.
	EventPattern *EventPattern `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// The name of the archive.
	// Default: - Automatically generated.
	//
	ArchiveName *string `field:"optional" json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	// Default: - none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	// Default: - Infinite.
	//
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
	// The event source associated with the archive.
	SourceEventBus IEventBus `field:"required" json:"sourceEventBus" yaml:"sourceEventBus"`
}

