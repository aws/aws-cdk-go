package awsqldb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLedger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLedgerProps := &CfnLedgerProps{
//   	PermissionsMode: jsii.String("permissionsMode"),
//
//   	// the properties below are optional
//   	DeletionProtection: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html
//
type CfnLedgerProps struct {
	// The permissions mode to assign to the ledger that you want to create.
	//
	// This parameter can have one of the following values:
	//
	// - `ALLOW_ALL` : A legacy permissions mode that enables access control with API-level granularity for ledgers.
	//
	// This mode allows users who have the `SendCommand` API permission for this ledger to run all PartiQL commands (hence, `ALLOW_ALL` ) on any tables in the specified ledger. This mode disregards any table-level or command-level IAM permissions policies that you create for the ledger.
	// - `STANDARD` : ( *Recommended* ) A permissions mode that enables access control with finer granularity for ledgers, tables, and PartiQL commands.
	//
	// By default, this mode denies all user requests to run any PartiQL commands on any tables in this ledger. To allow PartiQL commands to run, you must create IAM permissions policies for specific table resources and PartiQL actions, in addition to the `SendCommand` API permission for the ledger. For information, see [Getting started with the standard permissions mode](https://docs.aws.amazon.com/qldb/latest/developerguide/getting-started-standard-mode.html) in the *Amazon QLDB Developer Guide* .
	//
	// > We strongly recommend using the `STANDARD` permissions mode to maximize the security of your ledger data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html#cfn-qldb-ledger-permissionsmode
	//
	PermissionsMode *string `field:"required" json:"permissionsMode" yaml:"permissionsMode"`
	// Specifies whether the ledger is protected from being deleted by any user.
	//
	// If not defined during ledger creation, this feature is enabled ( `true` ) by default.
	//
	// If deletion protection is enabled, you must first disable it before you can delete the ledger. You can disable it by calling the `UpdateLedger` operation to set this parameter to `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html#cfn-qldb-ledger-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The key in AWS Key Management Service ( AWS KMS ) to use for encryption of data at rest in the ledger.
	//
	// For more information, see [Encryption at rest](https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html) in the *Amazon QLDB Developer Guide* .
	//
	// Use one of the following options to specify this parameter:
	//
	// - `AWS_OWNED_KMS_KEY` : Use an AWS KMS key that is owned and managed by AWS on your behalf.
	// - *Undefined* : By default, use an AWS owned KMS key.
	// - *A valid symmetric customer managed KMS key* : Use the specified symmetric encryption KMS key in your account that you create, own, and manage.
	//
	// Amazon QLDB does not support asymmetric keys. For more information, see [Using symmetric and asymmetric keys](https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	//
	// To specify a customer managed KMS key, you can use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. When using an alias name, prefix it with `"alias/"` . To specify a key in a different AWS account , you must use the key ARN or alias ARN.
	//
	// For example:
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Alias name: `alias/ExampleAlias`
	// - Alias ARN: `arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias`
	//
	// For more information, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html#cfn-qldb-ledger-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the ledger that you want to create.
	//
	// The name must be unique among all of the ledgers in your AWS account in the current Region.
	//
	// Naming constraints for ledger names are defined in [Quotas in Amazon QLDB](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming) in the *Amazon QLDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html#cfn-qldb-ledger-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qldb-ledger.html#cfn-qldb-ledger-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

