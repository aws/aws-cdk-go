package awsdynamodb


// Represents the settings used to enable server-side encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSESpecificationProperty := &sSESpecificationProperty{
//   	sseEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	sseType: jsii.String("sseType"),
//   }
//
type CfnTable_SSESpecificationProperty struct {
	// Indicates whether server-side encryption is done using an AWS managed key or an AWS owned key.
	//
	// If enabled (true), server-side encryption type is set to `KMS` and an AWS managed key is used ( AWS KMS charges apply). If disabled (false) or not specified, server-side encryption is set to AWS owned key.
	SseEnabled interface{} `field:"required" json:"sseEnabled" yaml:"sseEnabled"`
	// The AWS KMS key that should be used for the AWS KMS encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key `alias/aws/dynamodb` .
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Server-side encryption type. The only supported value is:.
	//
	// - `KMS` - Server-side encryption that uses AWS Key Management Service . The key is stored in your account and is managed by AWS KMS ( AWS KMS charges apply).
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

