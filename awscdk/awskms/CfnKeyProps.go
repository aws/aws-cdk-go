package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   cfnKeyProps := &CfnKeyProps{
//   	BypassPolicyLockoutSafetyCheck: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	EnableKeyRotation: jsii.Boolean(false),
//   	KeyPolicy: keyPolicy,
//   	KeySpec: jsii.String("keySpec"),
//   	KeyUsage: jsii.String("keyUsage"),
//   	MultiRegion: jsii.Boolean(false),
//   	Origin: jsii.String("origin"),
//   	PendingWindowInDays: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
//
type CfnKeyProps struct {
	// Skips ("bypasses") the key policy lockout safety check. The default value is false.
	//
	// > Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// >
	// > For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the *AWS Key Management Service Developer Guide* .
	//
	// Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-bypasspolicylockoutsafetycheck
	//
	// Default: - false.
	//
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// A description of the KMS key.
	//
	// Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations.
	//
	// When `Enabled` is `true` , the *key state* of the KMS key is `Enabled` . When `Enabled` is `false` , the key state of the KMS key is `Disabled` . The default value is `true` .
	//
	// The actual key state of the KMS key might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html) , [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html) , or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations.
	//
	// For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Enables automatic rotation of the key material for the specified KMS key.
	//
	// By default, automatic key rotation is not enabled.
	//
	// AWS KMS supports automatic rotation only for symmetric encryption KMS keys ( `KeySpec` = `SYMMETRIC_DEFAULT` ). For asymmetric KMS keys, HMAC KMS keys, and KMS keys with Origin `EXTERNAL` , omit the `EnableKeyRotation` property or set it to `false` .
	//
	// To enable automatic key rotation of the key material for a multi-Region KMS key, set `EnableKeyRotation` to `true` on the primary key (created by using `AWS::KMS::Key` ). AWS KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate) in the *AWS Key Management Service Developer Guide* .
	//
	// When you enable automatic rotation, AWS KMS automatically creates new key material for the KMS key one year after the enable date and every year thereafter. AWS KMS retains all key material until you delete the KMS key. For detailed information about automatic key rotation, see [Rotating KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	//
	EnableKeyRotation interface{} `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// The key policy to attach to the KMS key.
	//
	// If you provide a key policy, it must meet the following criteria:
	//
	// - The key policy must allow the caller to make a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) in the *AWS Key Management Service Developer Guide* . (To omit this condition, set `BypassPolicyLockoutSafetyCheck` to true.)
	// - Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to AWS KMS . When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to AWS KMS . For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *AWS Identity and Access Management User Guide* .
	//
	// If you do not provide a key policy, AWS KMS attaches a default key policy to the KMS key. For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the *AWS Key Management Service Developer Guide* .
	//
	// A key policy document can include only the following characters:
	//
	// - Printable ASCII characters
	// - Printable characters in the Basic Latin and Latin-1 Supplement character set
	// - The tab ( `\u0009` ), line feed ( `\u000A` ), and carriage return ( `\u000D` ) special characters
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `32768`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	//
	KeyPolicy interface{} `field:"optional" json:"keyPolicy" yaml:"keyPolicy"`
	// Specifies the type of KMS key to create.
	//
	// The default value, `SYMMETRIC_DEFAULT` , creates a KMS key with a 256-bit symmetric key for encryption and decryption. In China Regions, `SYMMETRIC_DEFAULT` creates a 128-bit symmetric key that uses SM4 encryption. You can't change the `KeySpec` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *AWS Key Management Service Developer Guide* .
	//
	// The `KeySpec` property determines the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see [AWS KMS condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms) in the *AWS Key Management Service Developer Guide* .
	//
	// > If you change the value of the `KeySpec` property on an existing KMS key, the update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) . This prevents you from accidentally deleting a KMS key by changing an immutable property value. > [AWS services that are integrated with AWS KMS](https://docs.aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support encryption with asymmetric KMS keys. For help determining whether a KMS key is asymmetric, see [Identifying asymmetric KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in the *AWS Key Management Service Developer Guide* .
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
	// - SM2 key pairs (China Regions only)
	//
	// - `SM2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
	//
	// Default: - "SYMMETRIC_DEFAULT".
	//
	KeySpec *string `field:"optional" json:"keySpec" yaml:"keySpec"`
	// Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is `ENCRYPT_DECRYPT` . This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the `KeyUsage` value after the KMS key is created.
	//
	// > If you change the value of the `KeyUsage` property on an existing KMS key, the update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) . This prevents you from accidentally deleting a KMS key by changing an immutable property value.
	//
	// Select only one valid value.
	//
	// - For symmetric encryption KMS keys, omit the property or specify `ENCRYPT_DECRYPT` .
	// - For asymmetric KMS keys with RSA key material, specify `ENCRYPT_DECRYPT` or `SIGN_VERIFY` .
	// - For asymmetric KMS keys with ECC key material, specify `SIGN_VERIFY` .
	// - For asymmetric KMS keys with SM2 (China Regions only) key material, specify `ENCRYPT_DECRYPT` or `SIGN_VERIFY` .
	// - For HMAC KMS keys, specify `GENERATE_VERIFY_MAC` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	//
	// Default: - "ENCRYPT_DECRYPT".
	//
	KeyUsage *string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Creates a multi-Region primary key that you can replicate in other AWS Regions .
	//
	// You can't change the `MultiRegion` value after the KMS key is created.
	//
	// For a list of AWS Regions in which multi-Region keys are supported, see [Multi-Region keys in AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the ** .
	//
	// > If you change the value of the `MultiRegion` property on an existing KMS key, the update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) . This prevents you from accidentally deleting a KMS key by changing an immutable property value.
	//
	// For a multi-Region key, set to this property to `true` . For a single-Region key, omit this property or set it to `false` . The default value is `false` .
	//
	// *Multi-Region keys* are an AWS KMS feature that lets you create multiple interoperable KMS keys in different AWS Regions . Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS Region and decrypt it in a different AWS Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You can create a symmetric encryption, HMAC, or asymmetric multi-Region KMS key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.
	//
	// To create a replica of this primary key in a different AWS Region , create an [AWS::KMS::ReplicaKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html) resource in a CloudFormation stack in the replica Region. Specify the key ARN of this primary key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-multiregion
	//
	// Default: - false.
	//
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// The source of the key material for the KMS key.
	//
	// You cannot change the origin after you create the KMS key. The default is `AWS_KMS` , which means that AWS KMS creates the key material.
	//
	// To [create a KMS key with no key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html) (for imported key material), set this value to `EXTERNAL` . For more information about importing key material into AWS KMS , see [Importing Key Material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// You can ignore `ENABLED` when Origin is `EXTERNAL` . When a KMS key with Origin `EXTERNAL` is created, the key state is `PENDING_IMPORT` and `ENABLED` is `false` . After you import the key material, `ENABLED` updated to `true` . The KMS key can then be used for Cryptographic Operations.
	//
	// > AWS CloudFormation doesn't support creating an `Origin` parameter of the `AWS_CLOUDHSM` or `EXTERNAL_KEY_STORE` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-origin
	//
	// Default: - "AWS_KMS".
	//
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
	//
	PendingWindowInDays *float64 `field:"optional" json:"pendingWindowInDays" yaml:"pendingWindowInDays"`
	// Assigns one or more tags to the replica key.
	//
	// > Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *AWS Key Management Service Developer Guide* .
	//
	// For information about tags in AWS KMS , see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) in the *AWS Key Management Service Developer Guide* . For information about tags in CloudFormation, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

