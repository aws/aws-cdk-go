package awsnovaact


// Properties for CfnWorkflowDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkflowDefinitionMixinProps := &CfnWorkflowDefinitionMixinProps{
//   	Description: jsii.String("description"),
//   	ExportConfig: &WorkflowExportConfigProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-novaact-workflowdefinition.html
//
type CfnWorkflowDefinitionMixinProps struct {
	// An optional description of the workflow definition's purpose and functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-novaact-workflowdefinition.html#cfn-novaact-workflowdefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration settings for exporting workflow execution data and logs to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-novaact-workflowdefinition.html#cfn-novaact-workflowdefinition-exportconfig
	//
	ExportConfig interface{} `field:"optional" json:"exportConfig" yaml:"exportConfig"`
	// The name of the workflow definition.
	//
	// Must be unique within your account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-novaact-workflowdefinition.html#cfn-novaact-workflowdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

