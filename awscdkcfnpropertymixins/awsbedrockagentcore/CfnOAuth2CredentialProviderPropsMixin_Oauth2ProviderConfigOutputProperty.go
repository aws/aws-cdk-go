package awsbedrockagentcore


// Output configuration for an OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oauth2ProviderConfigOutputProperty := &Oauth2ProviderConfigOutputProperty{
//   	ClientAuthenticationMethod: jsii.String("clientAuthenticationMethod"),
//   	ClientId: jsii.String("clientId"),
//   	OauthDiscovery: &Oauth2DiscoveryProperty{
//   		AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			Issuer: jsii.String("issuer"),
//   			ResponseTypes: []*string{
//   				jsii.String("responseTypes"),
//   			},
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		DiscoveryUrl: jsii.String("discoveryUrl"),
//   	},
//   	OnBehalfOfTokenExchangeConfig: &OnBehalfOfTokenExchangeConfigProperty{
//   		GrantType: jsii.String("grantType"),
//   		TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   			ActorTokenContent: jsii.String("actorTokenContent"),
//   			ActorTokenScopes: []*string{
//   				jsii.String("actorTokenScopes"),
//   			},
//   		},
//   	},
//   	PrivateEndpoint: &PrivateEndpointProperty{
//   		ManagedVpcResource: &ManagedVpcResourceProperty{
//   			EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   			RoutingDomain: jsii.String("routingDomain"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   			VpcIdentifier: jsii.String("vpcIdentifier"),
//   		},
//   		SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   			ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   		},
//   	},
//   	PrivateEndpointOverrides: []interface{}{
//   		&PrivateEndpointOverrideProperty{
//   			Domain: jsii.String("domain"),
//   			PrivateEndpoint: &PrivateEndpointProperty{
//   				ManagedVpcResource: &ManagedVpcResourceProperty{
//   					EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   					RoutingDomain: jsii.String("routingDomain"),
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetIds: []*string{
//   						jsii.String("subnetIds"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
//   					VpcIdentifier: jsii.String("vpcIdentifier"),
//   				},
//   				SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   					ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   				},
//   			},
//   		},
//   	},
//   	PrivateKeyJwtConfig: &PrivateKeyJwtConfigProperty{
//   		AdditionalHeaderClaims: map[string]*string{
//   			"additionalHeaderClaimsKey": jsii.String("additionalHeaderClaims"),
//   		},
//   		AdditionalPayloadClaims: map[string]*string{
//   			"additionalPayloadClaimsKey": jsii.String("additionalPayloadClaims"),
//   		},
//   		PrivateKeySource: &PrivateKeySourceProperty{
//   			KmsKeySource: &KmsKeySourceTypeProperty{
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   		SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html
//
type CfnOAuth2CredentialProviderPropsMixin_Oauth2ProviderConfigOutputProperty struct {
	// The client authentication method used when authenticating with the token endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-clientauthenticationmethod
	//
	ClientAuthenticationMethod *string `field:"optional" json:"clientAuthenticationMethod" yaml:"clientAuthenticationMethod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Discovery information for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-oauthdiscovery
	//
	OauthDiscovery interface{} `field:"optional" json:"oauthDiscovery" yaml:"oauthDiscovery"`
	// Configuration for on-behalf-of token exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-onbehalfoftokenexchangeconfig
	//
	OnBehalfOfTokenExchangeConfig interface{} `field:"optional" json:"onBehalfOfTokenExchangeConfig" yaml:"onBehalfOfTokenExchangeConfig"`
	// The private endpoint configuration for connecting to private resources in your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-privateendpoint
	//
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// The list of private endpoint overrides for the OAuth2 provider.
	//
	// Each override maps a specific domain to a private endpoint, enabling secure connectivity through VPC Lattice resource configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-privateendpointoverrides
	//
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
	// Configuration for private_key_jwt client authentication (RFC 7523).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-privatekeyjwtconfig
	//
	PrivateKeyJwtConfig interface{} `field:"optional" json:"privateKeyJwtConfig" yaml:"privateKeyJwtConfig"`
}

