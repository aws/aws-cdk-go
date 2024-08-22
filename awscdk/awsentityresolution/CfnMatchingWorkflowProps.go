package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMatchingWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMatchingWorkflowProps := &CfnMatchingWorkflowProps{
//   	InputSourceConfig: []interface{}{
//   		&InputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaArn: jsii.String("schemaArn"),
//
//   			// the properties below are optional
//   			ApplyNormalization: jsii.Boolean(false),
//   		},
//   	},
//   	OutputSourceConfig: []interface{}{
//   		&OutputSourceProperty{
//   			Output: []interface{}{
//   				&OutputAttributeProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Hashed: jsii.Boolean(false),
//   				},
//   			},
//   			OutputS3Path: jsii.String("outputS3Path"),
//
//   			// the properties below are optional
//   			ApplyNormalization: jsii.Boolean(false),
//   			KmsArn: jsii.String("kmsArn"),
//   		},
//   	},
//   	ResolutionTechniques: &ResolutionTechniquesProperty{
//   		ProviderProperties: &ProviderPropertiesProperty{
//   			ProviderServiceArn: jsii.String("providerServiceArn"),
//
//   			// the properties below are optional
//   			IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   				IntermediateS3Path: jsii.String("intermediateS3Path"),
//   			},
//   			ProviderConfiguration: map[string]*string{
//   				"providerConfigurationKey": jsii.String("providerConfiguration"),
//   			},
//   		},
//   		ResolutionType: jsii.String("resolutionType"),
//   		RuleBasedProperties: &RuleBasedPropertiesProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   			Rules: []interface{}{
//   				&RuleProperty{
//   					MatchingKeys: []*string{
//   						jsii.String("matchingKeys"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			MatchPurpose: jsii.String("matchPurpose"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	WorkflowName: jsii.String("workflowName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html
//
type CfnMatchingWorkflowProps struct {
	// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-inputsourceconfig
	//
	InputSourceConfig interface{} `field:"required" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// A list of `OutputSource` objects, each of which contains fields `OutputS3Path` , `ApplyNormalization` , and `Output` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-outputsourceconfig
	//
	OutputSourceConfig interface{} `field:"required" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// An object which defines the `resolutionType` and the `ruleBasedProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-resolutiontechniques
	//
	ResolutionTechniques interface{} `field:"required" json:"resolutionTechniques" yaml:"resolutionTechniques"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// AWS Entity Resolution assumes this role to create resources on your behalf as part of workflow execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the workflow.
	//
	// There can't be multiple `MatchingWorkflows` with the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-workflowname
	//
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// A description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

