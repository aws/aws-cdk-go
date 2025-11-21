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
//   replicaSSESpecificationProperty := &ReplicaSSESpecificationProperty{
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicassespecification.html
//
type CfnGlobalTable_ReplicaSSESpecificationProperty struct {
	// The AWS  key that should be used for the AWS  encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key `alias/aws/dynamodb` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicassespecification.html#cfn-dynamodb-globaltable-replicassespecification-kmsmasterkeyid
	//
	KmsMasterKeyId *string `field:"required" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

