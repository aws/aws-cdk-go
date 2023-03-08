package awsdynamodb


// Allows you to specify a KMS key identifier to be used for server-side encryption.
//
// The key can be specified via ARN, key ID, or alias. The key must be created in the same region as the replica.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaSSESpecificationProperty := &replicaSSESpecificationProperty{
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   }
//
type CfnGlobalTable_ReplicaSSESpecificationProperty struct {
	// The AWS KMS key that should be used for the AWS KMS encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key `alias/aws/dynamodb` .
	KmsMasterKeyId *string `field:"required" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

