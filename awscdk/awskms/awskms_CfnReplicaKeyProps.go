package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReplicaKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyPolicy interface{}
//
//   cfnReplicaKeyProps := &cfnReplicaKeyProps{
//   	keyPolicy: keyPolicy,
//   	primaryKeyArn: jsii.String("primaryKeyArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	pendingWindowInDays: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReplicaKeyProps struct {
	// The key policy that authorizes use of the replica key.
	//
	// The key policy is not a shared property of multi-Region keys. You can specify the same key policy or a different key policy for each key in a set of related multi-Region keys. AWS KMS does not synchronize this property.
	//
	// The key policy must conform to the following rules.
	//
	// - The key policy must give the caller [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) permission on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the **AWS Key Management Service Developer Guide** .
	// - Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to AWS KMS . When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to AWS KMS . For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *AWS Identity and Access Management User Guide* .
	//
	// A key policy document can include only the following characters:
	//
	// - Printable ASCII characters from the space character ( `\ u0020` ) through the end of the ASCII character range.
	// - Printable characters in the Basic Latin and Latin-1 Supplement character set (through `\ u00FF` ).
	// - The tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ) special characters
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `32768`.
	KeyPolicy interface{} `field:"required" json:"keyPolicy" yaml:"keyPolicy"`
	// Specifies the multi-Region primary key to replicate.
	//
	// The primary key must be in a different AWS Region of the same AWS partition. You can create only one replica of a given primary key in each AWS Region .
	//
	// > If you change the `PrimaryKeyArn` value of a replica key, the existing replica key is scheduled for deletion and a new replica key is created based on the specified primary key. While it is scheduled for deletion, the existing replica key becomes unusable. You can cancel the scheduled deletion of the key outside of CloudFormation.
	// >
	// > However, if you inadvertently delete a replica key, you can decrypt ciphertext encrypted by that replica key by using any related multi-Region key. If necessary, you can recreate the replica in the same Region after the previous one is completely deleted. For details, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide*
	//
	// Specify the key ARN of an existing multi-Region primary key. For example, `arn:aws:kms:us-east-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab` .
	PrimaryKeyArn *string `field:"required" json:"primaryKeyArn" yaml:"primaryKeyArn"`
	// A description of the KMS key.
	//
	// The default value is an empty string (no description).
	//
	// The description is not a shared property of multi-Region keys. You can specify the same description or a different description for each key in a set of related multi-Region keys. AWS Key Management Service does not synchronize this property.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations.
	//
	// When `Enabled` is `true` , the *key state* of the KMS key is `Enabled` . When `Enabled` is `false` , the key state of the KMS key is `Disabled` . The default value is `true` .
	//
	// The actual key state of the replica might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html) , [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html) , or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations. Also, while the replica key is being created, its key state is `Creating` . When the process is complete, the key state of the replica key changes to `Enabled` .
	//
	// For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the number of days in the waiting period before AWS KMS deletes a replica key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// When you remove a replica key from a CloudFormation stack, AWS KMS schedules the replica key for deletion and starts the mandatory waiting period. The `PendingWindowInDays` property determines the length of waiting period. During the waiting period, the key state of replica key is `Pending Deletion` , which prevents it from being used in cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the replica key.
	//
	// If the KMS key is a multi-Region primary key with replica keys, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	//
	// You cannot use a CloudFormation template to cancel deletion of the replica after you remove it from the stack, regardless of the waiting period. However, if you specify a replica key in your template that is based on the same primary key as the original replica key, CloudFormation creates a new replica key with the same key ID, key material, and other shared properties of the original replica key. This new replica key can decrypt ciphertext that was encrypted under the original replica key, or any related multi-Region key.
	//
	// For detailed information about deleting multi-Region keys, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *AWS Key Management Service Developer Guide* .
	//
	// For information about the `PendingDeletion` key state, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *AWS Key Management Service Developer Guide* . For more information about deleting KMS keys, see the [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operation in the *AWS Key Management Service API Reference* and [Deleting KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in the *AWS Key Management Service Developer Guide* .
	//
	// *Minimum* : 7
	//
	// *Maximum* : 30.
	PendingWindowInDays *float64 `field:"optional" json:"pendingWindowInDays" yaml:"pendingWindowInDays"`
	// Assigns one or more tags to the replica key.
	//
	// > Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Tags are not a shared property of multi-Region keys. You can specify the same tags or different tags for each key in a set of related multi-Region keys. AWS KMS does not synchronize this property.
	//
	// Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, AWS KMS replaces the current tag value with the specified one.
	//
	// When you assign tags to an AWS resource, AWS generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

