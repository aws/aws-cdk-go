package awsbedrockagentcore


// Contains the private key source configuration for a JWT client assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   privateKeySourceProperty := &PrivateKeySourceProperty{
//   	KmsKeySource: &KmsKeySourceTypeProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeysource.html
//
type CfnOAuth2CredentialProviderPropsMixin_PrivateKeySourceProperty struct {
	// Contains the KMS key configuration for a JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeysource.html#cfn-bedrockagentcore-oauth2credentialprovider-privatekeysource-kmskeysource
	//
	KmsKeySource interface{} `field:"optional" json:"kmsKeySource" yaml:"kmsKeySource"`
}

