package previewawssecretsmanagermixins


// Specifies a `Region` and the `KmsKeyId` for a replica secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicaRegionProperty := &ReplicaRegionProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-replicaregion.html
//
type CfnSecretPropsMixin_ReplicaRegionProperty struct {
	// The ARN, key ID, or alias of the KMS key to encrypt the secret.
	//
	// If you don't include this field, Secrets Manager uses `aws/secretsmanager` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-replicaregion.html#cfn-secretsmanager-secret-replicaregion-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A string that represents a `Region` , for example "us-east-1".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-replicaregion.html#cfn-secretsmanager-secret-replicaregion-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

