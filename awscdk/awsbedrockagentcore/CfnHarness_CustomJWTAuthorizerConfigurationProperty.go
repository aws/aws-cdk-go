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
//   	PrivateEndpoint: &PrivateEndpointProperty{
//   		ManagedVpcResource: &ManagedVpcResourceProperty{
//   			EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   			// the properties below are optional
//   			RoutingDomain: jsii.String("routingDomain"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
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
//   					SubnetIds: []*string{
//   						jsii.String("subnetIds"),
//   					},
//   					VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   					// the properties below are optional
//   					RoutingDomain: jsii.String("routingDomain"),
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
//   				},
//   				SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   					ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html
//
type CfnHarness_CustomJWTAuthorizerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-discoveryurl
	//
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-allowedaudience
	//
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-allowedclients
	//
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-allowedscopes
	//
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-customclaims
	//
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// Private endpoint configuration for connecting to the OpenID Connect discovery endpoint over a private network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-privateendpoint
	//
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-privateendpointoverrides
	//
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
}

