package previewawsentityresolutionmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdNamespacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdNamespaceMixinProps := &CfnIdNamespaceMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingWorkflowProperties: []interface{}{
//   		&IdNamespaceIdMappingWorkflowPropertiesProperty{
//   			IdMappingType: jsii.String("idMappingType"),
//   			ProviderProperties: &NamespaceProviderPropertiesProperty{
//   				ProviderConfiguration: map[string]*string{
//   					"providerConfigurationKey": jsii.String("providerConfiguration"),
//   				},
//   				ProviderServiceArn: jsii.String("providerServiceArn"),
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
//   	IdNamespaceName: jsii.String("idNamespaceName"),
//   	InputSourceConfig: []interface{}{
//   		&IdNamespaceInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html
//
type CfnIdNamespaceMixinProps struct {
	// The description of the ID namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-idmappingworkflowproperties
	//
	IdMappingWorkflowProperties interface{} `field:"optional" json:"idMappingWorkflowProperties" yaml:"idMappingWorkflowProperties"`
	// The name of the ID namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-idnamespacename
	//
	IdNamespaceName *string `field:"optional" json:"idNamespaceName" yaml:"idNamespaceName"`
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
	// The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
	//
	// The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
	//
	// The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html#cfn-entityresolution-idnamespace-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

