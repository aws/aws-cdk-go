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
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	RetentionDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html
//
type CfnArchiveProps struct {
	// The ARN of the event bus that sends events to the archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-sourcearn
	//
	SourceArn interface{} `field:"required" json:"sourceArn" yaml:"sourceArn"`
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
	// The identifier of the AWS  customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this archive.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
	//
	// If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the archive.
	//
	// For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// > If you have specified that EventBridge use a customer managed key for encrypting the source event bus, we strongly recommend you also specify a customer managed key for any archives for the event bus as well.
	// >
	// > For more information, see [Encrypting archives](https://docs.aws.amazon.com/eventbridge/latest/userguide/encryption-archives.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-archive.html#cfn-events-archive-retentiondays
	//
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

