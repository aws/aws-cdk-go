package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Options used to configure a DynamoDB table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stream stream
//
//   tableOptionsV2 := &TableOptionsV2{
//   	ContributorInsights: jsii.Boolean(false),
//   	DeletionProtection: jsii.Boolean(false),
//   	KinesisStream: stream,
//   	PointInTimeRecovery: jsii.Boolean(false),
//   	TableClass: awscdk.Aws_dynamodb.TableClass_STANDARD,
//   }
//
type TableOptionsV2 struct {
	// Whether CloudWatch contributor insights is enabled.
	// Default: false.
	//
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
	// Whether deletion protection is enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Kinesis Data Stream to capture item level changes.
	// Default: - no Kinesis Data Stream.
	//
	KinesisStream awskinesis.IStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// Whether point-in-time recovery is enabled.
	// Default: false.
	//
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The table class.
	// Default: TableClass.STANDARD
	//
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
}

