package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyPolicy interface{}
//
//   cfnKeyProps := &cfnKeyProps{
//   	keyPolicy: keyPolicy,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	enableKeyRotation: jsii.Boolean(false),
//   	keySpec: jsii.String("keySpec"),
//   	keyUsage: jsii.String("keyUsage"),
//   	multiRegion: jsii.Boolean(false),
//   	pendingWindowInDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnKeyProps struct {
	// The key policy that authorizes use of the KMS key. The key policy must conform to the following rules.
	//
	// - The key policy must allow the caller to make a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the **AWS Key Management Service Developer Guide** .
	// - Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to AWS KMS . When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to AWS KMS . For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *AWS Identity and Access Management User Guide* .
	//
	// If you are unsure of which policy to use, consider the *default key policy* . This is the key policy that AWS KMS applies to KMS keys that are created by using the CreateKey API with no specified key policy. It gives the AWS account that owns the key permission to perform all operations on the key. It also allows you write IAM policies to authorize access to the key. For details, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the *AWS Key Management Service Developer Guide* .
	//
	// A key policy document can include only the following characters:
	//
	// - Printable ASCII characters
	// - Printable characters in the Basic Latin and Latin-1 Supplement character set
	// - The tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ) special characters
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `32768`.
	KeyPolicy interface{} `field:"required" json:"keyPolicy" yaml:"keyPolicy"`
	// A description of the KMS key.
	//
	// Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations.
	//
	// When `Enabled` is `true` , the *key state* of the KMS key is `Enabled` . When `Enabled` is `false` , the key state of the KMS key is `Disabled` . The default value is `true` .
	//
	// The actual key state of the KMS key might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html) , [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html) , or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations.
	//
	// For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Enables automatic rotation of the key material for the specified KMS key.
	//
	// By default, automatic key rotation is not enabled.
	//
	// AWS KMS supports automatic rotation only for symmetric encryption KMS keys ( `KeySpec` = `SYMMETRIC_DEFAULT` ). For asymmetric KMS keys and HMAC KMS keys, omit the `EnableKeyRotation` property or set it to `false` .
	//
	// To enable automatic key rotation of the key material for a multi-Region KMS key, set `EnableKeyRotation` to `true` on the primary key (created by using `AWS::KMS::Key` ). AWS KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate) in the *AWS Key Management Service Developer Guide* .
	//
	// When you enable automatic rotation, AWS KMS automatically creates new key material for the KMS key one year after the enable date and every year thereafter. AWS KMS retains all key material until you delete the KMS key. For detailed information about automatic key rotation, see [Rotating KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) in the *AWS Key Management Service Developer Guide* .
	EnableKeyRotation interface{} `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Specifies the type of KMS key to create.
	//
	// The default value, `SYMMETRIC_DEFAULT` , creates a KMS key with a 256-bit symmetric key for encryption and decryption. You can't change the `KeySpec` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *AWS Key Management Service Developer Guide* .
	//
	// The `KeySpec` property determines the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see [AWS KMS condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms) in the *AWS Key Management Service Developer Guide* .
	//
	// > If you change the `KeySpec` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `KeySpec` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted. > [AWS services that are integrated with AWS KMS](https://docs.aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support encryption with asymmetric KMS keys. For help determining whether a KMS key is asymmetric, see [Identifying asymmetric KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in the *AWS Key Management Service Developer Guide* .
	//
	// AWS KMS supports the following key specs for KMS keys:
	//
	// - Symmetric encryption key (default)
	//
	// - `SYMMETRIC_DEFAULT` (AES-256-GCM)
	// - HMAC keys (symmetric)
	//
	// - `HMAC_224`
	// - `HMAC_256`
	// - `HMAC_384`
	// - `HMAC_512`
	// - Asymmetric RSA key pairs
	//
	// - `RSA_2048`
	// - `RSA_3072`
	// - `RSA_4096`
	// - Asymmetric NIST-recommended elliptic curve key pairs
	//
	// - `ECC_NIST_P256` (secp256r1)
	// - `ECC_NIST_P384` (secp384r1)
	// - `ECC_NIST_P521` (secp521r1)
	// - Other asymmetric elliptic curve key pairs
	//
	// - `ECC_SECG_P256K1` (secp256k1), commonly used for cryptocurrencies.
	KeySpec *string `field:"optional" json:"keySpec" yaml:"keySpec"`
	// Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is `ENCRYPT_DECRYPT` . This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the `KeyUsage` value after the KMS key is created.
	//
	// > If you change the `KeyUsage` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `KeyUsage` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted.
	//
	// Select only one valid value.
	//
	// - For symmetric encryption KMS keys, omit the property or specify `ENCRYPT_DECRYPT` .
	// - For asymmetric KMS keys with RSA key material, specify `ENCRYPT_DECRYPT` or `SIGN_VERIFY` .
	// - For asymmetric KMS keys with ECC key material, specify `SIGN_VERIFY` .
	// - For HMAC KMS keys, specify `GENERATE_VERIFY_MAC` .
	KeyUsage *string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Creates a multi-Region primary key that you can replicate in other AWS Regions .
	//
	// You can't change the `MultiRegion` value after the KMS key is created.
	//
	// > If you change the `MultiRegion` value of an existing KMS key, the existing KMS key is scheduled for deletion and a new KMS key is created with the specified `Multi-Region` value. While the scheduled deletion is pending, you can't use the existing KMS key. Unless you [cancel the scheduled deletion](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-keys-scheduling-key-deletion) of the KMS key outside of CloudFormation, all data encrypted under the existing KMS key becomes unrecoverable when the KMS key is deleted.
	//
	// For a multi-Region key, set to this property to `true` . For a single-Region key, omit this property or set it to `false` . The default value is `false` .
	//
	// *Multi-Region keys* are an AWS KMS feature that lets you create multiple interoperable KMS keys in different AWS Regions . Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS Region and decrypt it in a different AWS Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You can create a symmetric encryption, HMAC, or asymmetric multi-Region KMS key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.
	//
	// To create a replica of this primary key in a different AWS Region , create an [AWS::KMS::ReplicaKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html) resource in a CloudFormation stack in the replica Region. Specify the key ARN of this primary key.
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Specifies the number of days in the waiting period before AWS KMS deletes a KMS key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// When you remove a KMS key from a CloudFormation stack, AWS KMS schedules the KMS key for deletion and starts the mandatory waiting period. The `PendingWindowInDays` property determines the length of waiting period. During the waiting period, the key state of KMS key is `Pending Deletion` or `Pending Replica Deletion` , which prevents the KMS key from being used in cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the KMS key.
	//
	// AWS KMS will not delete a [multi-Region primary key](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) that has replica keys. If you remove a multi-Region primary key from a CloudFormation stack, its key state changes to `PendingReplicaDeletion` so it cannot be replicated or used in cryptographic operations. This state can persist indefinitely. When the last of its replica keys is deleted, the key state of the primary key changes to `PendingDeletion` and the waiting period specified by `PendingWindowInDays` begins. When this waiting period expires, AWS KMS deletes the primary key. For details, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You cannot use a CloudFormation template to cancel deletion of the KMS key after you remove it from the stack, regardless of the waiting period. If you specify a KMS key in your template, even one with the same name, CloudFormation creates a new KMS key. To cancel deletion of a KMS key, use the AWS KMS console or the [CancelKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_CancelKeyDeletion.html) operation.
	//
	// For information about the `Pending Deletion` and `Pending Replica Deletion` key states, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* . For more information about deleting KMS keys, see the [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operation in the *AWS Key Management Service API Reference* and [Deleting KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// *Minimum* : 7
	//
	// *Maximum* : 30.
	PendingWindowInDays *float64 `field:"optional" json:"pendingWindowInDays" yaml:"pendingWindowInDays"`
	// Assigns one or more tags to the replica key.
	//
	// > Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *AWS Key Management Service Developer Guide* .
	//
	// For information about tags in AWS KMS , see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) in the *AWS Key Management Service Developer Guide* . For information about tags in CloudFormation, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

