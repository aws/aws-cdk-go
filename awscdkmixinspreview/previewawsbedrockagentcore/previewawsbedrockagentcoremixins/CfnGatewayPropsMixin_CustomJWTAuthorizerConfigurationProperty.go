package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customJWTAuthorizerConfigurationProperty := &CustomJWTAuthorizerConfigurationProperty{
//   	AllowedAudience: []*string{
//   		jsii.String("allowedAudience"),
//   	},
//   	AllowedClients: []*string{
//   		jsii.String("allowedClients"),
//   	},
//   	AllowedScopes: []*string{
//   		jsii.String("allowedScopes"),
//   	},
//   	CustomClaims: []interface{}{
//   		&CustomClaimValidationTypeProperty{
//   			AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   				ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   				ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   					MatchValueString: jsii.String("matchValueString"),
//   					MatchValueStringList: []*string{
//   						jsii.String("matchValueStringList"),
//   					},
//   				},
//   			},
//   			InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   			InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   		},
//   	},
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html
//
type CfnGatewayPropsMixin_CustomJWTAuthorizerConfigurationProperty struct {
	// The allowed audience authorized for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-allowedaudience
	//
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-allowedclients
	//
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-allowedscopes
	//
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-customclaims
	//
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// The discovery URL for the authorizer configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-discoveryurl
	//
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

