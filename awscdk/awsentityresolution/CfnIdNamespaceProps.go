package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIdNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdNamespaceProps := &CfnIdNamespaceProps{
//   	IdNamespaceName: jsii.String("idNamespaceName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	IdMappingWorkflowProperties: []interface{}{
//   		&IdNamespaceIdMappingWorkflowPropertiesProperty{
//   			IdMappingType: jsii.String("idMappingType"),
//
//   			// the properties below are optional
//   			ProviderProperties: &NamespaceProviderPropertiesProperty{
//   				ProviderServiceArn: jsii.String("providerServiceArn"),
//
//   				// the properties below are optional
//   				ProviderConfiguration: map[string]*string{
//   					"providerConfigurationKey": jsii.String("providerConfiguration"),
//   				},
//   			},
//   			RuleBasedProperties: &NamespaceRuleBasedPropertiesProperty{
//   				AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   				RecordMatchingModels: []*string{
//   					jsii.String("recordMatchingModels"),
//   				},
//   				RuleDefinitionTypes: []*string{
//   					jsii.String("ruleDefinitionTypes"),
//   				},
//   				Rules: []interface{}{
//   					&RuleProperty{
//   						MatchingKeys: []*string{
//   							jsii.String("matchingKeys"),
//   						},
//   						RuleName: jsii.String("ruleName"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	InputSourceConfig: []interface{}{
//   		&IdNamespaceInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//
//   			// the properties below are optional
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html
//
type CfnIdNamespaceProps struct {
	// The name of the ID namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-idnamespacename
	//
	IdNamespaceName *string `field:"required" json:"idNamespaceName" yaml:"idNamespaceName"`
	// The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
	//
	// The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
	//
	// The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the ID namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-idmappingworkflowproperties
	//
	IdMappingWorkflowProperties interface{} `field:"optional" json:"idMappingWorkflowProperties" yaml:"idMappingWorkflowProperties"`
	// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-inputsourceconfig
	//
	InputSourceConfig interface{} `field:"optional" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

