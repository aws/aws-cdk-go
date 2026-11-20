package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customJWTAuthorizerConfigurationProperty := &CustomJWTAuthorizerConfigurationProperty{
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html
//
type CfnPaymentManager_CustomJWTAuthorizerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration-discoveryurl
	//
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration-allowedaudience
	//
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration-allowedclients
	//
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration-allowedscopes
	//
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-paymentmanager-customjwtauthorizerconfiguration-customclaims
	//
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
}

