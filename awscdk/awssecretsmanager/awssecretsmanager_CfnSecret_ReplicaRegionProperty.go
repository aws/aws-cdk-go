package awssecretsmanager


// Specifies a `Region` and the `KmsKeyId` for a replica secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaRegionProperty := &replicaRegionProperty{
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnSecret_ReplicaRegionProperty struct {
	// (Optional) A string that represents a `Region` , for example "us-east-1".
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN, key ID, or alias of the KMS key to encrypt the secret.
	//
	// If you don't include this field, Secrets Manager uses `aws/secretsmanager` .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

