package awsbedrockagentcore


// Configuration for on-behalf-of token exchange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onBehalfOfTokenExchangeConfigProperty := &OnBehalfOfTokenExchangeConfigProperty{
//   	GrantType: jsii.String("grantType"),
//
//   	// the properties below are optional
//   	TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   		ActorTokenContent: jsii.String("actorTokenContent"),
//
//   		// the properties below are optional
//   		ActorTokenScopes: []*string{
//   			jsii.String("actorTokenScopes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-onbehalfoftokenexchangeconfig.html
//
type CfnOAuth2CredentialProvider_OnBehalfOfTokenExchangeConfigProperty struct {
	// The grant type for on-behalf-of token exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-onbehalfoftokenexchangeconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-onbehalfoftokenexchangeconfig-granttype
	//
	GrantType *string `field:"required" json:"grantType" yaml:"grantType"`
	// Configuration for RFC 8693 Token Exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-onbehalfoftokenexchangeconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-onbehalfoftokenexchangeconfig-tokenexchangegranttypeconfig
	//
	TokenExchangeGrantTypeConfig interface{} `field:"optional" json:"tokenExchangeGrantTypeConfig" yaml:"tokenExchangeGrantTypeConfig"`
}

