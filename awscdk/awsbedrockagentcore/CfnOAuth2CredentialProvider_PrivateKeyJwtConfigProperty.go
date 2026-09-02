package awsbedrockagentcore


// Configuration for private_key_jwt client authentication (RFC 7523).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateKeyJwtConfigProperty := &PrivateKeyJwtConfigProperty{
//   	AdditionalHeaderClaims: map[string]*string{
//   		"additionalHeaderClaimsKey": jsii.String("additionalHeaderClaims"),
//   	},
//   	AdditionalPayloadClaims: map[string]*string{
//   		"additionalPayloadClaimsKey": jsii.String("additionalPayloadClaims"),
//   	},
//   	PrivateKeySource: &PrivateKeySourceProperty{
//   		KmsKeySource: &KmsKeySourceTypeProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig.html
//
type CfnOAuth2CredentialProvider_PrivateKeyJwtConfigProperty struct {
	// A map of additional claims to include in the JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig-additionalheaderclaims
	//
	AdditionalHeaderClaims interface{} `field:"optional" json:"additionalHeaderClaims" yaml:"additionalHeaderClaims"`
	// A map of additional claims to include in the JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig-additionalpayloadclaims
	//
	AdditionalPayloadClaims interface{} `field:"optional" json:"additionalPayloadClaims" yaml:"additionalPayloadClaims"`
	// Contains the private key source configuration for a JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig-privatekeysource
	//
	PrivateKeySource interface{} `field:"optional" json:"privateKeySource" yaml:"privateKeySource"`
	// The algorithm used to sign the JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-privatekeyjwtconfig-signingalgorithm
	//
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

