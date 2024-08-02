package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIdMappingWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdMappingWorkflowProps := &CfnIdMappingWorkflowProps{
//   	IdMappingTechniques: &IdMappingTechniquesProperty{
//   		IdMappingType: jsii.String("idMappingType"),
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
//   	},
//   	InputSourceConfig: []interface{}{
//   		&IdMappingWorkflowInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//
//   			// the properties below are optional
//   			SchemaArn: jsii.String("schemaArn"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	WorkflowName: jsii.String("workflowName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	OutputSourceConfig: []interface{}{
//   		&IdMappingWorkflowOutputSourceProperty{
//   			OutputS3Path: jsii.String("outputS3Path"),
//
//   			// the properties below are optional
//   			KmsArn: jsii.String("kmsArn"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html
//
type CfnIdMappingWorkflowProps struct {
	// An object which defines the ID mapping technique and any additional configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-idmappingtechniques
	//
	IdMappingTechniques interface{} `field:"required" json:"idMappingTechniques" yaml:"idMappingTechniques"`
	// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-inputsourceconfig
	//
	InputSourceConfig interface{} `field:"required" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// The Amazon Resource Name (ARN) of the IAM role.
	//
	// AWS Entity Resolution assumes this role to create resources on your behalf as part of workflow execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the workflow.
	//
	// There can't be multiple `IdMappingWorkflows` with the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-workflowname
	//
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// A description of the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of `IdMappingWorkflowOutputSource` objects, each of which contains fields `OutputS3Path` and `Output` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-outputsourceconfig
	//
	OutputSourceConfig interface{} `field:"optional" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html#cfn-entityresolution-idmappingworkflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

