package awsbedrockagentcore


// Configuration for RFC 8693 Token Exchange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tokenExchangeGrantTypeConfigProperty := &TokenExchangeGrantTypeConfigProperty{
//   	ActorTokenContent: jsii.String("actorTokenContent"),
//   	ActorTokenScopes: []*string{
//   		jsii.String("actorTokenScopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-tokenexchangegranttypeconfig.html
//
type CfnOAuth2CredentialProviderPropsMixin_TokenExchangeGrantTypeConfigProperty struct {
	// The actor token content type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-tokenexchangegranttypeconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-tokenexchangegranttypeconfig-actortokencontent
	//
	ActorTokenContent *string `field:"optional" json:"actorTokenContent" yaml:"actorTokenContent"`
	// The actor token scopes.
	//
	// Only valid when ActorTokenContent is M2M.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-tokenexchangegranttypeconfig.html#cfn-bedrockagentcore-oauth2credentialprovider-tokenexchangegranttypeconfig-actortokenscopes
	//
	ActorTokenScopes *[]*string `field:"optional" json:"actorTokenScopes" yaml:"actorTokenScopes"`
}

