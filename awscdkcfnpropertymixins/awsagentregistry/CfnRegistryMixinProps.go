package awsagentregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRegistryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRegistryMixinProps := &CfnRegistryMixinProps{
//   	ApprovalConfiguration: &ApprovalConfigurationProperty{
//   		AutoApprovalRules: []*string{
//   			jsii.String("autoApprovalRules"),
//   		},
//   	},
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Description: jsii.String("description"),
//   	DiscoveryConfiguration: &DiscoveryConfigurationProperty{
//   		AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   			CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   				AllowedAudience: []*string{
//   					jsii.String("allowedAudience"),
//   				},
//   				AllowedClients: []*string{
//   					jsii.String("allowedClients"),
//   				},
//   				AllowedScopes: []*string{
//   					jsii.String("allowedScopes"),
//   				},
//   				CustomClaims: []interface{}{
//   					&CustomClaimValidationTypeProperty{
//   						AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   							ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   							ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   								MatchValueString: jsii.String("matchValueString"),
//   								MatchValueStringList: []*string{
//   									jsii.String("matchValueStringList"),
//   								},
//   							},
//   						},
//   						InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   						InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   					},
//   				},
//   				DiscoveryUrl: jsii.String("discoveryUrl"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html
//
type CfnRegistryMixinProps struct {
	// Configuration for the registry's record approval workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-approvalconfiguration
	//
	ApprovalConfiguration interface{} `field:"optional" json:"approvalConfiguration" yaml:"approvalConfiguration"`
	// The type of authorizer that controls how consumers access the registry's search and MCP invoke operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-authorizertype
	//
	AuthorizerType *string `field:"optional" json:"authorizerType" yaml:"authorizerType"`
	// The description of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Discovery configuration for the registry.
	//
	// Controls how consumers are authorized to search the registry and invoke its MCP endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-discoveryconfiguration
	//
	DiscoveryConfiguration interface{} `field:"optional" json:"discoveryConfiguration" yaml:"discoveryConfiguration"`
	// The name of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags to assign to the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

