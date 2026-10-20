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
//   	AutoDetectionEnabled: jsii.Boolean(false),
//   	AutoDetectionScope: jsii.String("autoDetectionScope"),
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
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
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
	// Specifies whether auto-detection is requested for the registry.
	//
	// Must be specified together with AutoDetectionScope. Setting this to true is necessary but not sufficient for auto-detection to become active; the preconditions of the configured scope must also be met. To turn auto-detection off, explicitly set this to false - removing AutoDetectionEnabled and AutoDetectionScope from the template is a no-op and leaves the existing auto-detection settings unchanged. A registry cannot be deleted while auto-detection is enabled: set this to false and update the stack before deleting the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-autodetectionenabled
	//
	AutoDetectionEnabled interface{} `field:"optional" json:"autoDetectionEnabled" yaml:"autoDetectionEnabled"`
	// The source from which resources are detected.
	//
	// ORGANIZATION sources resources from all member accounts of an AWS Organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-autodetectionscope
	//
	AutoDetectionScope *string `field:"optional" json:"autoDetectionScope" yaml:"autoDetectionScope"`
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
	// The server-side encryption configuration for a registry.
	//
	// Specifies a customer managed key used to encrypt the registry's content. When omitted, the registry's content is encrypted with an AWS owned key. You cannot change the encryption configuration after registry creation. Specifying a different KMS key, adding this property to an existing registry, or removing it replaces the registry: CloudFormation creates a new registry with a new Amazon Resource Name (ARN) and then deletes the original, including all registry records it contains. Registry records that are not managed by the stack are not re-created in the new registry, and if any remain in the original registry its deletion fails and it is left behind.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags to assign to the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-agentregistry-registry.html#cfn-agentregistry-registry-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

