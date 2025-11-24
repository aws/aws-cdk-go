package mixinsawsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMatchingWorkflowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowMixinProps := &CfnMatchingWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	IncrementalRunConfig: &IncrementalRunConfigProperty{
//   		IncrementalRunType: jsii.String("incrementalRunType"),
//   	},
//   	InputSourceConfig: []interface{}{
//   		&InputSourceProperty{
//   			ApplyNormalization: jsii.Boolean(false),
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaArn: jsii.String("schemaArn"),
//   		},
//   	},
//   	OutputSourceConfig: []interface{}{
//   		&OutputSourceProperty{
//   			ApplyNormalization: jsii.Boolean(false),
//   			KmsArn: jsii.String("kmsArn"),
//   			Output: []interface{}{
//   				&OutputAttributeProperty{
//   					Hashed: jsii.Boolean(false),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			OutputS3Path: jsii.String("outputS3Path"),
//   		},
//   	},
//   	ResolutionTechniques: &ResolutionTechniquesProperty{
//   		ProviderProperties: &ProviderPropertiesProperty{
//   			IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   				IntermediateS3Path: jsii.String("intermediateS3Path"),
//   			},
//   			ProviderConfiguration: map[string]*string{
//   				"providerConfigurationKey": jsii.String("providerConfiguration"),
//   			},
//   			ProviderServiceArn: jsii.String("providerServiceArn"),
//   		},
//   		ResolutionType: jsii.String("resolutionType"),
//   		RuleBasedProperties: &RuleBasedPropertiesProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   			MatchPurpose: jsii.String("matchPurpose"),
//   			Rules: []interface{}{
//   				&RuleProperty{
//   					MatchingKeys: []*string{
//   						jsii.String("matchingKeys"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
//   		},
//   		RuleConditionProperties: &RuleConditionPropertiesProperty{
//   			Rules: []interface{}{
//   				&RuleConditionProperty{
//   					Condition: jsii.String("condition"),
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html
//
type CfnMatchingWorkflowMixinProps struct {
	// A description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// An object that defines the incremental run type. This object contains only the `incrementalRunType` field, which appears as "Automatic" in the console.
	//
	// > For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER` , incremental processing is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-incrementalrunconfig
	//
	IncrementalRunConfig interface{} `field:"optional" json:"incrementalRunConfig" yaml:"incrementalRunConfig"`
	// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-inputsourceconfig
	//
	InputSourceConfig interface{} `field:"optional" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// A list of `OutputSource` objects, each of which contains fields `outputS3Path` , `applyNormalization` , `KMSArn` , and `output` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-outputsourceconfig
	//
	OutputSourceConfig interface{} `field:"optional" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// An object which defines the `resolutionType` and the `ruleBasedProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-resolutiontechniques
	//
	ResolutionTechniques interface{} `field:"optional" json:"resolutionTechniques" yaml:"resolutionTechniques"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// AWS Entity Resolution assumes this role to create resources on your behalf as part of workflow execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workflow.
	//
	// There can't be multiple `MatchingWorkflows` with the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html#cfn-entityresolution-matchingworkflow-workflowname
	//
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

