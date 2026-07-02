package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html
//
type CfnHarnessPropsMixin_CustomJWTAuthorizerConfigurationProperty struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-discoveryurl
	//
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Private endpoint configuration for connecting to the OpenID Connect discovery endpoint over a private network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-privateendpoint
	//
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-harness-customjwtauthorizerconfiguration-privateendpointoverrides
	//
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
}

