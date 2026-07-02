package awsbedrockagentcore


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
//   		PrivateEndpoint: &PrivateEndpointProperty{
//   			ManagedVpcResource: &ManagedVpcResourceProperty{
//   				EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   				VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   				// the properties below are optional
//   				RoutingDomain: jsii.String("routingDomain"),
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Tags: map[string]*string{
//   					"tagsKey": jsii.String("tags"),
//   				},
//   			},
//   			SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   				ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   			},
//   		},
//   		PrivateEndpointOverrides: []interface{}{
//   			&PrivateEndpointOverrideProperty{
//   				Domain: jsii.String("domain"),
//   				PrivateEndpoint: &PrivateEndpointProperty{
//   					ManagedVpcResource: &ManagedVpcResourceProperty{
//   						EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   						SubnetIds: []*string{
//   							jsii.String("subnetIds"),
//   						},
//   						VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   						// the properties below are optional
//   						RoutingDomain: jsii.String("routingDomain"),
//   						SecurityGroupIds: []*string{
//   							jsii.String("securityGroupIds"),
//   						},
//   						Tags: map[string]*string{
//   							"tagsKey": jsii.String("tags"),
//   						},
//   					},
//   					SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   						ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizerconfiguration.html
//
type CfnHarness_AuthorizerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizerconfiguration.html#cfn-bedrockagentcore-harness-authorizerconfiguration-customjwtauthorizer
	//
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

