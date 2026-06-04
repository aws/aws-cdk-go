package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthCredentialProviderProperty := &OAuthCredentialProviderProperty{
//   	ProviderArn: jsii.String("providerArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   	DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   	GrantType: jsii.String("grantType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html
//
type CfnHarness_OAuthCredentialProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html#cfn-bedrockagentcore-harness-oauthcredentialprovider-providerarn
	//
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html#cfn-bedrockagentcore-harness-oauthcredentialprovider-scopes
	//
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html#cfn-bedrockagentcore-harness-oauthcredentialprovider-customparameters
	//
	CustomParameters interface{} `field:"optional" json:"customParameters" yaml:"customParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html#cfn-bedrockagentcore-harness-oauthcredentialprovider-defaultreturnurl
	//
	DefaultReturnUrl *string `field:"optional" json:"defaultReturnUrl" yaml:"defaultReturnUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-oauthcredentialprovider.html#cfn-bedrockagentcore-harness-oauthcredentialprovider-granttype
	//
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
}

