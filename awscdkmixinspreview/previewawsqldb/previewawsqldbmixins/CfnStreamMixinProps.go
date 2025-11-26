package previewawsqldbmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStreamPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamMixinProps := &CfnStreamMixinProps{
//   	ExclusiveEndTime: jsii.String("exclusiveEndTime"),
//   	InclusiveStartTime: jsii.String("inclusiveStartTime"),
//   	KinesisConfiguration: &KinesisConfigurationProperty{
//   		AggregationEnabled: jsii.Boolean(false),
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	LedgerName: jsii.String("ledgerName"),
//   	RoleArn: jsii.String("roleArn"),
//   	StreamName: jsii.String("streamName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html
//
type CfnStreamMixinProps struct {
	// The exclusive date and time that specifies when the stream ends.
	//
	// If you don't define this parameter, the stream runs indefinitely until you cancel it.
	//
	// The `ExclusiveEndTime` must be in `ISO 8601` date and time format and in Universal Coordinated Time (UTC). For example: `2019-06-13T21:36:34Z` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-exclusiveendtime
	//
	ExclusiveEndTime *string `field:"optional" json:"exclusiveEndTime" yaml:"exclusiveEndTime"`
	// The inclusive start date and time from which to start streaming journal data.
	//
	// This parameter must be in `ISO 8601` date and time format and in Universal Coordinated Time (UTC). For example: `2019-06-13T21:36:34Z` .
	//
	// The `InclusiveStartTime` cannot be in the future and must be before `ExclusiveEndTime` .
	//
	// If you provide an `InclusiveStartTime` that is before the ledger's `CreationDateTime` , QLDB effectively defaults it to the ledger's `CreationDateTime` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-inclusivestarttime
	//
	InclusiveStartTime *string `field:"optional" json:"inclusiveStartTime" yaml:"inclusiveStartTime"`
	// The configuration settings of the Kinesis Data Streams destination for your stream request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-kinesisconfiguration
	//
	KinesisConfiguration interface{} `field:"optional" json:"kinesisConfiguration" yaml:"kinesisConfiguration"`
	// The name of the ledger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-ledgername
	//
	LedgerName *string `field:"optional" json:"ledgerName" yaml:"ledgerName"`
	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
	//
	// To pass a role to QLDB when requesting a journal stream, you must have permissions to perform the `iam:PassRole` action on the IAM role resource. This is required for all journal stream requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name that you want to assign to the QLDB journal stream.
	//
	// User-defined names can help identify and indicate the purpose of a stream.
	//
	// Your stream name must be unique among other *active* streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in [Quotas in Amazon QLDB](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming) in the *Amazon QLDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-stream.html#cfn-qldb-stream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

