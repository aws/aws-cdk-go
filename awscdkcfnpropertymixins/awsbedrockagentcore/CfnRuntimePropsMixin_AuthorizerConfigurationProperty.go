package awsbedrockagentcore


// The authorizer configuration.
//
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
//   		AllowedWorkloadConfiguration: &AllowedWorkloadConfigurationProperty{
//   			HostingEnvironments: []interface{}{
//   				&HostingEnvironmentProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//   			WorkloadIdentities: []*string{
//   				jsii.String("workloadIdentities"),
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-authorizerconfiguration.html
//
type CfnRuntimePropsMixin_AuthorizerConfigurationProperty struct {
	// Represents inbound authorization configuration options used to authenticate incoming requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-authorizerconfiguration.html#cfn-bedrockagentcore-runtime-authorizerconfiguration-customjwtauthorizer
	//
	CustomJwtAuthorizer interface{} `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

