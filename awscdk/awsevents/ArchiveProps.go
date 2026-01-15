package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// The event archive properties.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var kmsKey IKey
//
//
//   stack := awscdk.Newstack()
//
//   eventBus := awscdk.NewEventBus(stack, jsii.String("Bus"))
//
//   archive := awscdk.NewArchive(stack, jsii.String("Archive"), &ArchiveProps{
//   	KmsKey: kmsKey,
//   	SourceEventBus: eventBus,
//   	EventPattern: &EventPattern{
//   		Source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
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
	// The customer managed key that encrypts this archive.
	// Default: - Use an AWS managed key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	// Default: - Infinite.
	//
	Retention awscdk.Duration `field:"optional" json:"retention" yaml:"retention"`
	// The event source associated with the archive.
	SourceEventBus interfacesawsevents.IEventBusRef `field:"required" json:"sourceEventBus" yaml:"sourceEventBus"`
}

