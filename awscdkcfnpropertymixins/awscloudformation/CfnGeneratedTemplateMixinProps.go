package awscloudformation


// Properties for CfnGeneratedTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGeneratedTemplateMixinProps := &CfnGeneratedTemplateMixinProps{
//   	GeneratedTemplateName: jsii.String("generatedTemplateName"),
//   	TemplateConfiguration: &TemplateConfigurationProperty{
//   		DeletionPolicy: jsii.String("deletionPolicy"),
//   		UpdateReplacePolicy: jsii.String("updateReplacePolicy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html
//
type CfnGeneratedTemplateMixinProps struct {
	// The name assigned to the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html#cfn-cloudformation-generatedtemplate-generatedtemplatename
	//
	GeneratedTemplateName *string `field:"optional" json:"generatedTemplateName" yaml:"generatedTemplateName"`
	// The configuration details of the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-generatedtemplate.html#cfn-cloudformation-generatedtemplate-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

