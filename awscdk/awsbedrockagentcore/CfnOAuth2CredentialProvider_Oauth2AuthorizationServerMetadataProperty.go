package awsbedrockagentcore


// Authorization server metadata for the OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oauth2AuthorizationServerMetadataProperty := &Oauth2AuthorizationServerMetadataProperty{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	Issuer: jsii.String("issuer"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   	// the properties below are optional
//   	ResponseTypes: []*string{
//   		jsii.String("responseTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.html
//
type CfnOAuth2CredentialProvider_Oauth2AuthorizationServerMetadataProperty struct {
	// The authorization endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The issuer URL for the OAuth2 authorization server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata-issuer
	//
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata-tokenendpoint
	//
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The supported response types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata-responsetypes
	//
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
}

