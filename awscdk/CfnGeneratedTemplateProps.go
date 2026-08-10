package awscdk


// Properties for defining a `CfnGeneratedTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGeneratedTemplateProps := &CfnGeneratedTemplateProps{
//   	GeneratedTemplateName: jsii.String("generatedTemplateName"),
//
//   	// the properties below are optional
//   	TemplateConfiguration: &TemplateConfigurationProperty{
//   		DeletionPolicy: jsii.String("deletionPolicy"),
//   		UpdateReplacePolicy: jsii.String("updateReplacePolicy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html
//
type CfnGeneratedTemplateProps struct {
	// The name assigned to the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html#cfn-cloudformation-generatedtemplate-generatedtemplatename
	//
	GeneratedTemplateName *string `field:"required" json:"generatedTemplateName" yaml:"generatedTemplateName"`
	// The configuration details of the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html#cfn-cloudformation-generatedtemplate-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

