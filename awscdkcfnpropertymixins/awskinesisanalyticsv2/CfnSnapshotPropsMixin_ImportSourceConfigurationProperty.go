package awskinesisanalyticsv2


// The Amazon S3 location to import a snapshot from, and the role the service assumes to read it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   importSourceConfigurationProperty := &ImportSourceConfigurationProperty{
//   	ImportRoleArn: jsii.String("importRoleArn"),
//   	SourceEncryptionConfiguration: &SnapshotEncryptionConfigurationDescriptionProperty{
//   		KeyId: jsii.String("keyId"),
//   		KeyType: jsii.String("keyType"),
//   	},
//   	SourceS3PrefixLocation: &S3PrefixLocationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-importsourceconfiguration.html
//
type CfnSnapshotPropsMixin_ImportSourceConfigurationProperty struct {
	// The ARN of the IAM role that the service assumes to read the snapshot files from Amazon S3 (and to decrypt them, if they are encrypted with a customer managed key).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-importsourceconfiguration.html#cfn-kinesisanalyticsv2-snapshot-importsourceconfiguration-importrolearn
	//
	ImportRoleArn *string `field:"optional" json:"importRoleArn" yaml:"importRoleArn"`
	// A KMS encryption configuration for snapshot data: the key type, and the key when it is a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-importsourceconfiguration.html#cfn-kinesisanalyticsv2-snapshot-importsourceconfiguration-sourceencryptionconfiguration
	//
	SourceEncryptionConfiguration interface{} `field:"optional" json:"sourceEncryptionConfiguration" yaml:"sourceEncryptionConfiguration"`
	// An Amazon S3 prefix location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-importsourceconfiguration.html#cfn-kinesisanalyticsv2-snapshot-importsourceconfiguration-sources3prefixlocation
	//
	SourceS3PrefixLocation interface{} `field:"optional" json:"sourceS3PrefixLocation" yaml:"sourceS3PrefixLocation"`
	// The type of storage the snapshot is imported from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-importsourceconfiguration.html#cfn-kinesisanalyticsv2-snapshot-importsourceconfiguration-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

