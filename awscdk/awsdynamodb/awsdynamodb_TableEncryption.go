package awsdynamodb


// What kind of server-side encryption to apply to this table.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   	encryption: dynamodb.tableEncryption_CUSTOMER_MANAGED,
//   })
//
//   // You can access the CMK that was added to the stack on your behalf by the Table construct via:
//   tableEncryptionKey := table.encryptionKey
//
type TableEncryption string

const (
	// Server-side KMS encryption with a master key owned by AWS.
	TableEncryption_DEFAULT TableEncryption = "DEFAULT"
	// Server-side KMS encryption with a customer master key managed by customer.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	//
	// > **NOTE**: if `encryptionKey` is not specified and the `Table` construct creates
	// > a KMS key for you, the key will be created with default permissions. If you are using
	// > CDKv2, these permissions will be sufficient to enable the key for use with DynamoDB tables.
	// > If you are using CDKv1, make sure the feature flag `@aws-cdk/aws-kms:defaultKeyPolicies`
	// > is set to `true` in your `cdk.json`.
	TableEncryption_CUSTOMER_MANAGED TableEncryption = "CUSTOMER_MANAGED"
	// Server-side KMS encryption with a master key managed by AWS.
	TableEncryption_AWS_MANAGED TableEncryption = "AWS_MANAGED"
)

