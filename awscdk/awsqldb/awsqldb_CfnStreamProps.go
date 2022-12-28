package awsqldb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamProps := &cfnStreamProps{
//   	inclusiveStartTime: jsii.String("inclusiveStartTime"),
//   	kinesisConfiguration: &kinesisConfigurationProperty{
//   		aggregationEnabled: jsii.Boolean(false),
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	ledgerName: jsii.String("ledgerName"),
//   	roleArn: jsii.String("roleArn"),
//   	streamName: jsii.String("streamName"),
//
//   	// the properties below are optional
//   	exclusiveEndTime: jsii.String("exclusiveEndTime"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamProps struct {
	// The inclusive start date and time from which to start streaming journal data.
	//
	// This parameter must be in `ISO 8601` date and time format and in Universal Coordinated Time (UTC). For example: `2019-06-13T21:36:34Z` .
	//
	// The `InclusiveStartTime` cannot be in the future and must be before `ExclusiveEndTime` .
	//
	// If you provide an `InclusiveStartTime` that is before the ledger's `CreationDateTime` , QLDB effectively defaults it to the ledger's `CreationDateTime` .
	InclusiveStartTime *string `field:"required" json:"inclusiveStartTime" yaml:"inclusiveStartTime"`
	// The configuration settings of the Kinesis Data Streams destination for your stream request.
	KinesisConfiguration interface{} `field:"required" json:"kinesisConfiguration" yaml:"kinesisConfiguration"`
	// The name of the ledger.
	LedgerName *string `field:"required" json:"ledgerName" yaml:"ledgerName"`
	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
	//
	// To pass a role to QLDB when requesting a journal stream, you must have permissions to perform the `iam:PassRole` action on the IAM role resource. This is required for all journal stream requests.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name that you want to assign to the QLDB journal stream.
	//
	// User-defined names can help identify and indicate the purpose of a stream.
	//
	// Your stream name must be unique among other *active* streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in [Quotas in Amazon QLDB](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming) in the *Amazon QLDB Developer Guide* .
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// The exclusive date and time that specifies when the stream ends.
	//
	// If you don't define this parameter, the stream runs indefinitely until you cancel it.
	//
	// The `ExclusiveEndTime` must be in `ISO 8601` date and time format and in Universal Coordinated Time (UTC). For example: `2019-06-13T21:36:34Z` .
	ExclusiveEndTime *string `field:"optional" json:"exclusiveEndTime" yaml:"exclusiveEndTime"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

