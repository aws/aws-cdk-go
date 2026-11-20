package awskinesisanalyticsv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotProps := &CfnSnapshotProps{
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	ApplicationName: jsii.String("applicationName"),
//   	ImportSourceConfiguration: &ImportSourceConfigurationProperty{
//   		ImportRoleArn: jsii.String("importRoleArn"),
//   		SourceS3PrefixLocation: &S3PrefixLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		SourceEncryptionConfiguration: &SnapshotEncryptionConfigurationDescriptionProperty{
//   			KeyType: jsii.String("keyType"),
//
//   			// the properties below are optional
//   			KeyId: jsii.String("keyId"),
//   		},
//   	},
//   	SnapshotDescription: jsii.String("snapshotDescription"),
//   	SnapshotEncryptionConfigurationDescription: &SnapshotEncryptionConfigurationDescriptionProperty{
//   		KeyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		KeyId: jsii.String("keyId"),
//   	},
//   	SnapshotFormat: jsii.String("snapshotFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html
//
type CfnSnapshotProps struct {
	// The customer-chosen name of the snapshot, unique within the account.
	//
	// Immutable after create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-snapshotname
	//
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// The name of the Managed Service for Apache Flink application to snapshot.
	//
	// Used only at create time to select the source application; immutable after create. Specify either ApplicationName or ImportSourceConfiguration, not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The Amazon S3 location to import a snapshot from, and the role the service assumes to read it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-importsourceconfiguration
	//
	ImportSourceConfiguration interface{} `field:"optional" json:"importSourceConfiguration" yaml:"importSourceConfiguration"`
	// Optional customer-supplied description of the snapshot.
	//
	// Immutable after create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-snapshotdescription
	//
	SnapshotDescription *string `field:"optional" json:"snapshotDescription" yaml:"snapshotDescription"`
	// A KMS encryption configuration for snapshot data: the key type, and the key when it is a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription
	//
	SnapshotEncryptionConfigurationDescription interface{} `field:"optional" json:"snapshotEncryptionConfigurationDescription" yaml:"snapshotEncryptionConfigurationDescription"`
	// The format of the snapshot (CANONICAL).
	//
	// Sent to ImportSnapshot with ImportSourceConfiguration; with ApplicationName it may only be CANONICAL, the format CreateSnapshot produces. Returned by DescribeSnapshot. Immutable after create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-snapshotformat
	//
	SnapshotFormat *string `field:"optional" json:"snapshotFormat" yaml:"snapshotFormat"`
	// Tags to apply to the snapshot.
	//
	// Applied at create time (with CreateSnapshot, or with TagResource right after ImportSnapshot) and updated in place with TagResource and UntagResource. Tag keys must not start with aws:.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html#cfn-kinesisanalyticsv2-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

