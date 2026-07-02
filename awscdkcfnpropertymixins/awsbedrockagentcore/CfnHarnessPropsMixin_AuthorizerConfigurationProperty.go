package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizerConfigurationProperty := &AuthorizerConfigurationProperty{
//   	CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
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
//   		DiscoveryUrl: jsii.String("discoveryUrl"),
//   		PrivateEndpoint: &PrivateEndpointProperty{
//   			ManagedVpcResource: &ManagedVpcResourceProperty{
//   				EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   				RoutingDomain: jsii.String("routingDomain"),
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   				Tags: map[string]*string{
//   					"tagsKey": jsii.String("tags"),
//   				},
//   				VpcIdentifier: jsii.String("vpcIdentifier"),
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
//   						RoutingDomain: jsii.String("routingDomain"),
//   						SecurityGroupIds: []*string{
//   							jsii.String("securityGroupIds"),
//   						},
//   						SubnetIds: []*string{
//   							jsii.String("subnetIds"),
//   						},
//   						Tags: map[string]*string{
//   							"tagsKey": jsii.String("tags"),
//   						},
//   						VpcIdentifier: jsii.String("vpcIdentifier"),
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
type CfnHarnessPropsMixin_AuthorizerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizerconfiguration.html#cfn-bedrockagentcore-harness-authorizerconfiguration-customjwtauthorizer
	//
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

