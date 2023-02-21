package awsevents


// Properties for defining a `CfnArchive`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventPattern interface{}
//
//   cfnArchiveProps := &cfnArchiveProps{
//   	sourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	archiveName: jsii.String("archiveName"),
//   	description: jsii.String("description"),
//   	eventPattern: eventPattern,
//   	retentionDays: jsii.Number(123),
//   }
//
type CfnArchiveProps struct {
	// The ARN of the event bus that sends events to the archive.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The name for the archive to create.
	ArchiveName *string `field:"optional" json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An event pattern to use to filter events sent to the archive.
	EventPattern interface{} `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

