package awsbedrockagentcore


// Contains the KMS configuration for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kmsConfigurationProperty := &KmsConfigurationProperty{
//   	KeyType: jsii.String("keyType"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-tokenvault-kmsconfiguration.html
//
type CfnTokenVaultPropsMixin_KmsConfigurationProperty struct {
	// The type of KMS key (CustomerManagedKey or ServiceManagedKey).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-tokenvault-kmsconfiguration.html#cfn-bedrockagentcore-tokenvault-kmsconfiguration-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The Amazon Resource Name (ARN) of the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-tokenvault-kmsconfiguration.html#cfn-bedrockagentcore-tokenvault-kmsconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

