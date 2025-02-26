package awsopensearchservice


// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestOptionsProperty := &EncryptionAtRestOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html
//
type CfnDomain_EncryptionAtRestOptionsProperty struct {
	// Specify `true` to enable encryption at rest. Required if you enable fine-grained access control in [AdvancedSecurityOptionsInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html) .
	//
	// If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template, the domain is deleted and recreated in order to modify the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html#cfn-opensearchservice-domain-encryptionatrestoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ID. Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a` . Required if you enable encryption at rest.
	//
	// You can also use `keyAlias` as a value.
	//
	// If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template, the domain is deleted and recreated in order to modify the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html#cfn-opensearchservice-domain-encryptionatrestoptions-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

