package awskinesisanalyticsv2


// A KMS encryption configuration for snapshot data: the key type, and the key when it is a customer managed key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotEncryptionConfigurationDescriptionProperty := &SnapshotEncryptionConfigurationDescriptionProperty{
//   	KeyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	KeyId: jsii.String("keyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription.html
//
type CfnSnapshot_SnapshotEncryptionConfigurationDescriptionProperty struct {
	// The type of KMS key: AWS_OWNED_KEY or CUSTOMER_MANAGED_KEY.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription.html#cfn-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription-keytype
	//
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ARN of the customer managed KMS key.
	//
	// Required when KeyType is CUSTOMER_MANAGED_KEY; must be omitted for AWS_OWNED_KEY.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription.html#cfn-kinesisanalyticsv2-snapshot-snapshotencryptionconfigurationdescription-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

