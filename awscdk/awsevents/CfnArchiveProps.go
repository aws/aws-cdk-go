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
//   cfnArchiveProps := &CfnArchiveProps{
//   	SourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	ArchiveName: jsii.String("archiveName"),
//   	Description: jsii.String("description"),
//   	EventPattern: eventPattern,
//   	RetentionDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html
//
type CfnArchiveProps struct {
	// The ARN of the event bus that sends events to the archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-sourcearn
	//
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The name for the archive to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-archivename
	//
	ArchiveName *string `field:"optional" json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An event pattern to use to filter events sent to the archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-eventpattern
	//
	EventPattern interface{} `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-retentiondays
	//
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

