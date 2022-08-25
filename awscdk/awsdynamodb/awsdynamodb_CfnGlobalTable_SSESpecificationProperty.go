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
//   	sseType: jsii.String("sseType"),
//   }
//
type CfnGlobalTable_SSESpecificationProperty struct {
	// Indicates whether server-side encryption is performed using an AWS managed key or an AWS owned key.
	//
	// If disabled (false) or not specified, server-side encryption uses an AWS owned key. If enabled (true), the server-side encryption type is set to KMS and an AWS managed key is used ( AWS KMS charges apply). If you choose to use KMS encryption, you can also use customer managed KMS keys by specifying them in the `ReplicaSpecification.SSESpecification` object. You cannot mix AWS managed and customer managed KMS keys.
	SseEnabled interface{} `field:"required" json:"sseEnabled" yaml:"sseEnabled"`
	// Server-side encryption type. The only supported value is:.
	//
	// - `KMS` - Server-side encryption that uses AWS Key Management Service . The key is stored in your account and is managed by AWS KMS ( AWS KMS charges apply).
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

