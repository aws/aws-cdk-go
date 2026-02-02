package awsbedrockagentcore


// The authorizer configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizerConfigurationProperty := &AuthorizerConfigurationProperty{
//   	CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   		DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   		// the properties below are optional
//   		AllowedAudience: []*string{
//   			jsii.String("allowedAudience"),
//   		},
//   		AllowedClients: []*string{
//   			jsii.String("allowedClients"),
//   		},
//   		AllowedScopes: []*string{
//   			jsii.String("allowedScopes"),
//   		},
//   		CustomClaims: []interface{}{
//   			&CustomClaimValidationTypeProperty{
//   				AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   					ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   					ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   						MatchValueString: jsii.String("matchValueString"),
//   						MatchValueStringList: []*string{
//   							jsii.String("matchValueStringList"),
//   						},
//   					},
//   				},
//   				InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   				InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-authorizerconfiguration.html
//
type CfnRuntime_AuthorizerConfigurationProperty struct {
	// Represents inbound authorization configuration options used to authenticate incoming requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-authorizerconfiguration.html#cfn-bedrockagentcore-runtime-authorizerconfiguration-customjwtauthorizer
	//
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

