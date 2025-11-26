package previewawsentityresolutionmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdMappingWorkflowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowMixinProps := &CfnIdMappingWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingIncrementalRunConfig: &IdMappingIncrementalRunConfigProperty{
//   		IncrementalRunType: jsii.String("incrementalRunType"),
//   	},
//   	IdMappingTechniques: &IdMappingTechniquesProperty{
//   		IdMappingType: jsii.String("idMappingType"),
//   		NormalizationVersion: jsii.String("normalizationVersion"),
//   		ProviderProperties: &ProviderPropertiesProperty{
//   			IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   				IntermediateS3Path: jsii.String("intermediateS3Path"),
//   			},
//   			ProviderConfiguration: map[string]*string{
//   				"providerConfigurationKey": jsii.String("providerConfiguration"),
//   			},
//   			ProviderServiceArn: jsii.String("providerServiceArn"),
//   		},
//   		RuleBasedProperties: &IdMappingRuleBasedPropertiesProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   			RecordMatchingModel: jsii.String("recordMatchingModel"),
//   			RuleDefinitionType: jsii.String("ruleDefinitionType"),
//   			Rules: []interface{}{
//   				&RuleProperty{
//   					MatchingKeys: []*string{
//   						jsii.String("matchingKeys"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
//   		},
//   	},
//   	InputSourceConfig: []interface{}{
//   		&IdMappingWorkflowInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaArn: jsii.String("schemaArn"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	OutputSourceConfig: []interface{}{
//   		&IdMappingWorkflowOutputSourceProperty{
//   			KmsArn: jsii.String("kmsArn"),
//   			OutputS3Path: jsii.String("outputS3Path"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html
//
type CfnIdMappingWorkflowMixinProps struct {
	// A description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-idmappingincrementalrunconfig
	//
	IdMappingIncrementalRunConfig interface{} `field:"optional" json:"idMappingIncrementalRunConfig" yaml:"idMappingIncrementalRunConfig"`
	// An object which defines the ID mapping technique and any additional configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques
	//
	IdMappingTechniques interface{} `field:"optional" json:"idMappingTechniques" yaml:"idMappingTechniques"`
	// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-inputsourceconfig
	//
	InputSourceConfig interface{} `field:"optional" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// A list of `IdMappingWorkflowOutputSource` objects, each of which contains fields `outputS3Path` and `KMSArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-outputsourceconfig
	//
	OutputSourceConfig interface{} `field:"optional" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// AWS Entity Resolution assumes this role to create resources on your behalf as part of workflow execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow.
	//
	// There can't be multiple `IdMappingWorkflows` with the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-workflowname
	//
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

