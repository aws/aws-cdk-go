package awsbedrockagentcore


// Discovery information for an OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oauth2DiscoveryProperty := &Oauth2DiscoveryProperty{
//   	AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		Issuer: jsii.String("issuer"),
//   		ResponseTypes: []*string{
//   			jsii.String("responseTypes"),
//   		},
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   	},
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2discovery.html
//
type CfnOAuth2CredentialProviderPropsMixin_Oauth2DiscoveryProperty struct {
	// Authorization server metadata for the OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2discovery.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2discovery-authorizationservermetadata
	//
	AuthorizationServerMetadata interface{} `field:"optional" json:"authorizationServerMetadata" yaml:"authorizationServerMetadata"`
	// The discovery URL for the OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2discovery.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2discovery-discoveryurl
	//
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

